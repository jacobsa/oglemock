package some_pkg

import (
	complicated_pkg "github.com/jacobsa/oglemock/generate/test_cases/complicated_pkg"

	fmt "fmt"

	image "image"

	io "io"

	net "net"

	oglemock "github.com/jacobsa/oglemock"

	reflect "reflect"

	renamed_pkg "github.com/jacobsa/oglemock/generate/test_cases/renamed_pkg"

	runtime "runtime"

	unsafe "unsafe"
)

type mockComplicatedThing struct {
	controller	oglemock.Controller
	description	string
}

func NewMockComplicatedThing(c oglemock.Controller,
	desc string) *mockComplicatedThing {
	return &mockComplicatedThing{
		controller:	c,
		description:	desc,
	}
}

func (m *mockComplicatedThing) Oglemock_Id() uintptr {
	return uintptr(unsafe.Pointer(m))
}

func (m *mockComplicatedThing) Oglemock_Description() string {
	return m.description
}

func (m *mockComplicatedThing) Arrays(p0 [3]string) (o0 [3]int, o1 error) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"Arrays",
		file,
		line,
		[]interface{}{p0})

	if len(retVals) != 2 {
		panic(fmt.Sprintf("mockComplicatedThing.Arrays: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 [3]int
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockComplicatedThing.: invalid return value 0: %v", v))
	}
	o0 = v.Interface().([3]int)

	// o1 error
	v = reflect.ValueOf(retVals[1])
	if v.Type() != reflect.TypeOf(o1) {
		panic(fmt.Sprintf("mockComplicatedThing.error: invalid return value 1: %v", v))
	}
	o1 = v.Interface().(error)

	return
}

func (m *mockComplicatedThing) Channels(p0 chan chan<- <-chan net.Conn) (o0 chan int) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"Channels",
		file,
		line,
		[]interface{}{p0})

	if len(retVals) != 1 {
		panic(fmt.Sprintf("mockComplicatedThing.Channels: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 chan int
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockComplicatedThing.: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(chan int)

	return
}

func (m *mockComplicatedThing) EmptyInterface(p0 interface{}) (o0 interface{}, o1 error) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"EmptyInterface",
		file,
		line,
		[]interface{}{p0})

	if len(retVals) != 2 {
		panic(fmt.Sprintf("mockComplicatedThing.EmptyInterface: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 interface {}
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockComplicatedThing.: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(interface{})

	// o1 error
	v = reflect.ValueOf(retVals[1])
	if v.Type() != reflect.TypeOf(o1) {
		panic(fmt.Sprintf("mockComplicatedThing.error: invalid return value 1: %v", v))
	}
	o1 = v.Interface().(error)

	return
}

func (m *mockComplicatedThing) Functions(p0 func(int, image.Image) int) (o0 func(string, int) net.Conn) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"Functions",
		file,
		line,
		[]interface{}{p0})

	if len(retVals) != 1 {
		panic(fmt.Sprintf("mockComplicatedThing.Functions: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 func(string, int) net.Conn
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockComplicatedThing.: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(func(string, int) net.Conn)

	return
}

func (m *mockComplicatedThing) Maps(p0 map[string]*int) (o0 map[int]*string, o1 error) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"Maps",
		file,
		line,
		[]interface{}{p0})

	if len(retVals) != 2 {
		panic(fmt.Sprintf("mockComplicatedThing.Maps: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 map[int]*string
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockComplicatedThing.: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(map[int]*string)

	// o1 error
	v = reflect.ValueOf(retVals[1])
	if v.Type() != reflect.TypeOf(o1) {
		panic(fmt.Sprintf("mockComplicatedThing.error: invalid return value 1: %v", v))
	}
	o1 = v.Interface().(error)

	return
}

func (m *mockComplicatedThing) NamedScalarType(p0 complicated_pkg.Byte) (o0 []complicated_pkg.Byte, o1 error) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"NamedScalarType",
		file,
		line,
		[]interface{}{p0})

	if len(retVals) != 2 {
		panic(fmt.Sprintf("mockComplicatedThing.NamedScalarType: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 []complicated_pkg.Byte
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockComplicatedThing.: invalid return value 0: %v", v))
	}
	o0 = v.Interface().([]complicated_pkg.Byte)

	// o1 error
	v = reflect.ValueOf(retVals[1])
	if v.Type() != reflect.TypeOf(o1) {
		panic(fmt.Sprintf("mockComplicatedThing.error: invalid return value 1: %v", v))
	}
	o1 = v.Interface().(error)

	return
}

func (m *mockComplicatedThing) Pointers(p0 *int, p1 *net.Conn, p2 **io.Reader) (o0 *int, o1 error) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"Pointers",
		file,
		line,
		[]interface{}{p0, p1, p2})

	if len(retVals) != 2 {
		panic(fmt.Sprintf("mockComplicatedThing.Pointers: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 *int
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockComplicatedThing.: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(*int)

	// o1 error
	v = reflect.ValueOf(retVals[1])
	if v.Type() != reflect.TypeOf(o1) {
		panic(fmt.Sprintf("mockComplicatedThing.error: invalid return value 1: %v", v))
	}
	o1 = v.Interface().(error)

	return
}

func (m *mockComplicatedThing) RenamedPackage(p0 tony.SomeUint8Alias) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"RenamedPackage",
		file,
		line,
		[]interface{}{p0})

	if len(retVals) != 0 {
		panic(fmt.Sprintf("mockComplicatedThing.RenamedPackage: invalid return values: %v", retVals))
	}

	var v reflect.Value

	return
}

func (m *mockComplicatedThing) Slices(p0 []string) (o0 []int, o1 error) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"Slices",
		file,
		line,
		[]interface{}{p0})

	if len(retVals) != 2 {
		panic(fmt.Sprintf("mockComplicatedThing.Slices: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 []int
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockComplicatedThing.: invalid return value 0: %v", v))
	}
	o0 = v.Interface().([]int)

	// o1 error
	v = reflect.ValueOf(retVals[1])
	if v.Type() != reflect.TypeOf(o1) {
		panic(fmt.Sprintf("mockComplicatedThing.error: invalid return value 1: %v", v))
	}
	o1 = v.Interface().(error)

	return
}

func (m *mockComplicatedThing) Variadic(p0 int, p1 []net.Conn) (o0 int) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"Variadic",
		file,
		line,
		[]interface{}{p0, p1})

	if len(retVals) != 1 {
		panic(fmt.Sprintf("mockComplicatedThing.Variadic: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 int
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockComplicatedThing.int: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(int)

	return
}
