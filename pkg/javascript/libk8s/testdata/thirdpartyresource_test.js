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

tpr = kubernetes.withNS("default").extensions.thirdpartyresource

res = tpr.create(mythirdpartyresource)
if (res.metadata.name != thirdpartyresourcename) {
	throw "expected thirdpartyresource named " + thirdpartyresourcename
}

// Get our new thirdpartyresource by name
pp = tpr.get(thirdpartyresourcename)
if (pp.metadata.name != thirdpartyresourcename) {
	throw "unexpected thirdpartyresource name: " + pp.metadata.name
}

// Search for our new thirdpartyresource.
matches = tpr.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one thirdpartyresource in list"
}

// Update the thirdpartyresource
res.metadata.annotations = {"foo": "bar"}
res2 = tpr.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

tpr.delete(thirdpartyresourcename, {})
tpr.deleteCollection({}, {labelSelector: "heritage=Quokka"})
