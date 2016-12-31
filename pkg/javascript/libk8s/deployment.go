package libk8s

import (
	"github.com/robertkrimen/otto"

	"k8s.io/client-go/kubernetes"
	v1beta1extensions "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	v1 "k8s.io/client-go/pkg/api/v1"
	extensions "k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

type Deployment struct {
	Create           func(map[string]interface{}) otto.Value `otto:"create"`
	Delete           func(string, map[string]interface{})    `otto:"delete"`
	List             func(map[string]interface{}) otto.Value `otto:"list"`
	Get              func(string) otto.Value                 `otto:"get"`
	Update           func(map[string]interface{}) otto.Value `otto:"update"`
	UpdateStatus     func(map[string]interface{}) otto.Value `otto:"updateStatus"`
	DeleteCollection func(dopt, lopt map[string]interface{}) `otto:"deleteCollection"`

	// Watch func() otto.Value `otto:"watch"`
	// Patch func() otto.Value `otto:"patch"`
}

func NewDeployment(c kubernetes.Interface, o *otto.Otto, ns string) *Deployment {
	pi := func() v1beta1extensions.DeploymentInterface {
		return c.ExtensionsV1beta1().Deployments(ns)
	}
	return &Deployment{
		Create: func(pod map[string]interface{}) otto.Value {
			gp := &extensions.Deployment{}
			poe(remarshal(pod, gp))
			p, err := pi().Create(gp)
			poe(err)
			return MustToObject(p, o)
		},
		Delete: func(name string, opts map[string]interface{}) {
			do := &v1.DeleteOptions{}
			poe(remarshal(opts, do))
			err := pi().Delete(name, do)
			poe(err)
		},
		List: func(opts map[string]interface{}) otto.Value {
			lo := v1.ListOptions{}
			poe(remarshal(opts, &lo))
			pl, err := pi().List(lo)
			poe(err)
			return MustToObject(pl, o)
		},
		Get: func(name string) otto.Value {
			out, err := pi().Get(name)
			poe(err)
			return MustToObject(out, o)
		},
		Update: func(pod map[string]interface{}) otto.Value {
			gp := &extensions.Deployment{}
			poe(remarshal(pod, gp))
			p, err := pi().Update(gp)
			poe(err)
			return MustToObject(p, o)
		},
		UpdateStatus: func(pod map[string]interface{}) otto.Value {
			gp := &extensions.Deployment{}
			poe(remarshal(pod, gp))
			p, err := pi().UpdateStatus(gp)
			poe(err)
			return MustToObject(p, o)
		},
		DeleteCollection: func(dopts, lopts map[string]interface{}) {
			do := &v1.DeleteOptions{}
			poe(remarshal(dopts, do))
			lo := v1.ListOptions{}
			poe(remarshal(lopts, &lo))
			err := pi().DeleteCollection(do, lo)
			poe(err)
		},
	}
}
