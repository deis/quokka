console.log("====> serverVersion test")

v = kubernetes.discovery.serverVersion()

console.log(v.major)

if (v.gitVersion == '') {
  throw "No gitVersion"
}

