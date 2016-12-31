package libk8s

import (
	"github.com/robertkrimen/otto"
	"github.com/technosophos/ottomatic"

	"k8s.io/client-go/kubernetes"
)

// Kubernetes is the top-level object exposed to JavaScript.
//
// It provides access to the latest supported versions of the built-in API, along
// with partiall tooling for ThirdPartyResources.
//
// Not all of the API is exposed. Emphasis has been given to commonly-used APIs
// that follow a consistent pattern.
//
// In many cases, the APIs function by executing Go functions and then marshaling
// the results to their JSON equivalent, which can then be loaded into the
// JavaScript runtime as-is.
//
// Objects created in JavaScript are passed into Go as maps, and are then
// typed by unmarshaling them into their respective JSON struct.
type Kubernetes struct {
	// WithNS initializes another Kubernetes object, but with a different namespace.
	WithNS func(string) otto.Value `otto:"withNS"`
	GetNS  func() otto.Value       `otto:"getNS"`

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

	// Batch

	Job *Job `otto:"job"`
}

func NewKubernetes(vm *otto.Otto, c kubernetes.Interface, ns string) *Kubernetes {
	k := &Kubernetes{
		WithNS: func(ns string) otto.Value {
			kk := NewKubernetes(vm, c, ns)
			obj, err := vm.Object(`ns = {}`)
			poe(err)
			err = ottomatic.RegisterTo("x", kk, vm, obj)
			poe(err)
			res, err := obj.Get("x")
			poe(err)
			return res
		},
		GetNS: func() otto.Value { r, _ := vm.ToValue(ns); return r },

		Namespace:             NewNamespace(c, vm),
		Discovery:             NewDiscovery(c, vm),
		Pod:                   NewPod(c, vm, ns),
		Secret:                NewSecret(c, vm, ns),
		ConfigMap:             NewConfigMap(c, vm, ns),
		Service:               NewService(c, vm, ns),
		ServiceAccount:        NewServiceAccount(c, vm, ns),
		PersistentVolumeClaim: NewPersistentVolumeClaim(c, vm, ns),
		ReplicationController: NewReplicationController(c, vm, ns),
		StatefulSet:           NewStatefulSet(c, vm, ns),
		Deployment:            NewDeployment(c, vm, ns),
		ReplicaSet:            NewReplicaSet(c, vm, ns),
		DaemonSet:             NewDaemonSet(c, vm, ns),
		Ingress:               NewIngress(c, vm, ns),
		PodSecurityPolicy:     NewPodSecurityPolicy(c, vm),
		ThirdPartyResource:    NewThirdPartyResource(c, vm),
		Job:                   NewJob(c, vm, ns),
	}
	return k
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
	k := NewKubernetes(vm, c, "default")
	return ottomatic.Register("kubernetes", k, vm)
}
