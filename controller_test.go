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

type ControllerTest struct {
}

func init() { RegisterTestSuite(&ControllerTest{}) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *ControllerTest) FinishWithoutAnyEvents() {
}

func (t *ControllerTest) HandleCallForUnknownObject() {
}

func (t *ControllerTest) ExpectCallForUnknownMethod() {
}

func (t *ControllerTest) PartialExpectationGivenWrongNumberOfArgs() {
}

func (t *ControllerTest) PartialExpectationCalledTwice() {
}

func (t *ControllerTest) ExpectThenNonMatchingCall() {
}

func (t *ControllerTest) ExplicitCardinalityNotSatisfied() {
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
