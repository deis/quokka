package javascript

import (
	"github.com/robertkrimen/otto"
	"time"
)

type timing struct {
	duration time.Duration
	interval bool
	timer    *time.Timer
	call     otto.FunctionCall
}

// clearTiming implements clearInterval and clearTimeout.
//
// clearInterval(<timer>)
// clearTimeout(<timer>)
func clearTiming(rt *Runtime) func(otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		if ifc, err := call.Argument(0).Export(); err != nil {
			if t, ok := ifc.(*timing); ok && t != nil {
				t.timer.Stop()
				rt.del(t)
			}
		}
		return otto.UndefinedValue()
	}
}

// setTiming is used to implement setInterval and setTimeout.
//
//
// <timer> = setTimeout(<function>, <delay>, [<arguments...>])
// <timer> = setInterval(<function>, <delay>, [<arguments...>])
func setTiming(rt *Runtime, interval bool) func(otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		return newTiming(rt, interval)(call)
	}
}

func newTiming(rt *Runtime, interval bool) func(otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		dur, err := call.Argument(1).ToInteger()
		switch {
		// What to do here: panic or log error? I think panic
		// for now... Alternative is to log the error and use
		// a default delay 'd' of 1 (or whatever the default
		// may be.)
		case err != nil:
			panic(err)
		case dur <= 0:
			dur = 1
		}
		t := &timing{duration: time.Duration(dur) * time.Millisecond, interval: interval, call: call}
		t.timer = time.AfterFunc(t.duration, func() {
			rt.ready <- t
		})

		v, err := call.Otto.ToValue(t)
		if err != nil {
			// Again, not sure how we handle these errors?
			// The beautiful evil hell of dynamic langs...
			panic(err)
		}

		rt.add(t)
		return v
	}
}
