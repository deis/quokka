/* Package libk8s provides a JavaScript library for Kubernetes.

This exposes substantial portions of the Go API to JavaScript. Using this library,
scripts can interact with a Kubernetes API server using a structure very similar
to the native Go API for Kubernetes.

In many cases, we emulate JavaScript patterns in Go (such as the frequent use
of 'Get' and 'Set' in function names). We do this because it is more important
that the user-facing JavaScript runtime follow idiomatic JS than that the internal
API follows idiomatic Go.

The JavaScript API is basically a simplified version of the Client-Go API. We've
made the following simplifications:

- Namespace is bubbled to almost the top of the API so that it is set once and
  re-used internally. 'k = kubernetes.withNS("default"); k.coreV1.pods.get(...)'
- Discovery and namespace info is moved to the top of the API for convenience.
- Using annotations, we export the most stable version of a group without version
  numbers. For example `k.coreV1.pods` is the same as `k.core.pods`.
- Credential info is not passed into the runtime.
- Kubernetes objects (e.g. a pod definition) are plain JavaScript objects, not JSON,
  exact replcias of Go types, or special constructions. When passed into a Go function,
  they are converted into the native Go type by unmarshaling JSON.
- Wherever possible, we've converted functions to properties.

*/
package libk8s
