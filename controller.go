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

// PartialExpecation is a function that should be called exactly once with
// expected arguments or matchers in order to set up an expected method call.
// See Controller.ExpectMethodCall below. It returns an expectation that can be
// further modified (e.g. by calling WillOnce).
//
// If the arguments are of the wrong type, the function panics.
type PartialExpecation func([]interface{}) Expectation

// Controller represents an object that implements the central logic of
// oglemock: recording and verifying expectations, responding to mock method
// calls, and so on.
type Controller interface {
	// ExpectCall expresses an expectation that the method of the given name
	// should be called on the supplied mock object. It returns a function that
	// should be called with the expected arguments, matchers for the arguments,
	// or a mix of both.
	//
	// For example:
	//
	//     mockWriter := [...]
	//     controller.ExpectCall(mockWriter, "Write")(ElementsAre(0x1))
	//         .WillOnce(Return(1, nil))
	//
	// If the mock object doesn't have a method of the supplied name, the
	// function panics.
	ExpectCall(o MockObject, methodName string) PartialExpecation

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
	return nil
}
