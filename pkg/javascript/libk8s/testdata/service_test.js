servicename = "myservice"
myservice = {
    "kind": "Service",
    "apiVersion": "v1",
    "metadata": {
        "name": servicename,
        "namespace": "default",
        "labels": {
            "heritage": "Quokka",
        },
    },
    "spec": {
      "selector": { "app": "myapp" },
      "ports": [
        {"protocol": "TCP", "port": 8080}
      ]
    },
};


res = kubernetes.service.create(myservice)
if (res.metadata.name != servicename) {
	throw "expected service named " + servicename
}

// Get our new service by name
pp = kubernetes.service.get(servicename)
if (pp.metadata.name != servicename) {
	throw "unexpected service name: " + pp.metadata.name
}

// Search for our new service.
matches = kubernetes.service.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one service in list"
}

// Update the service
res.metadata.annotations = {"foo": "bar"}
res2 = kubernetes.service.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

kubernetes.service.delete(servicename, {})
