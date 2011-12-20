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

package oglemock

import (
	"errors"
	"fmt"
	"reflect"
)

// Return creates an Action that simply returns the values passed to Return as
// arguments. You should make sure that these values are correctly typed to be
// return values for the mock method.
func Return(vals ...interface{}) Action {
	return &returnAction{vals}
}

type returnAction struct {
	returnVals []interface{}
}

func (a *returnAction) Invoke(vals []interface{}) []interface{} {
	return a.returnVals
}

func (a *returnAction) CheckType(signature reflect.Type) error {
	// Check the length of the return value.
	numOut := signature.NumOut()
	numVals := len(a.returnVals)

	if numOut != numVals {
		return errors.New(
			fmt.Sprintf("Return given %d vals; expected %d.", numVals, numOut))
	}

	// Check the type of each.
	for i, val := range a.returnVals {
		expectedType := signature.Out(i)
		actualType := reflect.TypeOf(val)

		if expectedType != actualType {
		return errors.New(
			fmt.Sprintf(
				"Return given %v for arg %d; expected %v.",
				actualType,
				i,
				expectedType))
		}
	}

	return nil
}
