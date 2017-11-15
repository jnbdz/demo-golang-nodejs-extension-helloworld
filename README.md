# demo-golang-nodejs-extension-helloworld
[![Support us on Patreon][badge_patreon]][patreon] [![Build Status][badge_travis]][travis]

Demo: Golang Node.js Extension: Hello World!

## :cloud: Installation
Make sure to have the `npm` package `node-gyp` installed.

```sh
$ npm i --save demo-golang-nodejs-extension-helloworld
```

You might need to run:

```sh
npm install
```

It will execute:

```sh
go build -buildmode c-archive -o src/helloworld.a src/helloworld.go && node-gyp rebuild
```

If you run into some errors you might need to install other packages on your system like: `gcc` and `build-essential`.

## :clipboard: Example

```js
const msg = require("demo-golang-nodejs-extension-helloworld");

console.log(msg.hello());
```

## :recycle: Testing
To run the unit test, you need to make sure you have [Jest][jest_link] installed.

You can install it with [npm][npm_link]:

```sh
npm install --save-dev jest
```

Or you can install it with [yarn][yarn_link]:

```sh
yarn add --dev jest
```

Or globally with [npm][npm_link]:

```sh
sudo npm install -g jest --unsafe-perm=true --allow-root
```

### Running the tests

```sh
npm test
```

## :memo: Documentation
Based on this tutorial: https://www.krishnaraman.net/blog/node-addons-written-in-go

When creating an Node.js extension with Golang make sure to import `C` package in your Golang script:

```go
import "C"
```

Then add the function or functions you want to implement in your extension.

For this demo:

```go
func HelloWorld() *C.char {
	cs := C.CString("Hello World!")
	return cs
}
```

When using strings you need to set the return type as `*C.char`.

All your strings need to be set in `C.CString()` method.

You can keep `main()` function empty:

```go
func main() {}
```

Then you need to add your `.cc` (C/C++) file.

For this demo the name of the file is `node-helloworld.cc`.

First you need to include these two libraries:

```c
#include "helloworld.h"
#include <node.h>
```

The `helloworld.h` will be generated later.

The `<node.h>` is the Node.js library.

Then you add the rest of the code:

```c
#include "helloworld.h"
#include <node.h>

namespace demo {

using v8::FunctionCallbackInfo;
using v8::Isolate;
using v8::Local;
using v8::Object;
using v8::String;
using v8::Value;

void Method(const FunctionCallbackInfo<Value>& args) {
  Isolate* isolate = args.GetIsolate();
  args.GetReturnValue().Set(String::NewFromUtf8(isolate, HelloWorld()));
}

void init(Local<Object> exports) {
  NODE_SET_METHOD(exports, "hello", Method);
}

NODE_MODULE(helloworld, init)

}
```

You need to add the C++ methods that you want to have access to in Node.js.

For this demo we are using:

```c
void Method(const FunctionCallbackInfo<Value>& args)
```

The last part is to create the methods for Node.js:

```c
void init(Local<Object> exports) {
NODE_SET_METHOD(exports, "hello", Method);
}

NODE_MODULE(helloworld, init)
```

You use `NODE_SET_METHOD` to set the methods you want to use inside of Node.js:

```c
NODE_SET_METHOD(exports, "hello", Method);
```

You also need this file at the root of your project: `binding.gyp`. It is self explanatory.

```json
{
  "targets": [
    {
      "target_name": "node-helloworld",
      "sources": [
        "src/node-helloworld.cc"
      ],
      "libraries": [
        "../src/helloworld.a"
      ],
    },
  ],
}
```

The file `helloworld.a` is also automatically generated. It should not be commited.

When you run `npm install` it runs:

```sh
go build -buildmode c-archive -o src/helloworld.a src/helloworld.go && node-gyp rebuild
```

The first command generates `helloworld.a` and `helloworld.h`. With the other command `node-gyp rebuild` it generates `build/` directory.

Make sure you have `node-gyp` installed. You can install it with `npm`.

## :scroll: License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

[badge_patreon]: https://propolisframework.github.io/assets/img/patreon.svg
[badge_travis]: https://travis-ci.org/jnbdz/demo-golang-nodejs-extension-helloworld.svg?branch=master

[patreon]: https://www.patreon.com/jnbdz
[travis]: https://travis-ci.org/jnbdz/demo-golang-nodejs-extension-helloworld

[jest_link]: https://facebook.github.io/jest/
[npm_link]: https://npmjs.com/
[yarn_link]: https://yarnpkg.com/