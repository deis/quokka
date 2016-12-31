thirdpartyresourcename = "mythirdpartyresource"
mythirdpartyresource = {
    "kind": "ThirdPartyResource",
    "apiVersion": "extensions/v1beta1",
    "metadata": {
        "name": thirdpartyresourcename,
        "labels": {
            "heritage": "Quokka",
        },
    },
    "versions": [
      {"name": "v1alpha1"},
    ]
};


res = kubernetes.thirdpartyresource.create(mythirdpartyresource)
if (res.metadata.name != thirdpartyresourcename) {
	throw "expected thirdpartyresource named " + thirdpartyresourcename
}

// Get our new thirdpartyresource by name
pp = kubernetes.thirdpartyresource.get(thirdpartyresourcename)
if (pp.metadata.name != thirdpartyresourcename) {
	throw "unexpected thirdpartyresource name: " + pp.metadata.name
}

// Search for our new thirdpartyresource.
matches = kubernetes.thirdpartyresource.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one thirdpartyresource in list"
}

// Update the thirdpartyresource
res.metadata.annotations = {"foo": "bar"}
res2 = kubernetes.thirdpartyresource.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

kubernetes.thirdpartyresource.delete(thirdpartyresourcename, {})
kubernetes.thirdpartyresource.deleteCollection({}, {labelSelector: "heritage=Quokka"})
