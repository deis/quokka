namespacename = "mynamespace"
mynamespace = {
    "kind": "Namespace",
    "apiVersion": "v1",
    "metadata": {
        "name": namespacename,
        "labels": {
            "heritage": "Quokka"
        }
    }
};


res = kubernetes.namespace.create(mynamespace)
if (res.metadata.name != namespacename) {
	throw "expected namespace named " + namespacename
}

// Get our new namespace by name
pp = kubernetes.namespace.get(namespacename)
if (pp.metadata.name != namespacename) {
	throw "unexpected namespace name: " + pp.metadata.name
}

// Search for our new namespace.
matches = kubernetes.namespace.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one namespace in list"
}

// Update the namespace
res.metadata.annotations = {"foo": "bar"}
res2 = kubernetes.namespace.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

kubernetes.namespace.delete(namespacename, {})
