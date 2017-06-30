console.log("====> serverResources tests")

v = kubernetes.discovery.serverResources()
if (v.length == 0) {
  throw "Expected at least one server resource"
}

v = kubernetes.discovery.serverPreferredResources()
if (v.length == 0) {
  throw "Expected at least one server preferred resource"
}

v = kubernetes.discovery.serverPreferredNamespacedResources()
if (v.length == 0) {
  throw "Expected at least one server preferred namespaced resource"
}

res = kubernetes.discovery.serverResourcesForGroupVersion("batch/v1")
if (res.groupVersion != "batch/v1") {
  throw "Expected groupVersion batch/v1, got " + res.groupVersion
}
