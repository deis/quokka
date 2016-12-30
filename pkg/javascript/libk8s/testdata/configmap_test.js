cmname = "myconfigmap"
myconfigmap = {
    "kind": "ConfigMap",
    "apiVersion": "v1",
    "metadata": {
        "name": cmname,
        "namespace": "default",
        "labels": {
            "heritage": "Quokka",
        },
    },
    "data": {
        "username": "hello"
    },
};


res = kubernetes.configmap.create(myconfigmap)
if (res.metadata.name != cmname) {
 throw "expected configmap named " + cmname
}

// Get our new configmap by name
pp = kubernetes.configmap.get(cmname)
if (pp.metadata.name != cmname) {
  throw "unexpected configmap name: " + pp.metadata.name
}

// Search for our new configmap.
matches = kubernetes.configmap.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
  throw "expected at least one configmap in list"
}

// Update the configmap
res.metadata.annotations = {"foo": "bar"}
res2 = kubernetes.configmap.update(res)
if (res2.metadata.annotations.foo != "bar") {
  throw "expected foo annotation"
}

kubernetes.configmap.delete(cmname, {})
