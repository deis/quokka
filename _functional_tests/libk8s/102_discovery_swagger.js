console.log("====> swaggerSchema tests")

v = kubernetes.discovery.swaggerSchema("batch", "v1")

if (v.apiVersion != 'batch/v1') {
  throw "Unexpected version: " + v.apiVersion
}
