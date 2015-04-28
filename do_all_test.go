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

func TestDoAll(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////
// Boilerplate
////////////////////////////////////////////////////////////

type DoAllTest struct {
}

func init() { RegisterTestSuite(&DoAllTest{}) }

////////////////////////////////////////////////////////////
// Test functions
////////////////////////////////////////////////////////////

func (t *DoAllTest) FirstActionDoesntLikeSignature() {
	f := func(a int, b string) {}

	a0 := oglemock.Invoke(func() {})
	a1 := oglemock.Invoke(f)
	a2 := oglemock.Return()

	err := oglemock.DoAll(a0, a1, a2).SetSignature(reflect.TypeOf(f))
	ExpectThat(err, Error(HasSubstr("TODO")))
}

func (t *DoAllTest) LastActionDoesntLikeSignature() {
	AssertFalse(true, "TODO")
}

func (t *DoAllTest) SingleAction() {
	AssertFalse(true, "TODO")
}

func (t *DoAllTest) MultipleActions() {
	AssertFalse(true, "TODO")
}
