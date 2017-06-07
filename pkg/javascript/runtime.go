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

// NewRuntime returns a new Runtime. Access the vm via the VM() accessor.
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
	// Execute src in the js vm.
	if ret, err = rt.vm.Run(src); err != nil {
		return otto.UndefinedValue(), err
	}
	defer rt.stop()
	// This is the main event loop. When a timing is added as a result
	// of `setTimeout` or `setInterval`, a *time.Timer is set to elapse
	// after the provided duration (given in the api call, e,g,
	// `setTimeout(timeoutCallback(), 2000)`).
	for {
		select {
		// Once the timer elapses, it's pushed on the runtime ready channel.
		// The runtime then invokes the callback given in the setTimeout
		// or setInterval call (in `fire`).
		case tm := <-rt.ready:
			if err := rt.fire(tm); err != nil {
				fmt.Printf("error: firing timer: %v\n", err)
				return otto.UndefinedValue(), err
			}
		default: // noop -- to avoid blocking on ready channel.
		}
		// if there are no times being tracked, then we can assume
		// all timers that were set (if any) have elapsed. We can
		// safely assume the script has finished executing and exit.
		if len(rt.times) == 0 {
			return
		}
	}
	return ret, nil
}

// VM is an accessor for the js vm.
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
