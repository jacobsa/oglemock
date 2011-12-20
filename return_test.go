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

	// Invoke
	result := action.Invoke([]interface{}{})
	ExpectThat(len(result), Equals(0))

	// CheckType
	emptyReturn := reflect.TypeOf(func(i int) {})
	nonEmptyReturn := reflect.TypeOf(func(i int) string { return "" })

	ExpectThat(action.CheckType(emptyReturn), Equals(nil))
	ExpectThat(action.CheckType(nonEmptyReturn), Not(Equals(nil)))
}

func (t *ReturnTest) OneValue() {
	action := Return("taco")

	// Invoke
	result := action.Invoke([]interface{}{})

	ExpectThat(len(result), Equals(1))
	ExpectThat(result[0], Equals("taco"))

	// CheckType
	type compatibleType string
	emptyReturn := reflect.TypeOf(func(i int) {})
	correctReturn := reflect.TypeOf(func(i int) string { return "" })
	aliasedTypeReturn := reflect.TypeOf(func(i int) compatibleType { return "" })
	tooManyReturn := reflect.TypeOf(func(i int) (string, int) { return "", i })

	ExpectThat(action.CheckType(emptyReturn), Not(Equals(nil)))
	ExpectThat(action.CheckType(correctReturn), Equals(nil))
	ExpectThat(action.CheckType(aliasedTypeReturn), Not(Equals(nil)))
	ExpectThat(action.CheckType(tooManyReturn), Not(Equals(nil)))
}

func (t *ReturnTest) MultipleValues() {
	someInt := 17

	// Invoke
	action := Return("taco", &someInt, 19)
	result := action.Invoke([]interface{}{})

	ExpectThat(len(result), Equals(3))
	ExpectThat(result[0], Equals("taco"))
	ExpectThat(result[1], Equals(&someInt))
	ExpectThat(result[2], Equals(19))

	// CheckType
	emptyReturn := reflect.TypeOf(func(i int) {})
	correctReturn := reflect.TypeOf(func(i int) (string, *int, int) { return "", &someInt, 19 })
	incorrectReturn := reflect.TypeOf(func(i int) (string, *int) { return "", &someInt })

	ExpectThat(action.CheckType(emptyReturn), Not(Equals(nil)))
	ExpectThat(action.CheckType(correctReturn), Equals(nil))
	ExpectThat(action.CheckType(incorrectReturn), Not(Equals(nil)))
}
