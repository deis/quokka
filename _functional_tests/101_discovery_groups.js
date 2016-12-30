console.log("====> serverGroups test")

v = kubernetes.discovery.serverGroups()
if (v.kind != 'APIGroupList') {
  throw "Unexpected kind: " + v.kind
}

if (v.groups.length == 0) {
  throw "Expected at least one group"
}

_.each(v.groups, function(e, i, l){
  console.log(e.name)
})
