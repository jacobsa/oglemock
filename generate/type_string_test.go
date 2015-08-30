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

package generate

import (
	"reflect"
	"testing"

	"github.com/jacobsa/oglemock/createmock/test_cases/gcs"
	. "github.com/jacobsa/ogletest"
)

func TestTypeString(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Boilerplate
////////////////////////////////////////////////////////////////////////

type TypeStringTest struct {
}

func init() { RegisterTestSuite(&TypeStringTest{}) }

////////////////////////////////////////////////////////////////////////
// Test functions
////////////////////////////////////////////////////////////////////////

func (t *TypeStringTest) TestCases() {
	const gcsPkgPath = "github.com/jacobsa/oglemock/createmock/test_cases/gcs"

	testCases := []struct {
		t        reflect.Type
		pkgPath  string
		expected string
	}{
		// Scalar types
		0: {reflect.TypeOf(true), "", "bool"},
		1: {reflect.TypeOf(true), "some/pkg", "bool"},
		2: {reflect.TypeOf(int(17)), "some/pkg", "int"},
		3: {reflect.TypeOf(int32(17)), "some/pkg", "int32"},
		4: {reflect.TypeOf(uint(17)), "some/pkg", "uint"},
		5: {reflect.TypeOf(uint32(17)), "some/pkg", "uint32"},
		6: {reflect.TypeOf(uintptr(17)), "some/pkg", "uintptr"},
		7: {reflect.TypeOf(float32(17)), "some/pkg", "float32"},
		8: {reflect.TypeOf(complex64(17)), "some/pkg", "complex64"},

		// Structs
		9: {
			reflect.TypeOf(gcs.CreateObjectRequest{}),
			"some/pkg",
			"gcs.CreateObjectRequest",
		},

		10: {
			reflect.TypeOf(gcs.CreateObjectRequest{}),
			gcsPkgPath,
			"CreateObjectRequest",
		},
	}

	for i, tc := range testCases {
		ExpectEq(
			tc.expected,
			typeString(tc.t, tc.pkgPath),
			"Case %d: %v, %q", i, tc.t, tc.pkgPath)
	}
}
