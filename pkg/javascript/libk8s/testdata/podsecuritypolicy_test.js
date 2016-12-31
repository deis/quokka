podsecuritypolicyname = "mypodsecuritypolicy"
mypodsecuritypolicy = {
    "kind": "PodSecurityPolicy",
    "apiVersion": "extensions/v1beta1",
    "metadata": {
        "name": podsecuritypolicyname,
        "labels": {
            "heritage": "Quokka",
        },
    },
    "spec": { }
};


res = kubernetes.podsecuritypolicy.create(mypodsecuritypolicy)
if (res.metadata.name != podsecuritypolicyname) {
	throw "expected podsecuritypolicy named " + podsecuritypolicyname
}

// Get our new podsecuritypolicy by name
pp = kubernetes.podsecuritypolicy.get(podsecuritypolicyname)
if (pp.metadata.name != podsecuritypolicyname) {
	throw "unexpected podsecuritypolicy name: " + pp.metadata.name
}

// Search for our new podsecuritypolicy.
matches = kubernetes.podsecuritypolicy.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one podsecuritypolicy in list"
}

// Update the podsecuritypolicy
res.metadata.annotations = {"foo": "bar"}
res2 = kubernetes.podsecuritypolicy.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

kubernetes.podsecuritypolicy.delete(podsecuritypolicyname, {})
kubernetes.podsecuritypolicy.deleteCollection({}, {labelSelector: "heritage=Quokka"})
