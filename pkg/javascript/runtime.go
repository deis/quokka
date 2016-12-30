package javascript

import (
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

type Runtime struct {
	VM *otto.Otto
}

func NewRuntime() *Runtime {
	return &Runtime{
		VM: otto.New(),
	}
}
