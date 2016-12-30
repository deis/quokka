package libk8s

import (
	"fmt"

	"github.com/robertkrimen/otto"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/unversioned"
)

// Arity0 is a JavaScript function with arity 0
type Arity0 func() otto.Value

type Discovery struct {
	ServerVersion                      Arity0                          `otto:"serverVersion"`
	ServerGroups                       Arity0                          `otto:"serverGroups"`
	SwaggerSchema                      func(string, string) otto.Value `otto:"swaggerSchema"`
	ServerResourcesForGroupVersion     func(string) otto.Value         `otto:"serverResourcesForGroupVersion"`
	ServerResources                    Arity0                          `otto:"serverResources"`
	ServerPreferredNamespacedResources Arity0                          `otto:"serverPreferredNamespacedResources"`
	ServerPreferredResources           Arity0                          `otto:"serverPreferredResources"`
}

func NewDiscovery(c kubernetes.Interface, o *otto.Otto) *Discovery {
	return &Discovery{
		ServerVersion: func() otto.Value {
			vi, err := c.Discovery().ServerVersion()
			poe(err)
			return MustToObject(vi, o)
		},
		ServerGroups: func() otto.Value {
			g, err := c.Discovery().ServerGroups()
			poe(err)
			fmt.Printf("Groups: %v", g)
			return MustToObject(g, o)
		},
		SwaggerSchema: func(group string, version string) otto.Value {
			//gv := schema.GroupVersion{Group: group, Version: version},
			gv := unversioned.GroupVersion{Group: group, Version: version}
			g, err := c.Discovery().SwaggerSchema(gv)
			poe(err)
			return MustToObject(g, o)
		},
		ServerResourcesForGroupVersion: func(gv string) otto.Value {
			g, err := c.Discovery().ServerResourcesForGroupVersion(gv)
			poe(err)
			return MustToObject(g, o)
		},
		ServerResources: func() otto.Value {
			g, err := c.Discovery().ServerResources()
			poe(err)
			return MustToObject(g, o)
		},
		ServerPreferredResources: func() otto.Value {
			g, err := c.Discovery().ServerPreferredResources()
			poe(err)
			return MustToObject(g, o)
		},
		ServerPreferredNamespacedResources: func() otto.Value {
			g, err := c.Discovery().ServerPreferredNamespacedResources()
			poe(err)
			return MustToObject(g, o)
		},
	}
}
