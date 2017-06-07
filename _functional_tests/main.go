package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/deis/quokka/pkg/javascript"
	"github.com/deis/quokka/pkg/javascript/libk8s"
	"github.com/spf13/cobra"
)

const usage = `
ftest is a functional test runner for quokka.

It is hard-coded to run the functional test suite found inside of the
_functional_tests directory in the Quokka source code.

It executes each functional JavaScript file in sequence, running against
whatever cluster $KUBECONFIG points to.
`

var until string

func main() {
	cmd := &cobra.Command{
		Use:   "ftest",
		Short: "run functional tests for quokka",
		Long:  usage,
		Run:   run,
	}
	cmd.PersistentFlags().StringVarP(&until, "until", "u", "", "Stop when it hits the named test (and don't execute that test).")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
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
		if strings.Contains(f, until) {
			break
		}
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
	if err := libk8s.Register(rt.VM()); err != nil {
		return err
	}
	_, err := rt.Run(script)
	return err
}
