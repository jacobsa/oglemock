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
	"bytes"
	"io"
	"math"
	"reflect"
	"testing"
	"unsafe"
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
	expectedSetSignatureErrorSubstring string
}

func (t *ReturnTest) runTestCases(signature reflect.Type, cases []returnTestCase) {
	for i, c := range cases {
		a := oglemock.Return(c.suppliedVal)

		// SetSignature
		err := a.SetSignature(signature)
		if c.expectedSetSignatureErrorSubstring == "" {
			ExpectEq(nil, err, "Test case %d: %v", i, c)
		} else {
			ExpectThat(err, Error(HasSubstr(c.expectedSetSignatureErrorSubstring)),
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

func (t *ReturnTest) SetSignatureNotCalled() {
	a := oglemock.Return()
	f := func() { a.Invoke([]interface{}{}) }
	ExpectThat(f, Panics(MatchesRegexp("SetSignature.*called")))
}

func (t *ReturnTest) NoReturnValues() {
	sig := reflect.TypeOf(func() {})
	var a oglemock.Action
	var err error
	var vals []interface{}

	// No values.
	a = oglemock.Return()
	err = a.SetSignature(sig)
	AssertEq(nil, err)

	vals = a.Invoke([]interface{}{})
	ExpectThat(vals, ElementsAre())

	// One value.
	a = oglemock.Return(17)
	err = a.SetSignature(sig)
	ExpectThat(err, Error(HasSubstr("given 1 val")))
	ExpectThat(err, Error(HasSubstr("expected 0")))

	// Two values.
	a = oglemock.Return(17, 19)
	err = a.SetSignature(sig)
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
	err = a.SetSignature(sig)
	ExpectThat(err, Error(HasSubstr("given 0 vals")))
	ExpectThat(err, Error(HasSubstr("expected 2")))

	// One value.
	a = oglemock.Return(17)
	err = a.SetSignature(sig)
	ExpectThat(err, Error(HasSubstr("given 1 val")))
	ExpectThat(err, Error(HasSubstr("expected 2")))

	// Two values.
	a = oglemock.Return(17, "taco")
	err = a.SetSignature(sig)
	AssertEq(nil, err)

	vals = a.Invoke([]interface{}{})
	ExpectThat(vals, ElementsAre(IdenticalTo(int(17)), "taco"))
}

func (t *ReturnTest) Bool() {
	sig := reflect.TypeOf(func() bool { return false })
	cases := []returnTestCase{
		// Identical types.
		{ bool(true), bool(true), "" },
		{ bool(false), bool(false), "" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Int() {
	type namedType int

	sig := reflect.TypeOf(func() int { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ int(math.MinInt32), int(math.MinInt32), "" },
		{ int(math.MaxInt32), int(math.MaxInt32), "" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int16" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Int8() {
	type namedType int8

	sig := reflect.TypeOf(func() int8 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ int8(math.MinInt8), int8(math.MinInt8), "" },
		{ int8(math.MaxInt8), int8(math.MaxInt8), "" },

		// In-range ints.
		{ int(math.MinInt8), int8(math.MinInt8), "" },
		{ int(math.MaxInt8), int8(math.MaxInt8), "" },

		// Out of range ints.
		{ int(math.MinInt8 - 1), nil, "out of range" },
		{ int(math.MaxInt8 + 1), nil, "out of range" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int16" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Int16() {
	type namedType int16

	sig := reflect.TypeOf(func() int16 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ int16(math.MinInt16), int16(math.MinInt16), "" },
		{ int16(math.MaxInt16), int16(math.MaxInt16), "" },

		// In-range ints.
		{ int(math.MinInt16), int16(math.MinInt16), "" },
		{ int(math.MaxInt16), int16(math.MaxInt16), "" },

		// Out of range ints.
		{ int(math.MinInt16 - 1), nil, "out of range" },
		{ int(math.MaxInt16 + 1), nil, "out of range" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int8(1), nil, "given int8" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Int32() {
	type namedType int32

	sig := reflect.TypeOf(func() int32 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ int32(math.MinInt32), int32(math.MinInt32), "" },
		{ int32(math.MaxInt32), int32(math.MaxInt32), "" },

		// Aliased version of type.
		{ rune(17), int32(17), "" },

		// In-range ints.
		{ int(math.MinInt32), int32(math.MinInt32), "" },
		{ int(math.MaxInt32), int32(math.MaxInt32), "" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int16" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Rune() {
	type namedType rune

	sig := reflect.TypeOf(func() rune { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ rune(math.MinInt32), rune(math.MinInt32), "" },
		{ rune(math.MaxInt32), rune(math.MaxInt32), "" },

		// Aliased version of type.
		{ int32(17), rune(17), "" },

		// In-range ints.
		{ int(math.MinInt32), rune(math.MinInt32), "" },
		{ int(math.MaxInt32), rune(math.MaxInt32), "" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int16" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Int64() {
	type namedType int64

	sig := reflect.TypeOf(func() int64 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ int64(math.MinInt64), int64(math.MinInt64), "" },
		{ int64(math.MaxInt64), int64(math.MaxInt64), "" },

		// In-range ints.
		{ int(math.MinInt32), int64(math.MinInt32), "" },
		{ int(math.MaxInt32), int64(math.MaxInt32), "" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int16" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Uint() {
	type namedType uint

	sig := reflect.TypeOf(func() uint { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ uint(0), uint(0), "" },
		{ uint(math.MaxUint32), uint(math.MaxUint32), "" },

		// In-range ints.
		{ int(0), uint(0), "" },
		{ int(math.MaxInt32), uint(math.MaxUint32), "" },

		// Out of range ints.
		{ int(-1), nil, "out of range" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int16" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Uint8() {
	type namedType uint8

	sig := reflect.TypeOf(func() uint8 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ uint8(0), uint8(0), "" },
		{ uint8(math.MaxUint8), uint8(math.MaxUint8), "" },

		// Aliased version of type.
		{ byte(17), uint8(17), "" },

		// In-range ints.
		{ int(0), uint8(0), "" },
		{ int(math.MaxUint8), uint8(math.MaxUint8), "" },

		// Out of range ints.
		{ int(-1), nil, "out of range" },
		{ int(math.MaxUint8 + 1), nil, "out of range" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int16" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Byte() {
	type namedType byte

	sig := reflect.TypeOf(func() byte { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ byte(0), byte(0), "" },
		{ byte(math.MaxUint8), byte(math.MaxUint8), "" },

		// Aliased version of type.
		{ uint8(17), byte(17), "" },

		// In-range ints.
		{ int(0), byte(0), "" },
		{ int(math.MaxUint8), byte(math.MaxUint8), "" },

		// Out of range ints.
		{ int(-1), nil, "out of range" },
		{ int(math.MaxUint8 + 1), nil, "out of range" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int16" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Uint16() {
	type namedType uint16

	sig := reflect.TypeOf(func() uint16 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ uint32(0), uint16(0), "" },
		{ uint32(math.MaxUint16), uint16(math.MaxUint16), "" },

		// In-range ints.
		{ int(0), uint16(0), "" },
		{ int(math.MaxUint16), uint16(math.MaxUint16), "" },

		// Out of range ints.
		{ int(-1), nil, "out of range" },
		{ int(math.MaxUint16 + 1), nil, "out of range" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int16" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Uint32() {
	type namedType uint32

	sig := reflect.TypeOf(func() uint32 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ uint32(0), uint32(0), "" },
		{ uint32(math.MaxUint32), uint32(math.MaxUint32), "" },

		// In-range ints.
		{ int(0), uint32(0), "" },
		{ int(math.MaxInt32), uint32(math.MaxInt32), "" },

		// Out of range ints.
		{ int(-1), nil, "out of range" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int16" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Uint64() {
	type namedType uint64

	sig := reflect.TypeOf(func() uint64 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ uint64(0), uint64(0), "" },
		{ uint64(math.MaxUint64), uint64(math.MaxUint64), "" },

		// In-range ints.
		{ int(0), uint64(0), "" },
		{ int(math.MaxInt32), uint32(math.MaxInt32), "" },

		// Out of range ints.
		{ int(-1), nil, "out of range" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int16" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Uintptr() {
	type namedType uintptr

	sig := reflect.TypeOf(func() uintptr { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ uintptr(17), uintptr(17), "" },

		// Named version of same underlying type.
		{ namedType(17), uintptr(17), "" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Float32() {
	type namedType float32

	sig := reflect.TypeOf(func() float32 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ float32(-17.5), float32(-17.5), "" },
		{ float32(17.5), float32(17.5), "" },

		// Named version of same underlying type.
		{ namedType(17.5), float32(17.5), "" },

		// In-range ints.
		{ int(-17), float32(-17), "" },
		{ int(17), float32(17), "" },

		// Float64s
		{ float64(-17.5), float32(-17.5), "" },
		{ float64(17.5), float32(17.5), "" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int8" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Float64() {
	type namedType float64

	sig := reflect.TypeOf(func() float64 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ float64(-17.5), float64(-17.5), "" },
		{ float64(17.5), float64(17.5), "" },

		// Named version of same underlying type.
		{ namedType(17.5), float64(17.5), "" },

		// In-range ints.
		{ int(-17), float64(-17), "" },
		{ int(17), float64(17), "" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int8" },
		{ float32(1), nil, "given float32" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Complex64() {
	type namedType complex64

	sig := reflect.TypeOf(func() complex64 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ complex64(-17.5-1i), complex64(-17.5-1i), "" },
		{ complex64(17.5+1i), complex64(17.5+1i), "" },

		// Named version of same underlying type.
		{ namedType(17.5+1i), complex64(17.5+1i), "" },

		// In-range ints.
		{ int(-17), complex64(-17), "" },
		{ int(17), complex64(17), "" },

		// Float64s
		{ float64(-17.5), complex64(-17.5), "" },
		{ float64(17.5), complex64(17.5), "" },

		// Complex128s
		{ complex128(-17.5-1i), complex64(-17.5-1i), "" },
		{ complex128(17.5+1i), complex64(17.5+1i), "" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int8" },
		{ float32(1), nil, "given float32" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Complex128() {
	type namedType complex128

	sig := reflect.TypeOf(func() complex128 { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ complex128(-17.5-1i), complex128(-17.5-1i), "" },
		{ complex128(17.5+1i), complex128(17.5+1i), "" },

		// Named version of same underlying type.
		{ namedType(17.5+1i), complex128(17.5+1i), "" },

		// In-range ints.
		{ int(-17), complex128(-17), "" },
		{ int(17), complex128(17), "" },

		// Float64s
		{ float64(-17.5), complex128(-17.5), "" },
		{ float64(17.5), complex128(17.5), "" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int16(1), nil, "given int8" },
		{ float32(1), nil, "given float32" },
		{ complex64(1), nil, "given complex64" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) ArrayOfInt() {
	type namedType [2]int
	type namedElemType int

	sig := reflect.TypeOf(func() [2]int { return [2]int{0, 0} })
	cases := []returnTestCase{
		// Identical types.
		{ [2]int{19, 23}, [2]int{19, 23}, "" },

		// Named version of same underlying type.
		{ namedType{19, 23}, [2]int{19, 23}, "" },

		// Wrong length.
		{ [1]int{17}, nil, "given [1]int" },

		// Wrong element types.
		{ [2]namedElemType{19, 23}, nil, "given [2]namedElemType" },
		{ [2]string{"", ""}, nil, "given [2]string" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) ChanOfInt() {
	type namedType chan int
	type namedElemType int

	someChan := make(chan int)
	someNamedTypeChan := namedType(make(namedType))

	sig := reflect.TypeOf(func() chan int { return nil })
	cases := []returnTestCase{
		// Identical types.
		{ someChan, someChan, "" },

		// Nil values.
		{ (interface{})(nil), (chan int)(nil), "" },
		{ (chan int)(nil), (chan int)(nil), "" },

		// Named version of same underlying type.
		{ someNamedTypeChan, (chan int)(someNamedTypeChan), "" },

		// Wrong element types.
		{ make(chan string), nil, "given chan string" },
		{ make(chan namedElemType), nil, "given chan namedElemType" },

		// Wrong direction
		{ (<-chan int)(someChan), nil, "given <-chan int" },
		{ (chan<- int)(someChan), nil, "given chan<- int" },

		// Wrong types.
		{ (func())(nil), nil, "given func()" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) SendChanOfInt() {
	type namedType chan<- int
	type namedElemType int

	someChan := make(chan<- int)
	someBidirectionalChannel := make(chan int)
	someNamedTypeChan := namedType(make(namedType))

	sig := reflect.TypeOf(func() chan<- int { return nil })
	cases := []returnTestCase{
		// Identical types.
		{ someChan, someChan, "" },

		// Nil values.
		{ (interface{})(nil), (chan<- int)(nil), "" },
		{ (chan int)(nil), (chan<- int)(nil), "" },

		// Named version of same underlying type.
		{ someNamedTypeChan, (chan<- int)(someNamedTypeChan), "" },

		// Bidirectional channel
		{ someBidirectionalChannel, (chan<- int)(someBidirectionalChannel), "" },

		// Wrong direction
		{ (<-chan int)(someBidirectionalChannel), nil, "given <-chan int" },

		// Wrong element types.
		{ make(chan string), nil, "given chan string" },
		{ make(chan namedElemType), nil, "given chan namedElemType" },

		// Wrong types.
		{ (func())(nil), nil, "given func()" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) RecvChanOfInt() {
	type namedType <-chan int
	type namedElemType int

	someChan := make(<-chan int)
	someBidirectionalChannel := make(chan int)
	someNamedTypeChan := namedType(make(namedType))

	sig := reflect.TypeOf(func() <-chan int { return nil })
	cases := []returnTestCase{
		// Identical types.
		{ someChan, someChan, "" },

		// Nil values.
		{ (interface{})(nil), (<-chan int)(nil), "" },
		{ (chan int)(nil), (<-chan int)(nil), "" },

		// Named version of same underlying type.
		{ someNamedTypeChan, (<-chan int)(someNamedTypeChan), "" },

		// Bidirectional channel
		{ someBidirectionalChannel, (<-chan int)(someBidirectionalChannel), "" },

		// Wrong direction
		{ (chan<- int)(someBidirectionalChannel), nil, "given chan<- int" },

		// Wrong element types.
		{ make(chan string), nil, "given chan string" },
		{ make(chan namedElemType), nil, "given chan namedElemType" },

		// Wrong types.
		{ (func())(nil), nil, "given func()" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Func() {
	type namedType func(string) int
	someFunc := func(string) int { return 0 }

	sig := reflect.TypeOf(func() func(string) int { return nil })
	cases := []returnTestCase{
		// Identical types.
		{ someFunc, someFunc, "" },

		// Nil values.
		{ (interface{})(nil), (func(string) int)(nil), "" },
		{ (func(string) int)(nil), (func(string) int)(nil), "" },

		// Named version of same underlying type.
		{ namedType(someFunc), someFunc, "" },

		// Wrong parameter and return types.
		{ func(int) int { return 0 }, nil, "given func(int) int" },
		{ func(string) string { return "" }, nil, "given func(string) string" },

		// Wrong types.
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ (chan int)(nil), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Interface() {
	sig := reflect.TypeOf(func() io.Reader { return nil })

	someBuffer := new(bytes.Buffer)

	cases := []returnTestCase{
		// Type that implements interface.
		{ someBuffer, someBuffer, "" },

		// Nil value.
		{ (interface{})(nil), (interface{})(nil), "" },

		// Non-implementing types.
		{ (chan int)(nil), (chan int)(nil), "" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) MapFromStringToInt() {
	type namedType map[string]int
	type namedElemType string

	someMap := make(map[string]int)
	someNamedTypeMap := namedType(make(namedType))

	sig := reflect.TypeOf(func() map[string]int { return nil })
	cases := []returnTestCase{
		// Identical types.
		{ someMap, someMap, "" },

		// Nil values.
		{ (interface{})(nil), (chan int)(nil), "" },
		{ (map[string]int)(nil), (map[string]int)(nil), "" },

		// Named version of same underlying type.
		{ someNamedTypeMap, map[string]int(someNamedTypeMap), "" },

		// Wrong element types.
		{ make(map[int]int), nil, "given map[int]int" },
		{ make(map[namedElemType]int), nil, "given map[namedElemType]int" },
		{ make(map[string]string), nil, "given map[string]string" },

		// Wrong types.
		{ (func())(nil), nil, "given func()" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) PointerToString() {
	type namedType *string
	type namedElemType string

	someStr := ""
	someNamedStr := namedElemType("")

	sig := reflect.TypeOf(func() *string { return nil })
	cases := []returnTestCase{
		// Identical types.
		{ (*string)(&someStr), (*string)(&someStr), "" },

		// Nil values.
		{ (interface{})(nil), (*string)(nil), "" },
		{ (*string)(nil), (*string)(nil), "" },

		// Named version of same underlying type.
		{ namedType(&someStr), (*string)(&someStr), "" },

		// Wrong element types.
		{ &someInt, nil, "given *int" },
		{ &someNamedStr, nil, "given *oglematchers_test.namedElemType" },

		// Wrong types.
		{ (func())(nil), nil, "given func()" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ unsafe.Pointer(&someStr), nil, "given unsafe.Pointer" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) SliceOfInts() {
	type namedType []int
	type namedElemType int

	someSlice := make([]int, 1)

	sig := reflect.TypeOf(func() []int { return nil })
	cases := []returnTestCase{
		// Identical types.
		{ someSlice, someSlice, "" },

		// Nil values.
		{ (interface{})(nil), ([]int)(nil), "" },
		{ ([]int)(nil), ([]int)(nil), "" },

		// Named version of same underlying type.
		{ namedType(someSlice), namedType(someSlice), "" },

		// Wrong element types.
		{ make([]string, 1), nil, "given []string" },
		{ make([]namedElemType, 1), nil, "given []oglematchers_test.namedElemType" },

		// Wrong types.
		{ (func())(nil), nil, "given func()" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) String() {
	type namedType string

	sig := reflect.TypeOf(func() string { return "" })
	cases := []returnTestCase{
		// Identical types.
		{ string(""), string(""), "" },
		{ string("taco"), string("taco"), "" },

		// Named version of same underlying type.
		{ namedType("burrito"), string("burrito"), "" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) Struct() {
	type myStruct struct {
		a int
	}

	type otherStruct struct {}
	type namedType myStruct

	sig := reflect.TypeOf(func() myStruct { return myStruct{0} })
	cases := []returnTestCase{
		// Identical types.
		{ myStruct{17}, myStruct{17}, "" },

		// Named version of same underlying type.
		{ myStruct{17}, nil, "given oglematchers_test.namedType" },

		// Wrong field types.
		{ otherStruct{}, nil, "given oglematchers_test.otherStruct" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) UnsafePointer() {
	type namedType unsafe.Pointer

	someStr := ""

	sig := reflect.TypeOf(func() unsafe.Pointer { return nil })
	cases := []returnTestCase{
		// Identical types.
		{ unsafe.Pointer(&someStr), (*string)(&someStr), "" },

		// Nil values.
		{ (interface{})(nil), unsafe.Pointer(nil), "" },
		{ unsafe.Pointer(nil), unsafe.Pointer(nil), "" },

		// Named version of same underlying type.
		{ namedType(&someStr), nil, "given unsafe.Pointer" },

		// Wrong types.
		{ (func())(nil), nil, "given func()" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ (*string)(&someStr), nil, "given *string" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) NamedNumericType() {
	type namedType int16
	type otherNamedType int16

	sig := reflect.TypeOf(func() namedType { return 0 })
	cases := []returnTestCase{
		// Identical types.
		{ namedType(math.MinInt16), namedType(math.MinInt16), "" },
		{ namedType(math.MaxInt16), namedType(math.MaxInt16), "" },

		// Non-named version of same type.
		{ int16(17), namedType(17), "" },

		// Other named version of same underlying type.
		{ otherNamedType(17), nil, "given oglematchers_test.otherNamedType" },

		// In-range ints.
		{ int(math.MinInt16), namedType(math.MinInt16), "" },
		{ int(math.MaxInt16), namedType(math.MaxInt16), "" },

		// Out of range ints.
		{ int(math.MinInt16 - 1), nil, "out of range" },
		{ int(math.MaxInt16 + 1), nil, "out of range" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int8(1), nil, "given int8" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) NamedNonNumericType() {
	type namedType string
	type otherNamedType string

	sig := reflect.TypeOf(func() namedType { return "" })
	cases := []returnTestCase{
		// Identical types.
		{ namedType("taco"), namedType("taco"), "" },

		// Non-named version of same type.
		{ string("taco"), namedType("taco"), "" },

		// Other named version of same underlying type.
		{ otherNamedType(""), nil, "given oglematchers_test.otherNamedType" },

		// Wrong types.
		{ nil, nil, "given <nil>" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
		{ make(chan int), nil, "given chan int" },
	}

	t.runTestCases(sig, cases)
}

func (t *ReturnTest) NamedChannelType() {
	type namedType chan int
	type otherNamedType chan int

	someChan := make(namedType)
	someUnnamedTypeChan := make(chan int)
	someOtherNamedTypeChan := make(otherNamedType)

	sig := reflect.TypeOf(func() namedType { return nil })
	cases := []returnTestCase{
		// Identical types.
		{ someChan, someChan, "" },

		// Non-named version of same type.
		{ someUnnamedTypeChan, namedType(someUnnamedTypeChan), "" },

		// Other named version of same underlying type.
		{ someOtherNamedTypeChan, nil, "given oglematchers_test.otherNamedType" },

		// Nil values.
		{ (interface{})(nil), namedType(nil), "" },
		{ namedType(nil), namedType(nil), "" },

		// Wrong element types.
		{ make(chan string), nil, "given chan string" },

		// Wrong direction
		{ (<-chan int)(someChan), nil, "given <-chan int" },
		{ (chan<- int)(someChan), nil, "given chan<- int" },

		// Wrong types.
		{ (func())(nil), nil, "given func()" },
		{ int(1), nil, "given int" },
		{ float64(1), nil, "given float64" },
		{ complex128(1), nil, "given complex128" },
		{ &someInt, nil, "given *int" },
	}

	t.runTestCases(sig, cases)
}
