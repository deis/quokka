console.log("====> withNS test")
k = kubernetes.withNS("quokkatest")
if (k.getNS() != "quokkatest") {
  throw "Expected quokkatest ns, got " + k.getNS()
}

if (k.discovery == undefined) {
  throw "discover object is missing"
}
