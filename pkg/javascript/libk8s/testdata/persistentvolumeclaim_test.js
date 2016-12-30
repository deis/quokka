persistentvolumeclaimname = "mypersistentvolumeclaim"
mypersistentvolumeclaim = {
    "kind": "PersistentVolumeClaim",
    "apiVersion": "v1",
    "metadata": {
        "name": persistentvolumeclaimname,
        "namespace": "default",
        "labels": {
            "heritage": "Quokka",
        },
    },
    "spec": {
      "selector": { "app": "myapp" },
      "accessModes": ["ReadWriteOnce"],
      "resources":{ "requests": {"storage": "8Gi"}}
    },
};


res = kubernetes.persistentvolumeclaim.create(mypersistentvolumeclaim)
if (res.metadata.name != persistentvolumeclaimname) {
	throw "expected persistentvolumeclaim named " + persistentvolumeclaimname
}

// Get our new persistentvolumeclaim by name
pp = kubernetes.persistentvolumeclaim.get(persistentvolumeclaimname)
if (pp.metadata.name != persistentvolumeclaimname) {
	throw "unexpected persistentvolumeclaim name: " + pp.metadata.name
}

// Search for our new persistentvolumeclaim.
matches = kubernetes.persistentvolumeclaim.list({labelSelector: "heritage = Quokka"})
if (matches.items.length == 0) {
	throw "expected at least one persistentvolumeclaim in list"
}

// Update the persistentvolumeclaim
res.metadata.annotations = {"foo": "bar"}
res2 = kubernetes.persistentvolumeclaim.update(res)
if (res2.metadata.annotations.foo != "bar") {
	throw "expected foo annotation"
}

kubernetes.persistentvolumeclaim.delete(persistentvolumeclaimname, {})
