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
	"github.com/jacobsa/oglematchers"
)

// callExpectation represents an expectation for zero or more calls to a mock
// method, and a set of actions to be taken when those calls are received.
type callExpectation struct {
	// Matchers that the arguments to the mock method must satisfy in order to
	// match this expectation.
	argMatchers []oglematchers.Matcher

	// The name of the file in which this expectation was expressed.
	fileName string

	// The line number at which this expectation was expressed.
	lineNumber int

	// The number of times this expectation should be matched, as explicitly
	// listed by the user. If there was no explicit number expressed, this is -1.
	expectedNumMatches int

	// Actions to be taken for the first N calls, one per call in order, where N is the
	// length of this slice.
	oneTimeActions []Action

	// An action to be taken when the one-time actions have expired, or nil if
	// there is no such action.
	fallbackAction Action
}

// newExpectation creates an expectation with the supplied info that is
// otherwise empty.
func newExpecation(args []interface{}, fileName string, lineNumber int) Expectation {
	return nil
}
