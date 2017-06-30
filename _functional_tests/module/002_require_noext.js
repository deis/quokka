var zed = require("./lib/zed")

if (zed.hello() != "hello") {
  throw "expected hello, got " + zed.hello();
}
