hello = require("minimalmodule/index.js");

console.log(typeof hello)

if (hello.msg != "hello world") {
  throw "expected 'hello world', Got " + hello.msg
}
