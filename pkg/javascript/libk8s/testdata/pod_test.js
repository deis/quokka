podname = "mypod"
mypod = {
    "kind": "Pod",
    "apiVersion": "v1",
    "metadata": {
        "name": podname,
        "namespace": "default",
        "labels": {
            "heritage": "Quokka",
        },
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
        ],
    },
};


res = kubernetes.pod.create(mypod)
if (res.metadata.name != podname) {
	throw "expected pod named " + podname
}

// Get our new pod by name
pp = kubernetes.pod.get(podname)
if (pp.metadata.name != podname) {
	throw "unexpected pod name: " + pp.metadata.name
}

// Search for our new pod.
matches = kubernetes.pod.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one pod in list"
}

// Update the pod
res.metadata.annotations = {"foo": "bar"}
res2 = kubernetes.pod.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

kubernetes.pod.delete(podname, {})
