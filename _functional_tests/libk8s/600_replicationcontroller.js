console.log("====> replicationcontroller test")

replicationcontrollername = "myreplicationcontroller"
myreplicationcontroller = {
    "kind": "ReplicationController",
    "apiVersion": "v1",
    "metadata": {
        "name": replicationcontrollername,
        "namespace": "quokkatest",
        "labels": {
            "heritage": "Quokka"
        },
    },
    "spec": {
      "replicas": 1,
      // The API for this changed, but is undocumented.
      "selector": {"app":"nginx"},
      "template": {
        "metadata": {
          "name": "nginx"
          "labels": {"app":"nginx"},
        },
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

RC = kubernetes.withNS("quokkatest").coreV1.replicationcontroller


res = RC.create(myreplicationcontroller)
if (res.metadata.name != replicationcontrollername) {
	throw "expected replicationcontroller named " + replicationcontrollername
}

// Get our new replicationcontroller by name
pp = RC.get(replicationcontrollername)
if (pp.metadata.name != replicationcontrollername) {
	throw "unexpected replicationcontroller name: " + pp.metadata.name
}

// Search for our new replicationcontroller.
matches = RC.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one replicationcontroller in list"
}

// Update the replicationcontroller
res.metadata.annotations = {"foo": "bar"}
res2 = RC.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

RC.delete(replicationcontrollername, {})
