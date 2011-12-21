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

type CallExpectationTest struct {
}

func init() { RegisterTestSuite(&CallExpectationTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *CallExpectationTest) StoresFileNameAndLineNumber() {
	args := []interface{}{}
	exp := InternalNewExpectation(emptyReturnSig, args, "taco", 17)

	ExpectThat(exp.FileName, Equals("taco"))
	ExpectThat(exp.LineNumber, Equals(17))
}

func (t *CallExpectationTest) NoArgs() {
	args := []interface{}{}
	exp := InternalNewExpectation(emptyReturnSig, args, "", 0)

	ExpectThat(len(exp.ArgMatchers), Equals(0))
}

func (t *CallExpectationTest) MixOfMatchersAndNonMatchers() {
	args := []interface{}{Equals(17), 19, Equals(23)}
	exp := InternalNewExpectation(emptyReturnSig, args, "", 0)

	// Matcher args
	ExpectThat(len(exp.ArgMatchers), Equals(3))
	ExpectThat(exp.ArgMatchers[0], Equals(args[0]))
	ExpectThat(exp.ArgMatchers[2], Equals(args[2]))

	// Non-matcher arg
	var res MatchResult
	matcher1 := exp.ArgMatchers[1]

	res, _ = matcher1.Matches(17)
	ExpectThat(res, Equals(MATCH_FALSE))

	res, _ = matcher1.Matches(19)
	ExpectThat(res, Equals(MATCH_TRUE))

	res, _ = matcher1.Matches(23)
	ExpectThat(res, Equals(MATCH_FALSE))
}

func (t *CallExpectationTest) NoTimes() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(exp.ExpectedNumMatches, Equals(-1))
}

func (t *CallExpectationTest) TimesN() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)
	exp.Times(17)

	ExpectThat(exp.ExpectedNumMatches, Equals(17))
}

func (t *CallExpectationTest) NoActions() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(len(exp.OneTimeActions), Equals(0))
	ExpectThat(exp.FallbackAction, Equals(nil))
}

func (t *CallExpectationTest) WillOnce() {
	action0 := Return(17.0)
	action1 := Return(19.0)

	exp := InternalNewExpectation(float64ReturnSig, []interface{}{}, "", 0)
	exp.WillOnce(action0).WillOnce(action1)

	ExpectThat(len(exp.OneTimeActions), Equals(2))
	ExpectThat(exp.OneTimeActions[0], Equals(action0))
	ExpectThat(exp.OneTimeActions[1], Equals(action1))
}

func (t *CallExpectationTest) WillRepeatedly() {
	action := Return(17.0)

	exp := InternalNewExpectation(float64ReturnSig, []interface{}{}, "", 0)
	exp.WillRepeatedly(action)

	ExpectThat(exp.FallbackAction, Equals(action))
}

func (t *CallExpectationTest) BothKindsOfAction() {
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

func (t *CallExpectationTest) TimesCalledWithHugeNumber() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.Times(1 << 30) },
		Panics(Error(HasSubstr("Times: N must be at most 1000"))))
}

func (t *CallExpectationTest) TimesCalledTwice() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.Times(17).Times(17) },
		Panics(Error(HasSubstr("Times called more than"))))
}

func (t *CallExpectationTest) TimesCalledAfterWillOnce() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillOnce(Return()).Times(17) },
		Panics(Error(HasSubstr("Times called after WillOnce"))))
}

func (t *CallExpectationTest) TimesCalledAfterWillRepeatedly() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillRepeatedly(Return()).Times(17) },
		Panics(Error(HasSubstr("Times called after WillRepeatedly"))))
}

func (t *CallExpectationTest) WillOnceCalledAfterWillRepeatedly() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillRepeatedly(Return()).WillOnce(Return()) },
		Panics(Error(HasSubstr("WillOnce called after WillRepeatedly"))))
}

func (t *CallExpectationTest) OneTimeActionRejectsSignature() {
	action := Return("taco")
	exp := InternalNewExpectation(float64ReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillOnce(action) },
		Panics(Error(HasSubstr("arg 0; expected float64"))))
}

func (t *CallExpectationTest) WillRepeatedlyCalledTwice() {
	exp := InternalNewExpectation(emptyReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillRepeatedly(Return()).WillRepeatedly(Return()) },
		Panics(Error(HasSubstr("WillRepeatedly called more than once"))))
}

func (t *CallExpectationTest) FallbackActionRejectsSignature() {
	action := Return("taco")
	exp := InternalNewExpectation(float64ReturnSig, []interface{}{}, "", 0)

	ExpectThat(
		func() { exp.WillRepeatedly(action) },
		Panics(HasSubstr("arg 0; expected float64")))
}
