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

[googlemock]: http://code.google.com/p/googlemock/
[google-js-test]: http://code.google.com/p/google-js-test/
[oglematchers]: https://github.com/jacobsa/oglematchers
[ogletest]: https://github.com/jacobsa/oglematchers
