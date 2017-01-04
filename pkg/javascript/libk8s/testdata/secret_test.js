secname = "mysecret"
mysecret = {
    "kind": "Secret",
    "apiVersion": "v1",
    "metadata": {
        "name": secname,
        "namespace": "default",
        "labels": {
            "heritage": "Quokka",
        },
    },
    "data": {
        "username": "YWRtaW4="
    },
};

secret = kubernetes.withNS("default").core.secret

res = secret.create(mysecret)
if (res.metadata.name != secname) {
 throw "expected secret named " + secname
}

// Get our new secret by name
pp = secret.get(secname)
if (pp.metadata.name != secname) {
  throw "unexpected secret name: " + pp.metadata.name
}

// Search for our new secret.
matches = secret.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
  throw "expected at least one secret in list"
}

// Update the secret
res.metadata.annotations = {"foo": "bar"}
res2 = secret.update(res)
if (res2.metadata.annotations.foo != "bar") {
  throw "expected foo annotation"
}

// Delete all the secrets
secret.deleteCollection({}, {labelSelector: "heritage = Quokka"})
