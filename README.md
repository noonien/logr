logr
====
[![Build Status](http://img.shields.io/travis/noonien/logr/master.svg?style=flat-square)](https://travis-ci.org/noonien/logr)
[![Coverage Status](http://img.shields.io/coveralls/noonien/logr/master.svg?style=flat-square)](https://coveralls.io/r/noonien/logr)
[![GoDoc](http://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/noonien/logr)
[![License: MIT](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](http://opensource.org/licenses/MIT)

`logr` aims to provide powerful logging for Go applications.


Goals
-----
  - Provide a way of logging as much data as possible from the application and
    its dependencies.
  - Add structure to the logging facilities needed throughout the application.
  - Simplify inserting context data.
  - Easy extensibility trough the use of few simple interfaces.
  - Even though filters are provided, filtering data is not encouraged, and
    should be done by logging processors.
  - Test as much of the code as possible.


Status
------
While unit tests are provided for most of the code, `logr` has not been
battle-tested, so it is not yet recommended for use in a production environment.

[Documentation](https://godoc.org/github.com/noonien/logr) is severely lacking.


Contributing
------------
Contributing is always welcome and appreciated! Pull requests and issues will
be processed as fast as possible!


License
---------
Copyright (c) 2014 George-Cristian Jiglau

This project is licensed under the [MIT license](http://opensource.org/licenses/MIT).
See the LICENSE file for more details.
