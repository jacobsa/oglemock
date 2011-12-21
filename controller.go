// Copyright 2011 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oglemock

import (
	"fmt"
	"github.com/jacobsa/oglematchers"
	"reflect"
)

// PartialExpecation is a function that should be called exactly once with
// expected arguments or matchers in order to set up an expected method call.
// See Controller.ExpectMethodCall below. It returns an expectation that can be
// further modified (e.g. by calling WillOnce).
//
// If the arguments are of the wrong type, the function panics.
type PartialExpecation func(...interface{}) Expectation

// Controller represents an object that implements the central logic of
// oglemock: recording and verifying expectations, responding to mock method
// calls, and so on.
type Controller interface {
	// ExpectCall expresses an expectation that the method of the given name
	// should be called on the supplied mock object. It returns a function that
	// should be called with the expected arguments, matchers for the arguments,
	// or a mix of both.
	//
	// fileName and lineNumber should indicate the line on which the expectation
	// was made, if known.
	//
	// For example:
	//
	//     mockWriter := [...]
	//     controller.ExpectCall(mockWriter, "Write")(ElementsAre(0x1))
	//         .WillOnce(Return(1, nil))
	//
	// If the mock object doesn't have a method of the supplied name, the
	// function panics.
	ExpectCall(
		o MockObject,
		methodName string,
		fileName string,
		lineNumber int) PartialExpecation

	// Finish causes the controller to check for any unsatisfied expectations,
	// and report them as errors if they exist.
	//
	// The controller may panic if any of its methods (including this one) are
	// called after Finish is called.
	Finish()

	// HandleMethodCall looks for a registered expectation matching the call of
	// the given method on mock object o, invokes the appropriate action (if
	// any), and returns the values returned by that action (if any).
	//
	// If the action returns nothing, the controller returns zero values. If
	// there is no matching expectation, the controller reports an error and
	// returns zero values.
	//
	// If the mock object doesn't have a method of the supplied name, the
	// arguments are of the wrong type, or the action returns the wrong types,
	// the function panics.
	//
	// HandleMethodCall is exported for the sake of mock implementations, and
	// should not be used directly.
	HandleMethodCall(
		o MockObject,
		methodName string,
		fileName string,
		lineNumber int,
		args ...interface{}) []interface{}
}

// NewController sets up a fresh controller, without any expectations set, and
// configures the controller to use the supplied error reporter.
func NewController(reporter ErrorReporter) Controller {
	return &controllerImpl{reporter, map[string][]*InternalExpectation{}}
}

type controllerImpl struct {
	reporter ErrorReporter
	expectations map[string][]*InternalExpectation
}

func getMapKey(o MockObject, methodName string) string {
  return fmt.Sprintf("%016x%s-", o.Oglemock_Id(), methodName)
}

func (c *controllerImpl) ExpectCall(
	o MockObject,
	methodName string,
	fileName string,
	lineNumber int) PartialExpecation {
	// Find the signature for the requested method.
	oType := reflect.TypeOf(o)
	method, ok := oType.MethodByName(methodName)
	if !ok {
		panic("Unknown method: " + methodName)
	}

	return func(args ...interface{}) Expectation {
		exp := InternalNewExpectation(method.Type, args, fileName, lineNumber)

		// Insert the expectation into the map.
		key := getMapKey(o, methodName)
		expList, ok := c.expectations[key]
		if !ok {
			expList = make([]*InternalExpectation, 0)
		}

		expList = append(expList, exp)
		c.expectations[key] = expList

		// Return the expectation to the user.
		return exp
	}
}

func (c *controllerImpl) Finish() {
	// TODO
}

// expectationMatches checks the matchers for the expectation against the
// supplied arguments.
func expectationMatches(exp *InternalExpectation, args []interface{}) bool {
	matchers := exp.ArgMatchers
	if len(args) != len(matchers) {
		panic(
			fmt.Sprintf(
				"Wrong number of arguments: expected %d; got %d",
				len(matchers),
				len(args)))
	}

	// Check each matcher.
	for i, matcher := range matchers {
		res, _ := matcher.Matches(args[i])
		if res != oglematchers.MATCH_TRUE {
			return false
		}
	}

	return true
}

// chooseExpectation returns the expectation that matches the supplied
// arguments. If there is more than one such expectation, the one furthest
// along in the list for the method is returned. If there is no such
// expectation, nil is returned.
func (c *controllerImpl) chooseExpectation(
	o MockObject,
	methodName string,
	args []interface{}) *InternalExpectation {
	// Do we have any expectations for this object?
	mapKey := getMapKey(o, methodName)
	expectations, ok := c.expectations[mapKey]
	if !ok || len(expectations) == 0 {
		return nil
	}

	for i := len(expectations) - 1; i >= 0; i-- {
		if (expectationMatches(expectations[i], args)) {
			return expectations[i]
		}
	}

	return nil
}

func (c *controllerImpl) HandleMethodCall(
	o MockObject,
	methodName string,
	fileName string,
	lineNumber int,
	args ...interface{}) []interface{} {
	// TODO
	return nil
}
