console.log("====> default namespace test")
if (kubernetes.getNS() != "default") {
  throw "Expected default namespace, got " + kubernetes.getNS()
}
