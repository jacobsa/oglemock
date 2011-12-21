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
	. "github.com/jacobsa/ogletest"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type CallExpectationTest struct {

}

func init()                     { RegisterTestSuite(&CallExpectationTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *CallExpectationTest) StoresFileNameAndLineNumber() {
}

func (t *CallExpectationTest) NoArgs() {
}

func (t *CallExpectationTest) MixOfMatchersAndNonMatchers() {
}

func (t *CallExpectationTest) NoTimes() {
}

func (t *CallExpectationTest) TimesN() {
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
