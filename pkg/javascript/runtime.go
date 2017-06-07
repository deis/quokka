package javascript

import (
	"fmt"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

type Runtime struct {
	vm    *otto.Otto
	ready chan *timing
	times map[*timing]struct{}
}

func NewRuntime() *Runtime {
	rt := &Runtime{
		vm:    otto.New(),
		ready: make(chan *timing),
		times: make(map[*timing]struct{}),
	}
	rt.vm.Set("clearInterval", clearTiming(rt))
	rt.vm.Set("clearTimeout", clearTiming(rt))
	rt.vm.Set("setInterval", setTiming(rt, true))
	rt.vm.Set("setTimeout", setTiming(rt, false))
	return rt
}

// Run will execute the given JavaScript, continuing to run until all timers
// have finished executing (if any). The following functions are included in
// the js vm:
//
//		clearInterval
// 		clearTimeout
// 		setInterval
//		setTimeout
func (rt *Runtime) Run(src interface{}) (ret otto.Value, err error) {
	if rt == nil {
		rt = NewRuntime()
	}
	if ret, err = rt.vm.Run(src); err != nil {
		return otto.UndefinedValue(), err
	}
	defer rt.stop()
	for {
		select {
		case tm := <-rt.ready:
			if err := rt.fire(tm); err != nil {
				fmt.Printf("error: firing timer: %v\n", err)
				return otto.UndefinedValue(), err
			}
		default:
		}
		if len(rt.times) == 0 {
			return
		}
	}
	return ret, nil
}

func (rt *Runtime) VM() *otto.Otto { return rt.vm }

// add adds the timing to the runtime.
func (rt *Runtime) add(t *timing) { rt.times[t] = struct{}{} }

// del deletes the timing from the runtime.
func (rt *Runtime) del(t *timing) { delete(rt.times, t) }

// stop stops and deletes all timers currently tracked by the runtime.
func (rt *Runtime) stop() {
	for t := range rt.times {
		t.timer.Stop()
		delete(rt.times, t)
	}
}

// fire handles an elapsed timeout (or interval) invoking the callback
// provided to setTimeout (or setInerval.)
func (rt *Runtime) fire(t *timing) error {
	// TODO: Refactor this block. It's weird and isn't
	// very intelligble at a glance. The idea is that
	// ```t.call.ArgumentList[2:]``` contains the args
	// for the callback given to setTimer and setInterval.
	// This block is directly from natto.
	var args []interface{}
	if len(t.call.ArgumentList) > 2 {
		tmp := t.call.ArgumentList[2:]
		args = make([]interface{}, 2+len(tmp))
		for i, arg := range tmp {
			args[i+2] = arg
		}
	} else {
		args = make([]interface{}, 1)
	}
	args[0] = t.call.ArgumentList[0]
	if _, err := rt.vm.Call("Function.call.call", nil, args...); err != nil {
		return err
	}
	if t.interval {
		t.timer.Reset(t.duration)
	} else {
		rt.del(t)
	}
	return nil
}
