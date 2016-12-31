ingressname = "myingress"
myingress = {
    "kind": "Ingress",
    "apiVersion": "extensions/v1beta1",
    "metadata": {
        "name": ingressname,
        "namespace": "default",
        "labels": {
            "heritage": "Quokka",
        },
    },
    "spec": {
      "rules": [
        {
          "http": {
            "paths": []
          }
        }
      ]
    }
};


res = kubernetes.ingress.create(myingress)
if (res.metadata.name != ingressname) {
	throw "expected ingress named " + ingressname
}

// Get our new ingress by name
pp = kubernetes.ingress.get(ingressname)
if (pp.metadata.name != ingressname) {
	throw "unexpected ingress name: " + pp.metadata.name
}

// Search for our new ingress.
matches = kubernetes.ingress.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one ingress in list"
}

// Update the ingress
res.metadata.annotations = {"foo": "bar"}
res2 = kubernetes.ingress.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

kubernetes.ingress.delete(ingressname, {})
kubernetes.ingress.deleteCollection({}, {labelSelector: "heritage=Quokka"})
