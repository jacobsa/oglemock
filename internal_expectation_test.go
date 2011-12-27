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
	"reflect"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

var emptyReturnSig reflect.Type = reflect.TypeOf(func(i int) {})
var float64ReturnSig reflect.Type = reflect.TypeOf(func(i int) float64 { return 17.0 })

type InternalExpectationTest struct {
}

func init() { RegisterTestSuite(&InternalExpectationTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *InternalExpectationTest) StoresFileNameAndLineNumber() {
	args := []interface{}{}
	exp := InternalNewExpectation(emptyReturnSig, args, "taco", 17)

	ExpectThat(exp.FileName, Equals("taco"))
	ExpectThat(exp.LineNumber, Equals(17))
}

func (t *InternalExpectationTest) NoArgs() {
	args := []interface{}{}
	exp := InternalNewExpectation(emptyReturnSig, args, "", 0)

	ExpectThat(len(exp.ArgMatchers), Equals(0))
}

func (t *InternalExpectationTest) MixOfMatchersAndNonMatchers() {
	args := []interface{}{Equals(17), 19, Equals(23)}
	exp := InternalNewExpectation(emptyReturnSig, args, "", 0)

	// Matcher args
	ExpectThat(len(exp.ArgMatchers), Equals(3))
	ExpectThat(exp.ArgMatchers[0], Equals(args[0]))
	ExpectThat(exp.ArgMatchers[2], Equals(args[2]))

	// Non-matcher arg
	var res bool
	matcher1 := exp.ArgMatchers[1]

	res, _ = matcher1.Matches(17)
	ExpectFalse(res)

	res, _ = matcher1.Matches(19)
	ExpectTrue(res)

	res, _ = matcher1.Matches(23)
	ExpectFalse(res)
}

func (t *InternalExpectationTest) NoTimes() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(exp.ExpectedNumMatches, Equals(-1))
}

func (t *InternalExpectationTest) TimesN() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)
	exp.Times(17)

	ExpectThat(exp.ExpectedNumMatches, Equals(17))
}

func (t *InternalExpectationTest) NoActions() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(len(exp.OneTimeActions), Equals(0))
	ExpectThat(exp.FallbackAction, Equals(nil))
}

func (t *InternalExpectationTest) WillOnce() {
	action0 := Return(17.0)
	action1 := Return(19.0)

	exp := InternalNewExpectation(float64ReturnSig, []interface{}{}, "", 0)
	exp.WillOnce(action0).WillOnce(action1)

	ExpectThat(len(exp.OneTimeActions), Equals(2))
	ExpectThat(exp.OneTimeActions[0], Equals(action0))
	ExpectThat(exp.OneTimeActions[1], Equals(action1))
}

func (t *InternalExpectationTest) WillRepeatedly() {
	action := Return(17.0)

	exp := InternalNewExpectation(float64ReturnSig, []interface{}{}, "", 0)
	exp.WillRepeatedly(action)

	ExpectThat(exp.FallbackAction, Equals(action))
}

func (t *InternalExpectationTest) BothKindsOfAction() {
	action0 := Return(17.0)
	action1 := Return(19.0)
	action2 := Return(23.0)

	exp := InternalNewExpectation(float64ReturnSig, []interface{}{}, "", 0)
	exp.WillOnce(action0).WillOnce(action1).WillRepeatedly(action2)

	ExpectThat(len(exp.OneTimeActions), Equals(2))
	ExpectThat(exp.OneTimeActions[0], Equals(action0))
	ExpectThat(exp.OneTimeActions[1], Equals(action1))
	ExpectThat(exp.FallbackAction, Equals(action2))
}

func (t *InternalExpectationTest) TimesCalledWithHugeNumber() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.Times(1 << 30) },
		Panics(HasSubstr("Times: N must be at most 1000")))
}

func (t *InternalExpectationTest) TimesCalledTwice() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.Times(17).Times(17) },
		Panics(HasSubstr("Times called more than")))
}

func (t *InternalExpectationTest) TimesCalledAfterWillOnce() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillOnce(Return()).Times(17) },
		Panics(HasSubstr("Times called after WillOnce")))
}

func (t *InternalExpectationTest) TimesCalledAfterWillRepeatedly() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillRepeatedly(Return()).Times(17) },
		Panics(HasSubstr("Times called after WillRepeatedly")))
}

func (t *InternalExpectationTest) WillOnceCalledAfterWillRepeatedly() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillRepeatedly(Return()).WillOnce(Return()) },
		Panics(HasSubstr("WillOnce called after WillRepeatedly")))
}

func (t *InternalExpectationTest) OneTimeActionRejectsSignature() {
	action := Return("taco")
	exp := InternalNewExpectation(float64ReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillOnce(action) },
		Panics(HasSubstr("arg 0; expected float64")))
}

func (t *InternalExpectationTest) WillRepeatedlyCalledTwice() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillRepeatedly(Return()).WillRepeatedly(Return()) },
		Panics(HasSubstr("WillRepeatedly called more than once")))
}

func (t *InternalExpectationTest) FallbackActionRejectsSignature() {
	action := Return("taco")
	exp := InternalNewExpectation(float64ReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillRepeatedly(action) },
		Panics(HasSubstr("arg 0; expected float64")))
}
