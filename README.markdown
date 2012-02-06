`oglemock` is a mocking framework for the Go programming language with the
following features:

 *  An extensive and extensible set of matchers for expressing call
    expectations (provided by the [oglematchers][] package).

 *  Clean, readable output that tells you exaclty what you need to know.

 *  Style and semantics similar to [Google Mock][googlemock] and
    [Google JS Test][google-js-test].

 *  Seamless integration with the [ogletest][] unit testing framework.

It can be integrated into any testing framework (including Go's `testing`
package), but out of the box support is built in to [ogletest][] and that is the
easiest place to use it.


Installation
------------

First, make sure you have installed a version of the Go tools at least as new as
`weekly/2012-01-27`. See [here][golang-install] for instructions. Until
release `r61` comes out, this involes using the `weekly` tag.

Use the following command to install `oglemock` and its dependencies, and to
keep them up to date:

    goinstall -u github.com/jacobsa/oglemock
    goinstall -u github.com/jacobsa/oglemock/createmock

Those commands will install the `oglemock` package itself, along with the
`createmock` tool that is used to auto-generate mock classes.


[golang-install]: http://golang.org/doc/install.html#releases
[google-js-test]: http://code.google.com/p/google-js-test/
[googlemock]: http://code.google.com/p/googlemock/
[oglematchers]: https://github.com/jacobsa/oglematchers
[ogletest]: https://github.com/jacobsa/oglematchers
