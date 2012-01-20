// This file was auto-generated using createmock. See the following page for
// more information:
//
//     https://github.com/jacobsa/oglemock
//

package some_pkg

import (
	fmt "fmt"
	oglemock "github.com/jacobsa/oglemock"
	reflect "reflect"
	runtime "runtime"
	tony "github.com/jacobsa/oglemock/generate/test_cases/renamed_pkg"
	unsafe "unsafe"
)

type MockSomeInterface interface {
	tony.SomeInterface
	oglemock.MockObject
}

type mockSomeInterface struct {
	controller	oglemock.Controller
	description	string
}

func NewMockSomeInterface(c oglemock.Controller,
	desc string) MockSomeInterface {
	return &mockSomeInterface{
		controller:	c,
		description:	desc,
	}
}

func (m *mockSomeInterface) Oglemock_Id() uintptr {
	return uintptr(unsafe.Pointer(m))
}

func (m *mockSomeInterface) Oglemock_Description() string {
	return m.description
}

func (m *mockSomeInterface) DoFoo(p0 int) (o0 int) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"DoFoo",
		file,
		line,
		[]interface{}{p0})

	if len(retVals) != 1 {
		panic(fmt.Sprintf("mockSomeInterface.DoFoo: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 int
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockSomeInterface.int: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(int)

	return
}
