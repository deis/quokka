package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/deis/quokka/pkg/javascript"
	"github.com/deis/quokka/pkg/javascript/libk8s"
	"github.com/deis/quokka/pkg/javascript/module"
	"github.com/spf13/cobra"
)

const usage = `
ftest is a functional test runner for quokka.

It is hard-coded to run the functional test suite found inside of the
_functional_tests directory in the Quokka source code.

It executes each functional JavaScript file in sequence, running against
whatever cluster $KUBECONFIG points to.
`

var (
	until   string
	verbose bool
)

var defaultTests = []string{"module", "libk8s"}

func main() {
	cmd := &cobra.Command{
		Use:   "ftest [dir1 [dir2 [...]]",
		Short: "run functional tests for quokka",
		Long:  usage,
		Run:   run,
	}
	cmd.PersistentFlags().StringVarP(&until, "until", "u", "", "Stop when it hits the named test (and don't execute that test).")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		args = []string{"module", "libk8s"}
	}
	cmd.Printf("##### Running test suites for %s\n", strings.Join(args, ", "))
	for _, dir := range args {
		runDir(cmd, dir)
	}
}

func runDir(cmd *cobra.Command, dir string) {
	testloc, _ := filepath.Abs(filepath.Join("_functional_tests", dir, "*.js"))
	cmd.Printf("#### Running tests in %s #### \n", testloc)

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
		if until != "" && strings.Contains(f, until) {
			break
		}
		if verbose {
			cmd.Printf("----> Executing %q\n", f)
		}
		d, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Printf("FAIL: failed to load %s: %s\n", f, err)
			os.Exit(1)
		}
		if err := execScript(d); err != nil {
			fmt.Printf("FAIL: script execution failed for %s: %s\n", f, err)
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
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	dest := filepath.Join(cwd, "_functional_tests")
	if verbose {
		fmt.Printf("module loader root set to %q", dest)
	}
	if err := module.Register(rt.VM(), module.DefaultLoader(dest)); err != nil {
		return err
	}
	_, err = rt.Run(script)
	return err
}
