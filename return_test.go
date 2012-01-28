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

	emptyReturn := reflect.TypeOf(func(i int) {})
	stringReturn := reflect.TypeOf(func(i int) string { return "" })
	interfaceReturn := reflect.TypeOf(func(i int) error { return nil })
	var err error

	// No return value.
	err = action.CheckType(emptyReturn)
	ExpectEq(nil, err)

	// String return value.
	err = action.CheckType(stringReturn)
	ExpectThat(err, Error(HasSubstr("0 vals; expected 1")))

	// Interface return value.
	err = action.CheckType(interfaceReturn)
	ExpectThat(err, Error(HasSubstr("0 vals; expected 1")))
}

func (t *ReturnTest) StringValue() {
	action := Return("taco")

	// Invoke
	result := action.Invoke([]interface{}{})

	ExpectThat(len(result), Equals(1))
	ExpectThat(result[0], Equals("taco"))

	type compatibleType string
	emptyReturn := reflect.TypeOf(func() {})
	stringReturn := reflect.TypeOf(func() string { return "" })
	aliasedTypeReturn := reflect.TypeOf(func() compatibleType { return "" })
	intReturn := reflect.TypeOf(func() int { return 0 })
	unsatisfiedInterfaceReturn := reflect.TypeOf(func() error { return nil })
	tooManyReturn := reflect.TypeOf(func() (string, int) { return "", 0 })
	var err error

	// No return value.
	err = action.CheckType(emptyReturn)
	ExpectThat(err, Error(HasSubstr("1 vals; expected 0")))

	// String return value.
	err = action.CheckType(stringReturn)
	ExpectEq(nil, err)

	// Aliased string return value.
	err = action.CheckType(aliasedTypeReturn)
	ExpectEq(nil, err)

	// Int return value.
	err = action.CheckType(intReturn)
	ExpectThat(err, Error(HasSubstr("val 0")))
	ExpectThat(err, Error(HasSubstr("given string")))
	ExpectThat(err, Error(HasSubstr("expected int")))

	// Unsatisfied interface return value.
	err = action.CheckType(unsatisfiedInterfaceReturn)
	ExpectThat(err, Error(HasSubstr("val 0")))
	ExpectThat(err, Error(HasSubstr("given string")))
	ExpectThat(err, Error(HasSubstr("expected error")))

	// Multiple return values.
	err = action.CheckType(tooManyReturn)
	ExpectThat(err, Error(HasSubstr("1 vals; expected 2")))
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
