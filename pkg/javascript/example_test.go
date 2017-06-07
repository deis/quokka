package javascript_test

import (
	"fmt"
	"github.com/deis/quokka/pkg/javascript"
	"os"
)

func ExampleSetTimeout() {
	const testScript = `
		function TimeoutCallBack() {
			console.log("C");
		}
		console.log("A");
		setTimeout(TimeoutCallBack, 1000);
		console.log("B");
	`
	jsrt := javascript.NewRuntime()
	if _, err := jsrt.Run(testScript); err != nil {
		fmt.Fprintf(os.Stderr, "I'm sorry Dave, I'm afraid I can't do that: %v\n", err)
		os.Exit(2)
	}

	// Output:
	// A
	// B
	// C
}
