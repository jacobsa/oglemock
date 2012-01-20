// This file was auto-generated using createmock. See the following page for
// more information:
//
//     https://github.com/jacobsa/oglemock
//

package some_pkg

import (
	fmt "fmt"
	io "io"
	oglemock "github.com/jacobsa/oglemock"
	reflect "reflect"
	runtime "runtime"
	unsafe "unsafe"
)

type MockReader interface {
	io.Reader
	oglemock.MockObject
}

type mockReader struct {
	controller	oglemock.Controller
	description	string
}

func NewMockReader(c oglemock.Controller,
	desc string) MockReader {
	return &mockReader{
		controller:	c,
		description:	desc,
	}
}

func (m *mockReader) Oglemock_Id() uintptr {
	return uintptr(unsafe.Pointer(m))
}

func (m *mockReader) Oglemock_Description() string {
	return m.description
}

func (m *mockReader) Read(p0 []uint8) (o0 int, o1 error) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"Read",
		file,
		line,
		[]interface{}{p0})

	if len(retVals) != 2 {
		panic(fmt.Sprintf("mockReader.Read: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 int
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockReader.int: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(int)

	// o1 error
	v = reflect.ValueOf(retVals[1])
	if v.Type() != reflect.TypeOf(o1) {
		panic(fmt.Sprintf("mockReader.error: invalid return value 1: %v", v))
	}
	o1 = v.Interface().(error)

	return
}

type MockWriter interface {
	io.Writer
	oglemock.MockObject
}

type mockWriter struct {
	controller	oglemock.Controller
	description	string
}

func NewMockWriter(c oglemock.Controller,
	desc string) MockWriter {
	return &mockWriter{
		controller:	c,
		description:	desc,
	}
}

func (m *mockWriter) Oglemock_Id() uintptr {
	return uintptr(unsafe.Pointer(m))
}

func (m *mockWriter) Oglemock_Description() string {
	return m.description
}

func (m *mockWriter) Write(p0 []uint8) (o0 int, o1 error) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"Write",
		file,
		line,
		[]interface{}{p0})

	if len(retVals) != 2 {
		panic(fmt.Sprintf("mockWriter.Write: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 int
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockWriter.int: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(int)

	// o1 error
	v = reflect.ValueOf(retVals[1])
	if v.Type() != reflect.TypeOf(o1) {
		panic(fmt.Sprintf("mockWriter.error: invalid return value 1: %v", v))
	}
	o1 = v.Interface().(error)

	return
}
