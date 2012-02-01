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
	"math"
	"reflect"
	"testing"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

var someInt int = 17

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
		ExpectThat(res[0], IdenticalTo(c.expectedVal), "Test case %d: %v", i, c)
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

func (t *ReturnTest) MultipleReturnValues() {
	sig := reflect.TypeOf(func() (int, string) { return 0, "" })
	var a oglemock.Action
	var err error
	var vals []interface{}

	// No values.
	a = oglemock.Return()
	err = a.CheckType(sig)
	ExpectThat(err, Error(HasSubstr("given 0 vals")))
	ExpectThat(err, Error(HasSubstr("expected 2")))

	// One value.
	a = oglemock.Return(17)
	err = a.CheckType(sig)
	ExpectThat(err, Error(HasSubstr("given 1 val")))
	ExpectThat(err, Error(HasSubstr("expected 2")))

	// Two values.
	a = oglemock.Return(17, "taco")
	err = a.CheckType(sig)
	AssertEq(nil, err)

	vals = a.Invoke([]interface{}{})
	ExpectThat(vals, ElementsAre(IdenticalTo(int(17)), "taco"))
}

func (t *ReturnTest) Bool() {
	type namedType bool

	sig := reflect.TypeOf(func() bool { return false })
	cases := []returnTestCase{
		// Identical types.
		{ bool(true), bool(true), true, "" },
		{ bool(false), bool(false), true, "" },

		// Named version of same underlying type.
		{ namedType(true), bool(true), true, "" },

		// Wrong types.
		{ nil, nil, false, "given <nil>" },
		{ int16(1), nil, false, "given int" },
		{ float64(1), nil, false, "given float64" },
		{ complex128(1), nil, false, "given complex128" },
		{ &someInt, nil, false, "given *int" },
		{ make(chan int), nil, false, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Int() {
	type namedType int

	sig := reflect.TypeOf(func() int { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ int(0), int(0), true, "" },
		{ int(math.MaxInt32), int(math.MaxInt32), true, "" },

		// Named version of same underlying type.
		{ namedType(17), int(17), true, "" },

		// Wrong types.
		{ nil, nil, false, "given <nil>" },
		{ int16(1), nil, false, "given int16" },
		{ float64(1), nil, false, "given float64" },
		{ complex128(1), nil, false, "given complex128" },
		{ &someInt, nil, false, "given *int" },
		{ make(chan int), nil, false, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Int8() {
	type namedType int8

	sig := reflect.TypeOf(func() int8 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ int8(0), int8(0), true, "" },
		{ int8(math.MaxInt8), int8(math.MaxInt8), true, "" },

		// Named version of same underlying type.
		{ namedType(17), int8(17), true, "" },

		// In-range ints.
		{ int(17), int8(17), true, "" },
		{ int(math.MaxInt8), int8(math.MaxInt8), true, "" },

		// Out of range ints.
		{ int(math.MaxInt8 + 1), nil, false, "out of range" },

		// Wrong types.
		{ nil, nil, false, "given <nil>" },
		{ int16(1), nil, false, "given int16" },
		{ float64(1), nil, false, "given float64" },
		{ complex128(1), nil, false, "given complex128" },
		{ &someInt, nil, false, "given *int" },
		{ make(chan int), nil, false, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Int16() {
	type namedType int16

	sig := reflect.TypeOf(func() int16 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ int16(0), int16(0), true, "" },
		{ int16(math.MaxInt16), int16(math.MaxInt16), true, "" },

		// Named version of same underlying type.
		{ namedType(17), int16(17), true, "" },

		// In-range ints.
		{ int(17), int16(17), true, "" },
		{ int(math.MaxInt16), int16(math.MaxInt16), true, "" },

		// Out of range ints.
		{ int(math.MaxInt16 + 1), nil, false, "out of range" },

		// Wrong types.
		{ nil, nil, false, "given <nil>" },
		{ int8(1), nil, false, "given int8" },
		{ float64(1), nil, false, "given float64" },
		{ complex128(1), nil, false, "given complex128" },
		{ &someInt, nil, false, "given *int" },
		{ make(chan int), nil, false, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Int32() {
	type namedType int32

	sig := reflect.TypeOf(func() int32 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ int32(0), int32(0), true, "" },
		{ int32(math.MaxInt32), int32(math.MaxInt32), true, "" },

		// Named version of same underlying type.
		{ namedType(17), int32(17), true, "" },

		// In-range ints.
		{ int(17), int32(17), true, "" },
		{ int(math.MaxInt32), int32(math.MaxInt32), true, "" },

		// Wrong types.
		{ nil, nil, false, "given <nil>" },
		{ int16(1), nil, false, "given int16" },
		{ float64(1), nil, false, "given float64" },
		{ complex128(1), nil, false, "given complex128" },
		{ &someInt, nil, false, "given *int" },
		{ make(chan int), nil, false, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Rune() {
	type namedType rune

	sig := reflect.TypeOf(func() rune { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ rune(0), rune(0), true, "" },
		{ rune(math.MaxInt32), rune(math.MaxInt32), true, "" },

		// Named version of same underlying type.
		{ namedType(17), rune(17), true, "" },

		// Aliased version of type.
		{ int32(17), rune(17), true, "" },

		// In-range ints.
		{ int(17), rune(17), true, "" },
		{ int(math.MaxInt32), rune(math.MaxInt32), true, "" },

		// Wrong types.
		{ nil, nil, false, "given <nil>" },
		{ int16(1), nil, false, "given int16" },
		{ float64(1), nil, false, "given float64" },
		{ complex128(1), nil, false, "given complex128" },
		{ &someInt, nil, false, "given *int" },
		{ make(chan int), nil, false, "given chan int" },
	}

	t.runTestCases(sig, cases)
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

func (t *ReturnTest) Byte() {
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

func (t *ReturnTest) NamedNumericType() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) NamedNonNumericType() {
	ExpectTrue(false, "TODO")
}

func (t *ReturnTest) NamedChannelType() {
	ExpectTrue(false, "TODO")
}
