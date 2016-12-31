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

**Core:**

- `discovery`: API discovery
- `pod`: Pod management
- `namespace`: Namespace management
- `secret`: Secrets
- `configmap`: ConfigMaps
- `replicationcontroller`: Replication controllers
- `persistentvolumeclaim`: Persistent volume claims

**Apps:**

- `statefulset`: Stateful set (formerly PetSet)

**Extensions:**

- `daemonset`
- `deployment`
- `replicaset`
- `ingress`

## The Standard Methods

Many of the objects on `kubernetes` follow the same API. 

- create(objectDefinition): create a new object. For example, `kuberetes.pod.create` takes a Pod definition and creates a pod.
- delete(name, deleteOpts): delete an object
- list(listOpts): list objects. The `listOpts` object provides filters.
- get(name): Get an object by name
- update(objectDefinition): update an existing object
- updateStatus(objectDefinition): update an object status (expert)
- deleteCollection(deleteOpts, listOpts): delete a set of objects by `listOpts` filter.

```js
res = kubernetes.pods.list({labelSelector: "foo = bar"});
```


## The `kubernetes.discovery` object.

This object is quite a bit different than the others. This object provides access to API discovery mechanisms.

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

