package libk8s

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/robertkrimen/otto"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/deis/quokka/pkg/javascript/jsutil"
)

type Runtime struct {
	DefaultNamespace string `otto:"defaultNamespace"`

	o *otto.Otto `otto:"-"`
}

// KubernetesErrorType is the string name of the JavaScript error type.
const KubernetesErrorType = "KubernetesError"

// POE panics on error (in a way that is try/catchable to JS)
//
// If the given error is non-nil, this panics. THe panic value is a
// custom JavaScript type (KubernetesError)
func (r *Runtime) POE(err error) {
	if err != nil {
		panic(r.o.MakeCustomError(KubernetesErrorType, err.Error()))
	}
}

// ToObject marshals an interface to JSON, and then creates a JavaScript object.
// The object is then returned as a value.
func (r *Runtime) ToObject(v interface{}) (otto.Value, error) {
	// Let Go build the JavaScript object for us.
	j, err := json.Marshal(v)
	if err != nil {
		return otto.UndefinedValue(), fmt.Errorf("ToObject marshal %T: %s", v, err)
	}

	log.Printf("Out: %s", trunc(string(j), 1024))

	// Note that the variable name is not returned with obj.Value --
	// just the rval is.
	obj, err := r.o.Object("v = " + string(j))
	if err != nil {
		return otto.UndefinedValue(), fmt.Errorf("ToObject object %T: %s", v, err)
	}
	return obj.Value(), nil
}

// MustToObject runs ToObject, and panics on failure.
func (r *Runtime) MustToObject(v interface{}) otto.Value {
	val, e := r.ToObject(v)
	r.POE(e)
	return val
}

// MustRemarshal remarshals the src interface into the dest interface.
//
// It does this via the built-in JSON marshal.
//
// The intent of this function is to leverage the JSON marshal to convert
// between untyped native JS objects and typed Go Kubernetes API objects.
func (r *Runtime) MustRemarshal(src, dest interface{}) {
	r.POE(remarshal(src, dest))
}

func kubeConfig() (*rest.Config, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	rules.DefaultClientConfig = &clientcmd.DefaultClientConfig

	overrides := &clientcmd.ConfigOverrides{ClusterDefaults: clientcmd.ClusterDefaults}

	cfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, overrides)

	return cfg.ClientConfig()
}

// KubeClient returns an initialized Kubernetes client.
//
// Normally, it is best to use Register() instead.
//
// This will attempt to load a KUBECONFIG from a variety of well-known locations
// and formats. Then it will initialize a client with the found configuration.
func KubeClient() (kubernetes.Interface, error) {
	cfg, err := kubeConfig()
	if err != nil {
		return nil, err
	}

	// This merely ensures that the proxy can auth.
	//log.Printf("Connecting to %q", cli.Api)
	cli, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Printf("cannot get new Kube client: %s", err)
		return cli, err
		//} else if _, err := cli.Namespaces().List(v1.ListOptions{}); err != nil {
		//	log.Printf("cannot get namespaces: %s", err)
		//	return cli, err
	}
	return cli, nil
}

func trunc(s string, m int) string {
	if len(s) < m {
		return s
	}
	return s[:m] + "..."
}

// EVERYTHING BELOW THIS LINE IS DEPRECATED AND SHOULD BE REMOVED.

// ToObject marshals an interface to JSON, and then creates a JavaScript object.
// The object is then returned as a value.
func ToObject(v interface{}, o *otto.Otto) (otto.Value, error) {
	// Let Go build the JavaScript object for us.
	j, err := json.Marshal(v)
	if err != nil {
		return otto.UndefinedValue(), fmt.Errorf("ToObject marshal %T: %s", v, err)
	}

	log.Printf("Out: %s", trunc(string(j), 1024))

	// Note that the variable name is not returned with obj.Value --
	// just the rval is.
	obj, err := o.Object("v = " + string(j))
	if err != nil {
		return otto.UndefinedValue(), fmt.Errorf("ToObject object %T: %s", v, err)
	}
	return obj.Value(), nil
}

// MustToObject runs ToObject, and panics on failure.
func MustToObject(v interface{}, o *otto.Otto) otto.Value {
	val, e := ToObject(v, o)
	if e != nil {
		panic(e)
	}
	return val
}

// deprecated
func poe(e error)                      { jsutil.POE(e) }
func remarshal(a, b interface{}) error { return jsutil.Remarshal(a, b) }
