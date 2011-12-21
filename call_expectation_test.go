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

type CallExpectationTest struct {

}

func init() { RegisterTestSuite(&CallExpectationTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *CallExpectationTest) StoresFileNameAndLineNumber() {
	args := []interface{}{}
	exp := InternalNewExpectation(args, "taco", 17)

	ExpectThat(exp.FileName, Equals("taco"))
	ExpectThat(exp.LineNumber, Equals(17))
}

func (t *CallExpectationTest) NoArgs() {
	args := []interface{}{}
	exp := InternalNewExpectation(args, "", 0)

	ExpectThat(len(exp.ArgMatchers), Equals(0))
}

func (t *CallExpectationTest) MixOfMatchersAndNonMatchers() {
	args := []interface{}{ Equals(17), 19, Equals(23) }
	exp := InternalNewExpectation(args, "", 0)

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
	exp := InternalNewExpectation([]interface{}{}, "", 0)

	ExpectThat(exp.ExpectedNumMatches, Equals(-1))
}

func (t *CallExpectationTest) TimesN() {
	exp := InternalNewExpectation([]interface{}{}, "", 0)
	exp.Times(17)

	ExpectThat(exp.ExpectedNumMatches, Equals(17))
}

func (t *CallExpectationTest) NoActions() {
}

func (t *CallExpectationTest) WillOnce() {
}

func (t *CallExpectationTest) WillRepeatedly() {
}

func (t *CallExpectationTest) TimesCalledTwice() {
}

func (t *CallExpectationTest) TimesCalledAfterWillOnce() {
}

func (t *CallExpectationTest) TimesCalledAfterWillRepeatedly() {
}

func (t *CallExpectationTest) WillOnceCalledAfterWillRepeatedly() {
}

func (t *CallExpectationTest) OneTimeActionRejectsSignature() {
}

func (t *CallExpectationTest) WillRepeatedlyCalledTwice() {
}

func (t *CallExpectationTest) FallbackActionRejectsSignature() {
}
