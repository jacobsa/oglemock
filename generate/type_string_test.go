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
	"io"
	"reflect"
	"testing"
	"unsafe"

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
	to := reflect.TypeOf

	testCases := []struct {
		t        reflect.Type
		pkgPath  string
		expected string
	}{
		/////////////////////////
		// Scalar types
		/////////////////////////

		0: {to(true), "", "bool"},
		1: {to(true), "some/pkg", "bool"},
		2: {to(int(17)), "some/pkg", "int"},
		3: {to(int32(17)), "some/pkg", "int32"},
		4: {to(uint(17)), "some/pkg", "uint"},
		5: {to(uint32(17)), "some/pkg", "uint32"},
		6: {to(uintptr(17)), "some/pkg", "uintptr"},
		7: {to(float32(17)), "some/pkg", "float32"},
		8: {to(complex64(17)), "some/pkg", "complex64"},

		/////////////////////////
		// Structs
		/////////////////////////

		9:  {to(gcs.Object{}), "some/pkg", "gcs.Object"},
		10: {to(gcs.Object{}), gcsPkgPath, "Object"},

		/////////////////////////
		// Pointers
		/////////////////////////

		11: {to((*int)(nil)), gcsPkgPath, "*int"},
		12: {to((*gcs.Object)(nil)), "some/pkg", "*gcs.Object"},
		13: {to((*gcs.Object)(nil)), gcsPkgPath, "*Object"},

		/////////////////////////
		// Arrays
		/////////////////////////

		14: {to([3]int{}), "some/pkg", "[3]int"},
		15: {to([3]gcs.Object{}), gcsPkgPath, "[3]Object"},

		/////////////////////////
		// Channels
		/////////////////////////

		16: {to((chan int)(nil)), "some/pkg", "chan int"},
		17: {to((<-chan int)(nil)), "some/pkg", "<-chan int"},
		18: {to((chan<- int)(nil)), "some/pkg", "chan<- int"},
		19: {to((<-chan gcs.Object)(nil)), gcsPkgPath, "<-chan Object"},

		/////////////////////////
		// Functions
		/////////////////////////

		20: {
			to(func(int, gcs.Object) {}),
			gcsPkgPath,
			"func(int, Object) ()",
		},

		21: {
			to(func() (*gcs.Object, error) { return nil, nil }),
			gcsPkgPath,
			"func() (*Object, error)",
		},

		22: {
			to(func(int, gcs.Object) (*gcs.Object, error) { return nil, nil }),
			gcsPkgPath,
			"func(int, Object) (*Object, error)",
		},

		/////////////////////////
		// Interfaces
		/////////////////////////

		23: {to((*error)(nil)).Elem(), "some/pkg", "error"},
		24: {to((*io.Reader)(nil)).Elem(), "some/pkg", "io.Reader"},
		25: {to((*io.Reader)(nil)).Elem(), "io", "Reader"},

		/////////////////////////
		// Maps
		/////////////////////////

		26: {to(map[*gcs.Object]gcs.Object{}), gcsPkgPath, "map[*Object]Object"},

		/////////////////////////
		// Slices
		/////////////////////////

		27: {to([]int{}), "some/pkg", "[]int"},
		28: {to([]gcs.Object{}), gcsPkgPath, "[]Object"},

		/////////////////////////
		// Strings
		/////////////////////////

		29: {to(""), gcsPkgPath, "string"},

		/////////////////////////
		// Unsafe pointer
		/////////////////////////

		30: {to(unsafe.Pointer(nil)), gcsPkgPath, "unsafe.Pointer"},

		/////////////////////////
		// Named types
		/////////////////////////

		31: {to(gcs.Int(17)), "some/pkg", "gcs.Int"},
		32: {to(gcs.Int(17)), gcsPkgPath, "Int"},

		33: {to(gcs.Array{}), "some/pkg", "gcs.Array"},
		34: {to(gcs.Array{}), gcsPkgPath, "Array"},

		35: {to(gcs.Chan(nil)), "some/pkg", "gcs.Chan"},
		36: {to(gcs.Chan(nil)), gcsPkgPath, "Chan"},

		37: {to(gcs.Ptr(nil)), "some/pkg", "gcs.Ptr"},
		38: {to(gcs.Ptr(nil)), gcsPkgPath, "Ptr"},

		39: {to((*gcs.Int)(nil)), "some/pkg", "*gcs.Int"},
		40: {to((*gcs.Int)(nil)), gcsPkgPath, "*Int"},
	}

	for i, tc := range testCases {
		ExpectEq(
			tc.expected,
			typeString(tc.t, tc.pkgPath),
			"Case %d: %v, %q", i, tc.t, tc.pkgPath)
	}
}
