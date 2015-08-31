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
	"fmt"
	"log"
	"reflect"
)

// Return the string that should be used to refer to the supplied type within
// the given package.
//
// For example, a pointer to an io.Reader may be rendered as "*Reader" or
// "*io.Reader" depending on whether the package path is "io" or not.
func typeString(
	t reflect.Type,
	pkgPath string) (s string) {
	// Is this type named? If so we use its name, possibly with a package prefix.
	//
	// Examples:
	//
	//     int
	//     string
	//     error
	//     gcs.Bucket
	//
	if t.Name() != "" {
		if t.PkgPath() == pkgPath {
			s = t.Name()
		} else {
			s = t.String()
		}

		return
	}

	// This type is unnamed. Recurse.
	switch t.Kind() {
	case reflect.Array:
		s = fmt.Sprintf("[%d]%s", t.Len(), typeString(t.Elem(), pkgPath))

	case reflect.Chan:
		s = fmt.Sprintf("%s %s", t.ChanDir(), typeString(t.Elem(), pkgPath))

	case reflect.Func:
		s = funcTypeString(t)

	case reflect.Map:
		s = fmt.Sprintf(
			"map[%s]%s",
			typeString(t.Key(), pkgPath),
			typeString(t.Elem(), pkgPath))

	case reflect.Ptr:
		s = fmt.Sprintf("*%s", typeString(t.Elem(), pkgPath))

	case reflect.Slice:
		s = fmt.Sprintf("[]%s", typeString(t.Elem(), pkgPath))

	default:
		log.Panicf("Unhandled kind: %v", t.Kind())
	}

	return
}

func funcTypeString(t reflect.Type) (s string) {
	s = "TODO"
	return
}
