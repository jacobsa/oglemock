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

// Return creates an Action that returns the values passed to Return as
// arguments, after suitable legal type conversions. The following rules apply.
// Given an argument x to Return and a corresponding type T in the method's
// signature, at least one of the following must hold:
//
//  *  x is assignable to T. (See "Assignability" in the language spec.)
//
//  *  x is the nil literal and T is a pointer, function, interface, slice,
//     channel, or map type.
//
//  *  T is any numeric type, and x is an int that is in-range for that type.
//     This facilities using raw integer constants: Return(17).
//
//  *  T is a floating-point or complex number type, and x is a float64.  This
//     facilities using raw floating-point constants: Return(17.5).
//
//  *  T is a complex number type, and x is a complex128. This facilities using
//     raw complex constants: Return(17+2i).
//
func Return(vals ...interface{}) Action {
	return &returnAction{vals}
}

type returnAction struct {
	returnVals []interface{}
}

func (a *returnAction) Invoke(vals []interface{}) []interface{} {
	return a.returnVals
}

func (a *returnAction) SetSignature(signature reflect.Type) error {
	// Check the length of the return value.
	numOut := signature.NumOut()
	numVals := len(a.returnVals)

	if numOut != numVals {
		return errors.New(
			fmt.Sprintf("Return given %d vals; expected %d.", numVals, numOut))
	}

	// Check the type of each value to be returned.
	for i, val := range a.returnVals {
		expectedType := signature.Out(i)
		actualType := reflect.TypeOf(val)

		if expectedType != actualType {
			return errors.New(
				fmt.Sprintf(
					"Return arg %d: given %v; expected %v.",
					i,
					actualType,
					expectedType))
		}
	}

	return nil
}
