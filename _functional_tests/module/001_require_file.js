
var zed = require("./lib/zed.js")

if (zed.hello() != "hello") {
  throw "expected hello, got " + zed.hello();
}
