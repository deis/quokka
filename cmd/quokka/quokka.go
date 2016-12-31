package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/deis/quokka/pkg/javascript"
	"github.com/deis/quokka/pkg/javascript/libk8s"
	"github.com/spf13/cobra"
)

const usage = `Execute JavaScript files inside of the Quokka runtime.

Quokka is a Kubernetes client that can execute arbitrary JavaScript.

A Quokka script will have access to the following libraries:

  - Underscore.js
  - The 'kubernetes' object, which has access to a variety of features, including:
    * 'pod': Pod functions
    * 'configmap': ConfigMap functions
    * 'secret': Access to Kubernetes secrets
    * and many more

Quokka is not a drop-in replacement for Node.js. It is a special-purpose client
for working with Kubernetes.
`

func main() {
	cmd := &cobra.Command{
		Use:   "quokka [flags] file [file]...",
		Short: "execute one or more JavaScript files inside the Quokka runtime",
		Long:  usage,
		RunE:  run,
	}
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("at least one filename must be specified")
	}
	rt := javascript.NewRuntime()
	libk8s.Register(rt.VM)

	for _, file := range args {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		if _, err = rt.VM.Run(data); err != nil {
			return err
		}
	}

	return nil
}
