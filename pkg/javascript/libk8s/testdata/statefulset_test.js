statefulsetname = "mystatefulset"
mystatefulset = {
    "kind": "StatefulSet",
    "apiVersion": "apps/v1beta1",
    "metadata": {
        "name": statefulsetname,
        "namespace": "default",
        "labels": {
            "heritage": "Quokka",
        },
    },
    "spec": {
      "serviceName": "nginx",
      "replicas": 1,
      "selector": {"app": "nginx"},
      "template": {
        "metadata": {"name": "nginx"},
        "spec": {
          "containers": [
              {
                  "name": "waiter",
                  "image": "alpine:3.3",
                  "command": [
                      "/bin/sleep",
                      "9000"
                  ],
                  "imagePullPolicy": "IfNotPresent"
              }
          ]
        }
      }
    }
};


res = kubernetes.statefulset.create(mystatefulset)
if (res.metadata.name != statefulsetname) {
	throw "expected statefulset named " + statefulsetname
}

// Get our new statefulset by name
pp = kubernetes.statefulset.get(statefulsetname)
if (pp.metadata.name != statefulsetname) {
	throw "unexpected statefulset name: " + pp.metadata.name
}

// Search for our new statefulset.
matches = kubernetes.statefulset.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one statefulset in list"
}

// Update the statefulset
res.metadata.annotations = {"foo": "bar"}
res2 = kubernetes.statefulset.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

kubernetes.statefulset.delete(statefulsetname, {})
