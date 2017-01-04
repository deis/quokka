serviceaccountname = "myserviceaccount"
myserviceaccount = {
    "kind": "ServiceAccount",
    "apiVersion": "v1",
    "metadata": {
        "name": serviceaccountname,
        "namespace": "default",
        "labels": {
            "heritage": "Quokka",
        }
    }
};

sa = kubernetes.withNS("default").core.serviceaccount


res = sa.create(myserviceaccount)
if (res.metadata.name != serviceaccountname) {
	throw "expected serviceaccount named " + serviceaccountname
}

// Get our new serviceaccount by name
pp = sa.get(serviceaccountname)
if (pp.metadata.name != serviceaccountname) {
	throw "unexpected serviceaccount name: " + pp.metadata.name
}

// Search for our new serviceaccount.
matches = sa.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one serviceaccount in list"
}

// Update the serviceaccount
res.metadata.annotations = {"foo": "bar"}
res2 = sa.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

sa.delete(serviceaccountname, {})
