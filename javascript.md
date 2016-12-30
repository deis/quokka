# JavaScript API

Quokka provides a JavaScript API to Kubernetes.

You can see this API in action by looking at the `_functional_tests/*.js` files.

## About Quokka's JS API

_Why not do this as a Node module?_ We have no doubt that at some point a Node.js
module will be built for Kubernetes. However, this project is trying to accomplish
something a little different: We want to expose a JavaScript engine inside of
Kubernetes' Go runtime. This allows us to:

- Re-use massive amounts of tested and type-safe code
- Use JavaScript to embed logic in Go at runtime without recompilation

To that end, we took the Otto JavaScript implementation and the `client-go`
Kubernetes library and combined them.

## Basic Usage

Quokka ships with:

- An ES5 JavaScript engine
- The `underscore.js` library
- The `kubernetes` library

## Objects and Functions

This section documents all of the objects that are currently available. Note
that this library may not expose _all_ of the core Kubernetes types in the near
future. Quokka will implement the most frequently used Kubernetes kinds first,
then add more as time/contributions allow.

## The `kubernetes` object.

The top-level object for accessing Kubernetes is the `kubernetes` object.

The `kubernetes` object has the following objects:

- `discovery`: API discovery
- `pod`: Pod management
- `namespace`: Namespace management

## The `kubernetes.discovery` object.

This object provides access to API discovery mechanisms.

```js
ver = kubernetes.discovery.serverVersion()
console.log(ver.gitVersion)
```

### Methods

- `serverVersion()`: Information about the server version
- `serverGroups()`: Information about what API groups are supported.
- `swaggerSchema(group, version)`: Access to the Swagger API schema by group and version (`batch`, `v1`)
- `serverResources()`: Server resource definitions
- `serverResourcesForGroupVersion(groupVersion)`: Server resource definitions by `group/version`: `batch/v1`
- `serverPreferredResources`: Preferred server resources.
- `serverPreferredNamespacedResources`: Namespaced preferred server resources.

## The `kubernetes.pod` object

This object provides access to Kubernetes' pods API.

```js
res = kubernetes.pods.list({labelSelector: "foo = bar"});
```

- create(podDef)
- delete(name, deleteOpts)
- list(listOpts)
- get(name)
- update(podDef)
- updateStatus(podDef)
- deleteCollection(deleteOpts, listOpts)
