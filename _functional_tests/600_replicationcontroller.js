console.log("====> replicationcontroller test")

replicationcontrollername = "myreplicationcontroller"
myreplicationcontroller = {
    "kind": "ReplicationController",
    "apiVersion": "v1",
    "metadata": {
        "name": replicationcontrollername,
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


res = kubernetes.replicationcontroller.create(myreplicationcontroller)
if (res.metadata.name != replicationcontrollername) {
	throw "expected replicationcontroller named " + replicationcontrollername
}

// Get our new replicationcontroller by name
pp = kubernetes.replicationcontroller.get(replicationcontrollername)
if (pp.metadata.name != replicationcontrollername) {
	throw "unexpected replicationcontroller name: " + pp.metadata.name
}

// Search for our new replicationcontroller.
matches = kubernetes.replicationcontroller.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one replicationcontroller in list"
}

// Update the replicationcontroller
res.metadata.annotations = {"foo": "bar"}
res2 = kubernetes.replicationcontroller.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

kubernetes.replicationcontroller.delete(replicationcontrollername, {})
