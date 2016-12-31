jobname = "myjob"
myjob = {
    "kind": "Job",
    "apiVersion": "batch/v1",
    "metadata": {
        "name": jobname,
        "namespace": "default",
        "labels": {
            "heritage": "Quokka",
        },
    },
    "spec": {
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


res = kubernetes.job.create(myjob)
if (res.metadata.name != jobname) {
	throw "expected job named " + jobname
}

// Get our new job by name
pp = kubernetes.job.get(jobname)
if (pp.metadata.name != jobname) {
	throw "unexpected job name: " + pp.metadata.name
}

// Search for our new job.
matches = kubernetes.job.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one job in list"
}

// Update the job
res.metadata.annotations = {"foo": "bar"}
res2 = kubernetes.job.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

kubernetes.job.delete(jobname, {})
kubernetes.job.deleteCollection({}, {labelSelector: "heritage=Quokka"})
