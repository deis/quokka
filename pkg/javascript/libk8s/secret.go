package libk8s

import (
	"github.com/robertkrimen/otto"

	"k8s.io/client-go/kubernetes"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/pkg/api/v1"
)

type Secret struct {
	Create           func(map[string]interface{}) otto.Value `otto:"create"`
	Delete           func(string, map[string]interface{})    `otto:"delete"`
	List             func(map[string]interface{}) otto.Value `otto:"list"`
	Get              func(string) otto.Value                 `otto:"get"`
	Update           func(map[string]interface{}) otto.Value `otto:"update"`
	DeleteCollection func(dopt, lopt map[string]interface{}) `otto:"deleteCollection"`

	// Watch func() otto.Value `otto:"watch"`
	// Patch func() otto.Value `otto:"patch"`

	// We do not currently support secret expansion
}

func NewSecret(c kubernetes.Interface, o *otto.Otto, ns string) *Secret {
	iface := func() v1core.SecretInterface {
		return c.CoreV1().Secrets(ns)
	}
	return &Secret{
		Create: func(sec map[string]interface{}) otto.Value {
			gp := &v1.Secret{}
			poe(remarshal(sec, gp))
			p, err := iface().Create(gp)
			poe(err)
			return MustToObject(p, o)
		},
		Delete: func(name string, opts map[string]interface{}) {
			do := &v1.DeleteOptions{}
			poe(remarshal(opts, do))
			err := iface().Delete(name, do)
			poe(err)
		},
		List: func(opts map[string]interface{}) otto.Value {
			lo := v1.ListOptions{}
			poe(remarshal(opts, &lo))
			pl, err := iface().List(lo)
			poe(err)
			return MustToObject(pl, o)
		},
		Get: func(name string) otto.Value {
			out, err := iface().Get(name)
			poe(err)
			return MustToObject(out, o)
		},
		Update: func(sec map[string]interface{}) otto.Value {
			gp := &v1.Secret{}
			poe(remarshal(sec, gp))
			p, err := iface().Update(gp)
			poe(err)
			return MustToObject(p, o)
		},
		DeleteCollection: func(dopts, lopts map[string]interface{}) {
			do := &v1.DeleteOptions{}
			poe(remarshal(dopts, do))
			lo := v1.ListOptions{}
			poe(remarshal(lopts, &lo))
			err := iface().DeleteCollection(do, lo)
			poe(err)
		},
	}
}
