// Copyright 2012 Aaron Jacobs. All Rights Reserved.
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

// Package generate implements code generation for mock classes. This is an
// implementation detail of the createmock command, which you probably want to
// use directly instead.
package generate

import (
	"errors"
	"io"
	"reflect"
)

// Given a set of interfaces to mock, write out source code for a package named
// `pkg` that contains mock implementations of those interfaces.
func GenerateMockSource(w io.Writer, pkg string, interfaces []reflect.Type) error {
	return errors.New("Not implemented.")
}
