console.log("====> configmap test")
cmname = "myconfigmap"
myns = "quokkatest"
//myns = "default"
myconfigmap = {
    "kind": "ConfigMap",
    "apiVersion": "v1",
    "metadata": {
        "name": cmname,
        "namespace": myns,
        "labels": {
            "heritage": "Quokka",
        },
    },
    "data": {
        "username": "hello"
    },
};
k = kubernetes.withNS(myns).extensions

// Clean up after any failed tests
k.configmap.deleteCollection({}, {labelSelector: "heritage=Quokka"})


res = k.configmap.create(myconfigmap)
if (res.metadata.name != cmname) {
 throw "expected configmap named " + cmname
}

// Get our new configmap by name
pp = k.configmap.get(cmname)
if (pp.metadata.name != cmname) {
  throw "unexpected configmap name: " + pp.metadata.name
}

// Search for our new configmap.
matches = k.configmap.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
  throw "expected at least one configmap in list"
}

// Update the configmap
res.metadata.annotations = {"foo": "bar"}
res2 = k.configmap.update(res)
if (res2.metadata.annotations.foo != "bar") {
  throw "expected foo annotation"
}

k.configmap.delete(cmname, {})
