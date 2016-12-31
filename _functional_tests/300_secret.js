console.log("====> secret test") // Oooo... mysterious.

secname = "mysecret"
ns = "quokkatest"
k = kubernetes.withNS(ns)
mysecret = {
    "kind": "Secret",
    "apiVersion": "v1",
    "metadata": {
        "name": secname,
        //"namespace": "quokkatest",
        "labels": {
            "heritage": "Quokka",
        },
    },
    "type": "Opaque",
    "data": {
        "username": "YWRtaW4="
    },
};

res = k.secret.create(mysecret)
if (res.metadata.name != secname) {
	throw "expected secret named " + secname
}

// Get our new secret by name
pp = k.secret.get(secname)
if (pp.metadata.name != secname) {
	throw "unexpected secret name: " + pp.metadata.name
}

// Search for our new secret.
matches = k.secret.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one secret in list"
}

// Update the secret
res.metadata.annotations = {"foo": "bar"}
res2 = k.secret.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

// Delete all the secrets
k.secret.deleteCollection({}, {labelSelector: "heritage = Quokka"})

// Verify delete
pp = k.secret.get(secname)
if (pp != undefined) {
	throw "expected secret to be deleted"
}
