// Copyright 2015 Aaron Jacobs. All Rights Reserved.
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
	"reflect"
	"testing"

	. "github.com/jacobsa/oglematchers"
	"github.com/jacobsa/oglemock"
	. "github.com/jacobsa/ogletest"
)

func TestSaveArg(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////
// Boilerplate
////////////////////////////////////////////////////////////

type SaveArgTest struct {
}

func init() { RegisterTestSuite(&SaveArgTest{}) }

////////////////////////////////////////////////////////////
// Test functions
////////////////////////////////////////////////////////////

func (t *SaveArgTest) FunctionHasNoArguments() {
	const index = 0
	var dst int
	f := func() (int, string) { return 0, "" }

	err := oglemock.SaveArg(0, &dst).SetSignature(reflect.TypeOf(f))
	ExpectThat(err, Error(HasSubstr("index 0")))
	ExpectThat(err, Error(HasSubstr("out of range")))
	ExpectThat(err, Error(HasSubstr("func() (int, string)")))
}

func (t *SaveArgTest) ArgumentIndexOutOfRange() {
	const index = 2
	var dst int
	f := func(a int, b int) {}

	err := oglemock.SaveArg(0, &dst).SetSignature(reflect.TypeOf(f))
	ExpectThat(err, Error(HasSubstr("index 2")))
	ExpectThat(err, Error(HasSubstr("out of range")))
	ExpectThat(err, Error(HasSubstr("func(int, int)")))
}

func (t *SaveArgTest) DestinationIsLiteralNil() {
	AssertFalse(true, "TODO")
}

func (t *SaveArgTest) DestinationIsNotAPointer() {
	AssertFalse(true, "TODO")
}

func (t *SaveArgTest) DestinationIsNilPointer() {
	AssertFalse(true, "TODO")
}

func (t *SaveArgTest) DestinationNotAssignable() {
	AssertFalse(true, "TODO")
}

func (t *SaveArgTest) ExactTypeMatch() {
	AssertFalse(true, "TODO")
}

func (t *SaveArgTest) AssignableTypeMatch() {
	AssertFalse(true, "TODO")
}
