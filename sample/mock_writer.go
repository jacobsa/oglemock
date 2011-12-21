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

package sample

import (
	"fmt"
	"github.com/jacobsa/oglemock"
	"reflect"
	"unsafe"
)

type mockWriter struct {
	controller oglemock.Controller
}

func (w *mockWriter) Oglemock_Id() uintptr {
	return uintptr(unsafe.Pointer(w))
}

func (w *mockWriter) Oglemock_Description() string {
	return "TODO"
}

func (w *mockWriter) Write(p []byte) (n int, err error) {
	retVals := w.controller.HandleMethodCall(w, "Write")
	if len(returnVals) != 2 {
		panic(fmt.Sprintf("mockWriter.Write: invalid values: %v", retVals))
	}

	var v reflect.Value

	// n int
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(n) {
		panic(fmt.Sprintf("mockWriter.Write: invalid return value 0: %v", v))
	}
	n = v.Int()

	// err error
	v = reflect.ValueOf(retVals[1])
	if v.Type() != reflect.TypeOf(err) {
		panic(fmt.Sprintf("mockWriter.Write: invalid return value 0: %v", v))
	}
	err = v.Interface.(error)

	return
}
