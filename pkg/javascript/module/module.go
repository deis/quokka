package module

import (
	"fmt"

	"github.com/robertkrimen/otto"
	//"github.com/technosophos/ottomatic"

	js "github.com/deis/quokka/pkg/javascript/jsutil"
)

// DefaultLoader creates a strategic loader backed by a filesystem loader.
//
// This is suitable for loading off of a local filesystem.
func DefaultLoader(path string) FileLoader {
	return NewStrategicLoader(FilesystemLoader(path))
}

// Module describes a JavaScript module.
//
// The structure is based off of Node.js's module function, but compatibility
// is not ensured.
// https://nodejs.org/api/modules.html
type Module struct {
	Require  func(string) otto.Value `otto:"require"`
	Children []otto.Value            `otto:"children"`
	Exports  otto.Value              `otto:"exports"`
	Filename string                  `otto:"filename"`
	Id       string                  `otto:"id"`
	Loaded   bool                    `otto:"loaded"`
	Parent   otto.Value              `otto:"parent"`
}

// Register registers the module library to a JS runtime.
//
// It is advised to use the StategicLoader to wrap loader here. That will
// handle the various Node.js loading strategies.
func Register(vm *otto.Otto, loader FileLoader) error {
	requireFn := func(filename string) otto.Value {
		data, err := loader.Load(filename)
		js.POE(err)
		script := wrapModule(data)
		fn, err := vm.Eval(script)
		js.POE(err)
		return fn
	}
	return vm.Set("require", requireFn)
}

// This is largely derived from Node.js documentation
const moduleWrapper = `
(function () {
  var module = {exports: {}};
  (function (module, exports) {
// +src
%s
// -src
;
  })(module, module.exports);
  return module.exports;
})();
`

func wrapModule(data []byte) string {
	return fmt.Sprintf(moduleWrapper, string(data))
}
