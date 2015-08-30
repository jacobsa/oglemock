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

import "reflect"

// Return the string that should be used to refer to the supplied type within
// the given package.
//
// For example, a pointer to an io.Reader may be rendered as "*Reader" or
// "*io.Reader" depending on whether the package path is "io" or not.
func typeString(
	t reflect.Type,
	pkgPath string) (s string) {
	s = t.String()
	return
}
