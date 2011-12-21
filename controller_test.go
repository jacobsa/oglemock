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

package oglemock_test

import (
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/oglemock"
	. "github.com/jacobsa/ogletest"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type errorReport struct {
	fileName string
	lineNumber int
	err error
}

type fakeErrorReporter struct {
	errorsReported []errorReport
}

func (r *fakeErrorReporter) ReportError(fileName string, lineNumber int, err error) {
	report := errorReport{fileName, lineNumber, err}
	r.errorsReported = append(r.errorsReported, report)
}

type trivialMockObject struct {
	id uintptr
	desc string
}

func (o *trivialMockObject) Oglemock_Id() uintptr {
	return o.id
}

func (o *trivialMockObject) Oglemock_Description() string {
	return o.desc
}

// Method being mocked
func (o *trivialMockObject) StringToInt(s string) int {
	return 0
}

// Method being mocked
func (o *trivialMockObject) TwoIntsToString(i, j int) string {
	return ""
}

type ControllerTest struct {
	reporter fakeErrorReporter
	controller Controller

	mock1 MockObject
	mock2 MockObject
}

func (t *ControllerTest) SetUp() {
	t.reporter.errorsReported = make([]errorReport, 0)
	t.controller = NewController(&t.reporter)

	t.mock1 = &trivialMockObject{17, "taco"}
	t.mock2 = &trivialMockObject{19, "burrito"}
}

func init() { RegisterTestSuite(&ControllerTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *ControllerTest) FinishWithoutAnyEvents() {
	t.controller.Finish()
	ExpectThat(len(t.reporter.errorsReported), Equals(0))
}

func (t *ControllerTest) HandleCallForUnknownObject() {
	p := []byte{255}
	t.controller.HandleMethodCall(
		t.mock1,
		"StringToInt",
		"taco.go",
		112,
		[]interface{}{p})

	// The error should be reported immediately.
	ExpectThat(len(t.reporter.errorsReported), Equals(1))
	ExpectThat(t.reporter.errorsReported[0].fileName, Equals("taco.go"))
	ExpectThat(t.reporter.errorsReported[0].lineNumber, Equals(112))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("Unexpected")))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("Read")))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("[255]")))

	// Finish should change nothing.
	t.controller.Finish()
	ExpectThat(len(t.reporter.errorsReported), Equals(1))
}

func (t *ControllerTest) ExpectCallForUnknownMethod() {
	ExpectThat(
		func() {
			t.controller.ExpectCall(t.mock1, "Frobnicate", "", 0)
		},
		Panics(HasSubstr("Unknown method: Frobnicate")))
}

func (t *ControllerTest) PartialExpectationGivenWrongNumberOfArgs() {
	ExpectThat(
		func() {
			t.controller.ExpectCall(t.mock1, "TwoIntsToString", "", 0)(17, 19, 23)
		},
		Panics(HasSubstr("arguments: expected 2, got 3")))
}

func (t *ControllerTest) PartialExpectationCalledTwice() {
	ExpectThat(
		func() {
			partial := t.controller.ExpectCall(t.mock1, "StringToInt", "", 0)
			partial("taco")
			partial("taco")
		},
		Panics(HasSubstr("called more than once")))
}

func (t *ControllerTest) ExpectThenNonMatchingCall() {
	// Expectation
	partial := t.controller.ExpectCall(
		t.mock1,
		"TwoIntsToString",
		"burrito.go",
		117)

	partial(LessThan(10), Equals(2))

	// Call
	t.controller.HandleMethodCall(
		t.mock1,
		"TwoIntsToString",
		"taco.go",
		112,
		[]interface{}{8, 1})


	// The error should be reported immediately.
	ExpectThat(len(t.reporter.errorsReported), Equals(1))
	ExpectThat(t.reporter.errorsReported[0].fileName, Equals("taco.go"))
	ExpectThat(t.reporter.errorsReported[0].lineNumber, Equals(112))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("Unexpected")))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("Tried")))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("burrito.go:117")))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("arg 1")))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("Expected: 2")))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("Actual:   1")))

	// Finish should change nothing.
	t.controller.Finish()
	ExpectThat(len(t.reporter.errorsReported), Equals(1))
}

func (t *ControllerTest) ExplicitCardinalityNotSatisfied() {
	// Expectation
	partial := t.controller.ExpectCall(
		t.mock1,
		"StringToInt",
		"burrito.go",
		117)

	exp := partial(HasSubstr(""))
	exp.Times(3)

	// Call twice.
	t.controller.HandleMethodCall(
		t.mock1,
		"TwoIntsToString",
		"",
		0,
		[]interface{}{""})

	t.controller.HandleMethodCall(
		t.mock1,
		"TwoIntsToString",
		"",
		0,
		[]interface{}{""})

	// The error should not yet be reported.
	ExpectThat(len(t.reporter.errorsReported), Equals(0))

	// Finish should cause the error to be reported.
	t.controller.Finish()

	ExpectThat(len(t.reporter.errorsReported), Equals(1))
	ExpectThat(t.reporter.errorsReported[0].fileName, Equals("burrito.go"))
	ExpectThat(t.reporter.errorsReported[0].lineNumber, Equals(117))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("Unsatisfied")))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("StringToInt")))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("has substring \"\"")))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("called 3 times")))
	ExpectThat(t.reporter.errorsReported[0].err, Error(HasSubstr("called 2 times")))
}

func (t *ControllerTest) ImplicitOneTimeActionCountNotSatisfied() {
}

func (t *ControllerTest) ImplicitOneTimeActionLowerBoundNotSatisfied() {
}

func (t *ControllerTest) ImplicitCardinalityOfOneNotSatisfied() {
}

func (t *ControllerTest) ExplicitCardinalitySatisfied() {
}

func (t *ControllerTest) ImplicitOneTimeActionCountSatisfied() {
}

func (t *ControllerTest) ImplicitOneTimeActionLowerBoundSatisfied() {
}

func (t *ControllerTest) FallbackActionConfiguredWithZeroCalls() {
}

func (t *ControllerTest) FallbackActionConfiguredWithMultipleCalls() {
}

func (t *ControllerTest) ImplicitCardinalityOfOneSatisfied() {
}

func (t *ControllerTest) InvokesOneTimeActions() {
}

func (t *ControllerTest) InvokesFallbackActions() {
}

func (t *ControllerTest) InvokesImplicitActions() {
}

func (t *ControllerTest) ExpectationsAreMatchedLastToFirst() {
}

func (t *ControllerTest) ExpectationsAreSegregatedByMockObject() {
}

func (t *ControllerTest) ExpectationsAreSegregatedByMethodName() {
}
