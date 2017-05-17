Go Hello World
==============

Golang is gaining a lot of traction at work so here is my "Hello World"

In order to get started with Golang you will need 
* Install Go, see https://golang.org/doc/install
* An editor (I use Visual Studio Code), see https://code.visualstudio.com/

Optional
* Go plugin for your editor, see https://code.visualstudio.com/docs/languages/go
* Debugger for your editor, see https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code

For this project
* Gingko https://onsi.github.io/ginkgo/
* Gomega http://onsi.github.io/gomega/

Usage
-----
* Download source code of this project from github
* From command line type "go build"
* From command line type "hello"

You will recive an appropriate greeting for the time of day!

Tests
-----
When a Go program is built the end result is a single runnable program (well almost, some libraries are linked at runtime).

The Go compiler can compile for multiple operating systems but there are some quirks, you can include or exclude code based on filenames or build flags.

One "known" filename is for unit test files, these must end _test.go, see overview at https://golang.org/pkg/testing/

This code comes complete with a couple of basic unit tests ... just to ensure I can refactor the code safely without messing up the time based greeting.

There are three test files
* hello_test.go - tests the greeting program using standard Go
* ginkgo_reference_test.go - self contained Gingko example
* gingko_hello_test.go - tests the greeting program with Gingko

To run the tests from command line type
```
$ go test
```

NB: If you do not have Gingko and Gomega installed then you will need to install ginkgo or delete the files prefixed with ginkgo.

Contributing
------------
Please see [CONTRIBUTING.md][contributor] and read the [CODE_OF_CONDUCT.md][conduct]

License and Authors
-------------------
Please see [LICENSE][licence]
Authors: Chris Sullivan

[contributor]: https://github.com/chrisgit/go-hello_world/blob/master/CONTRIBUTING.md
[conduct]: https://github.com/chrisgit/go-hello_world/blob/master/CODE_OF_CONDUCT.md
[licence]: https://github.com/chrisgit/go-hello_world/blob/master/LICENSE