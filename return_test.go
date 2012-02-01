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
	. "github.com/jacobsa/ogletest"
	"github.com/jacobsa/oglemock"
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

type returnTestCase struct {
	suppliedVal interface{}
	expectedVal interface{}
	expectedCheckTypeResult bool
	expectedCheckTypeErrorSubstring string
}

func (t *ReturnTest) runTestCases(signature reflect.Type, cases []returnTestCase) {
	for i, c := range cases {
		a := oglemock.Return(c.suppliedVal)

		// CheckType
		err := a.CheckType(signature)
		if c.expectedCheckTypeResult {
			ExpectEq(nil, err, "Test case %d: %v", i, c)
		} else {
			ExpectThat(err, Error(HasSubstr(c.expectedCheckTypeErrorSubstring)),
				"Test case %d: %v", i, c)
			continue
		}

		// Invoke
		res := a.Invoke([]interface{}{})
		AssertThat(res, ElementsAre(Any()))
		ExpectThat(res[0], IdenticalTo(c.expectedVal))
	}
}

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *ReturnTest) NoReturnValues() {
	sig := reflect.TypeOf(func() {})
	var a oglemock.Action
	var err error
	var vals []interface{}

	// No values.
	a = oglemock.Return()
	err = a.CheckType(sig)
	AssertEq(nil, err)

	vals = a.Invoke([]interface{}{})
	ExpectThat(vals, ElementsAre())

	// One value.
	a = oglemock.Return(17)
	err = a.CheckType(sig)
	ExpectThat(err, Error(HasSubstr("given 1 val")))
	ExpectThat(err, Error(HasSubstr("expected 0")))

	// Two values.
	a = oglemock.Return(17, 19)
	err = a.CheckType(sig)
	ExpectThat(err, Error(HasSubstr("given 2 vals")))
	ExpectThat(err, Error(HasSubstr("expected 0")))
}

func (t *ReturnTest) Bool() {
	sig := reflect.TypeOf(func() bool { return false })
	cases := []returnTestCase {
		{ true, true, true, "" },
		{ false, false, true, "" },
		{ nil, false, false, "given <nil>; expected bool" },
		{ int(17), false, false, "given int; expected bool" },
		{ "false", false, false, "given string; expected bool" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Int() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Int8() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Int16() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Int32() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Int64() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Uint() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Uint8() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Uint16() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Uint32() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Uint64() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Uintptr() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Float32() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Float64() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Complex64() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Complex128() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) ArrayOfInt() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) ChanOfInt() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Func() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Interface() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) MapFromStringToInt() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) PointerToString() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) SliceOfInts() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) String() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) Struct() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) UnsafePointer() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) MultipleReturnValues() {
	ExpectTrue(false, "TODO")
}
