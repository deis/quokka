hello = require("simplemodule/index.js");

console.log(hello)

r = hello("world");
if (r != "hello, world") {
  throw "Got " + r;
}
