# gotests [![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://github.com/hexdigest/gotests/blob/master/LICENSE) [![godoc](https://img.shields.io/badge/go-documentation-blue.svg)](https://godoc.org/github.com/hexdigest/gotests) [![Build Status](https://travis-ci.org/hexdigest/gotests.svg?branch=master)](https://travis-ci.org/hexdigest/gotests) [![Coverage Status](https://coveralls.io/repos/github/hexdigest/gotests/badge.svg?branch=master)](https://coveralls.io/github/hexdigest/gotests?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/hexdigest/gotests)](https://goreportcard.com/report/github.com/hexdigest/gotests) [![Coverage Status](https://coveralls.io/repos/github/cweill/gotests/badge.svg?branch=master)]

This repository is a fork of [cweill/gotests](https://github.com/cweill/gotests)
The difference from the original code is that generated test is integrated with [minimock](https://github.com/cweill/gotests) and has setup function when testing a method of a struct.

`gotests` makes writing Go tests easy. It's a Golang commandline tool that generates [table driven tests](https://github.com/golang/go/wiki/TableDrivenTests) based on its target source files' function and method signatures. Any new dependencies in the test files are automatically imported.

## Demo

The following shows `gotests` in action using the [official Sublime Text 3 plugin](https://github.com/cweill/GoTests-Sublime). Plugins also exist for [Emacs](https://github.com/damienlevin/GoTests-Emacs), [Vim](https://github.com/buoto/gotests-vim) and [Atom Editor](https://atom.io/packages/gotests).

![demo](https://github.com/cweill/GoTests-Sublime/blob/master/gotests.gif)

## Installation

__Minimum Go version:__ Go 1.6

Use [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install and update:

```sh
$ go get -u github.com/hexdigest/gotests/...
```

## Usage

From the commandline, `gotests` can generate Go tests for specific source files or an entire directory. By default, it prints its output to `stdout`.

```sh
$ gotests [options] PATH ...
```

Available options:

```
  -all         generate go tests for all functions and methods
  
  -excl        regexp. generate go tests for functions and methods that don't 
               match. Takes precedence over -only, -exported, and -all
    	   
  -exported    generate go tests for exported functions and methods. Takes 
               precedence over -only and -all

  -i	       print test inputs in error messages
  
  -only        regexp. generate go tests for functions and methods that match only.
               Takes precedence over -all
  
  -w           write output to (test) files instead of stdout
  
  -nosubtests  disable subtest generation. Only available for Go 1.7+
```

## License

`gotests` is released under the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0).
