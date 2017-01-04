daemonsetname = "mydaemonset"
mydaemonset = {
    "kind": "DaemonSet",
    "apiVersion": "extensions/v1beta1",
    "metadata": {
        "name": daemonsetname,
        "namespace": "default",
        "labels": {
            "heritage": "Quokka",
        },
    },
    "spec": {
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

ext = kubernetes.withNS("default").extensions


res = ext.daemonset.create(mydaemonset)
if (res.metadata.name != daemonsetname) {
	throw "expected daemonset named " + daemonsetname
}

// Get our new daemonset by name
pp = ext.daemonset.get(daemonsetname)
if (pp.metadata.name != daemonsetname) {
	throw "unexpected daemonset name: " + pp.metadata.name
}

// Search for our new daemonset.
matches = ext.daemonset.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one daemonset in list"
}

// Update the daemonset
res.metadata.annotations = {"foo": "bar"}
res2 = ext.daemonset.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

ext.daemonset.delete(daemonsetname, {})
ext.daemonset.deleteCollection({}, {labelSelector: "heritage=Quokka"})
