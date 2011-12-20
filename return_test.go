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
	"testing"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type ReturnTest struct {
}

func init()                     { RegisterTestSuite(&ReturnTest{}) }
func TestOgletest(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *ReturnTest) EmptySet() {
	action := Return()
	result := action.Invoke([]interface{}{})

	ExpectThat(len(result), Equals(0))
}

func (t *ReturnTest) OneValue() {
	action := Return("taco")
	result := action.Invoke([]interface{}{})

	ExpectThat(len(result), Equals(1))
	ExpectThat(result[0], Equals("taco"))
}

func (t *ReturnTest) MultipleValues() {
	someInt := 17
	action := Return("taco", &someInt, 19)
	result := action.Invoke([]interface{}{})

	ExpectThat(len(result), Equals(3))
	ExpectThat(result[0], Equals("taco"))
	ExpectThat(result[1], Equals(&someInt))
	ExpectThat(result[2], Equals(19))
}
