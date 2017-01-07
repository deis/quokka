package jsutil

import (
	"encoding/json"
	"fmt"

	"github.com/robertkrimen/otto"
)

// POE is panic-on-error designed to bubble errors into the JS runtime.
//
// We use a custom error type KubernetesError to indicate the source of
// the problem.
//
// In the JavaScript runtime, throw/catch is implemented using panic. So
// to bubble an error to a try/catch, we need to panic.
func POE(err error) {
	if err != nil {
		//panic(otto.MakeCustomError(KubernetesErrorType, err.Error()))
		v, _ := otto.ToValue(err.Error())
		panic(v)
	}
}

// Remarshal turns src into JSON, and then unmarshals it into dest.
//
// This is a cheap and easy way to take generic data from JavaScript and convert
// it to strongly typed data in Go.
func Remarshal(src, dest interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("Remarshal %T=>%T (1): %s", src, dest, err)
	}
	if err = json.Unmarshal(data, dest); err != nil {
		return fmt.Errorf("Remarshal %T=>%T (2): %s", src, dest, err)
	}
	return nil
}
