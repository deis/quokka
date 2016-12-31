package libk8s

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/robertkrimen/otto"
	"github.com/technosophos/ottomatic"

	"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/1.4/pkg/api/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Kubernetes is the top-level object exposed to JavaScript.
// It consisists of two things: A high-level API to work wtih Kubernetes and
// low-level access to the Go Kubernetes API.
//
// The high-level API is designed to feel native to JavaScript. It follows the
// common JavaScript patterns and conventions.
//
// Access to the low-level API is granted with knowledge that it will not be easy
// to use or even possible to use. But it offers dedicated and knowledgeable
// developers the opportunity to experiment. It may be deprecated prior to
// the initial release.
type Kubernetes struct {
	// Impl provides the JavaScript runtime with access to the raw implementation.
	//
	// This is for expert use only.
	//Impl kubernetes.Interface `otto:"impl"`

	// Core
	Namespace             *Namespace             `otto:"namespace"`
	Discovery             *Discovery             `otto:"discovery"`
	Pod                   *Pod                   `otto:"pod"`
	Secret                *Secret                `otto:"secret"`
	ConfigMap             *ConfigMap             `otto:"configmap"`
	Service               *Service               `otto:"service"`
	ServiceAccount        *ServiceAccount        `otto:"serviceaccount"`
	PersistentVolumeClaim *PersistentVolumeClaim `otto:"persistentvolumeclaim"`
	ReplicationController *ReplicationController `otto:"replicationcontroller"`

	// Apps
	StatefulSet *StatefulSet `otto:"statefulset"`

	//Extensions
	Deployment         *Deployment         `otto:"deployment"`
	ReplicaSet         *ReplicaSet         `otto:"replicaset"`
	DaemonSet          *DaemonSet          `otto:"daemonset"`
	Ingress            *Ingress            `otto:"ingress"`
	PodSecurityPolicy  *PodSecurityPolicy  `otto:"podsecuritypolicy"`
	ThirdPartyResource *ThirdPartyResource `otto:"thirdpartyresource"`
}

// Register registers the top-level Kubernetes API objects with the JS runtime.
func Register(vm *otto.Otto) error {
	c, err := kubeClient()
	if err != nil {
		return err
	}
	return RegisterWithClient(vm, c)
}

// RegisterWithClient registers the Kubernetes object to an existing Kubernetes client.
func RegisterWithClient(vm *otto.Otto, c kubernetes.Interface) error {
	k := &Kubernetes{
		Namespace:             NewNamespace(c, vm),
		Discovery:             NewDiscovery(c, vm),
		Pod:                   NewPod(c, vm),
		Secret:                NewSecret(c, vm),
		ConfigMap:             NewConfigMap(c, vm),
		Service:               NewService(c, vm),
		ServiceAccount:        NewServiceAccount(c, vm),
		PersistentVolumeClaim: NewPersistentVolumeClaim(c, vm),
		ReplicationController: NewReplicationController(c, vm),
		StatefulSet:           NewStatefulSet(c, vm),
		Deployment:            NewDeployment(c, vm),
		ReplicaSet:            NewReplicaSet(c, vm),
		DaemonSet:             NewDaemonSet(c, vm),
		Ingress:               NewIngress(c, vm),
		PodSecurityPolicy:     NewPodSecurityPolicy(c, vm),
		ThirdPartyResource:    NewThirdPartyResource(c, vm),
	}
	return ottomatic.Register("kubernetes", k, vm)
}

func kubeConfig() (*rest.Config, error) {
	// Try in-cluster config:
	c, err := rest.InClusterConfig()
	if err == nil {
		log.Print("in-cluster")
		return c, nil
	}

	// TODO: Env vars support colon-separated list (and kubectl allows this).
	// Figure out how to do this with client-go.
	kconf := os.Getenv("KUBECONFIG")
	parts := strings.Split(kconf, ":")
	kconf = parts[0]
	if len(parts) > 0 {
		log.Printf("WARNING: Building config only from %q", kconf)
	}

	cfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kconf},
		&clientcmd.ConfigOverrides{})

	//rc, _ := cfg.RawConfig()
	//fmt.Printf("Use context: %s\n", rc.CurrentContext)

	return cfg.ClientConfig()
}

func kubeClient() (kubernetes.Interface, error) {
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

// ToObject marshals an interface to JSON, and then creates a JavaScript object.
// The object is then returned as a value.
func ToObject(v interface{}, o *otto.Otto) (otto.Value, error) {
	// Let Go buikd the JavaScript object for us.
	j, err := json.Marshal(v)
	if err != nil {
		return otto.UndefinedValue(), err
	}

	log.Printf("Out: %s", trunc(string(j), 1024))

	// Note that the variable name is not returned with obj.Value --
	// just the rval is.
	obj, err := o.Object("v = " + string(j))
	if err != nil {
		return otto.UndefinedValue(), err
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

// poe is panic-on-error
//
// In the JavaScript runtime, throw/catch is implemented using panic. So
// to bubble an error to a try/catch, we need to panic.
func poe(err error) {
	if err != nil {
		panic(err)
	}
}

// remarshal turns src into JSON, and then unmarshals it into dest.
//
// This is a cheap and easy way to take generic data from JavaScript and convert
// it to strongly typed data in Go.
func remarshal(src, dest interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}
