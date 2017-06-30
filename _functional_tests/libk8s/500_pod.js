console.log("====> pod tests")

ns = "quokkatest"
k = kubernetes.withNS(ns).coreV1

podname = "mypod"
mypod = {
    "kind": "Pod",
    "apiVersion": "v1",
    "metadata": {
        "name": podname,
        //"namespace": ns,
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


res = k.pod.create(mypod)
if (res.metadata.name != podname) {
	throw "expected pod named " + podname
}

// TODO: Might need to sleep here.

// Get our new pod by name
pp = k.pod.get(podname)
if (pp.metadata.name != podname) {
	throw "unexpected pod name: " + pp.metadata.name
}

// Search for our new pod.
matches = k.pod.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one pod in list"
}

// Update the pod
// Need a sleep here to run these tests on Minikube. And then need to re-get.
// Right now, running tests on minikube basically causes a race condition
// because the pod takes several seconds to initialize.
/*
pp.metadata.annotations = {"foo": "bar"}
res2 = k.pod.update(pp)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}
*/

// TODO: Might need to sleep here.

k.pod.delete(podname, {})
