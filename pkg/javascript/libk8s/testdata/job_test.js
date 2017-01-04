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

batch = kubernetes.withNS("default").batch


res = batch.job.create(myjob)
if (res.metadata.name != jobname) {
	throw "expected job named " + jobname
}

// Get our new job by name
pp = batch.job.get(jobname)
if (pp.metadata.name != jobname) {
	throw "unexpected job name: " + pp.metadata.name
}

// Search for our new job.
matches = batch.job.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one job in list"
}

// Update the job
res.metadata.annotations = {"foo": "bar"}
res2 = batch.job.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

batch.job.delete(jobname, {})
batch.job.deleteCollection({}, {labelSelector: "heritage=Quokka"})
