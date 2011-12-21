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

// InternalCallExpectation is exported for purposes of testing only. You should
// not touch it.
//
// InternalCallExpectation represents an expectation for zero or more calls to
// a mock method, and a set of actions to be taken when those calls are
// received.
type InternalCallExpectation struct {
	// Matchers that the arguments to the mock method must satisfy in order to
	// match this expectation.
	ArgMatchers []oglematchers.Matcher

	// The name of the file in which this expectation was expressed.
	FileName string

	// The line number at which this expectation was expressed.
	LineNumber int

	// The number of times this expectation should be matched, as explicitly
	// listed by the user. If there was no explicit number expressed, this is -1.
	ExpectedNumMatches int

	// Actions to be taken for the first N calls, one per call in order, where N
	// is the length of this slice.
	OneTimeActions []Action

	// An action to be taken when the one-time actions have expired, or nil if
	// there is no such action.
	FallbackAction Action
}

// InternalNewExpectation is exported for purposes of testing only. You should
// not touch it.
func InternalNewExpectation(
	args []interface{},
	fileName string,
	lineNumber int) *InternalCallExpectation {
	return nil
}
