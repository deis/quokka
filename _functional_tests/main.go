package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/deis/quokka/pkg/javascript"
	"github.com/deis/quokka/pkg/javascript/libk8s"
)

func main() {
	testloc, _ := filepath.Abs("./_functional_tests/*.js")

	scripts, err := filepath.Glob(testloc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(scripts) == 0 {
		fmt.Printf("Found no scripts to execute at %q.\n", testloc)
		os.Exit(1)
	}

	failures := []string{}
	for _, f := range scripts {
		d, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Printf("failed to load %s: %s\n", f, err)
			os.Exit(1)
		}
		if err := execScript(d); err != nil {
			fmt.Printf("script execution failed for %s: %s\n", f, err)
			failures = append(failures, f)
		}
	}

	if len(failures) > 0 {
		fmt.Printf("FAILED: %d scripts failed: %v", len(failures), failures)
		os.Exit(1)
	}
	fmt.Printf("all %d scripts executed cleanly\n", len(scripts))
}

func execScript(script []byte) error {
	rt := javascript.NewRuntime()
	if err := libk8s.Register(rt.VM); err != nil {
		return err
	}
	_, err := rt.VM.Run(script)
	return err
}
