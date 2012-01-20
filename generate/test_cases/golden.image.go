// This file was auto-generated using createmock. See the following page for
// more information:
//
//     https://github.com/jacobsa/oglemock
//

package some_pkg

import (
	color "image/color"
	fmt "fmt"
	image "image"
	oglemock "github.com/jacobsa/oglemock"
	reflect "reflect"
	runtime "runtime"
	unsafe "unsafe"
)

type mockImage struct {
	controller	oglemock.Controller
	description	string
}

func NewMockImage(c oglemock.Controller,
	desc string) *mockImage {
	return &mockImage{
		controller:	c,
		description:	desc,
	}
}

func (m *mockImage) Oglemock_Id() uintptr {
	return uintptr(unsafe.Pointer(m))
}

func (m *mockImage) Oglemock_Description() string {
	return m.description
}

func (m *mockImage) At(p0 int, p1 int) (o0 color.Color) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"At",
		file,
		line,
		[]interface{}{p0, p1})

	if len(retVals) != 1 {
		panic(fmt.Sprintf("mockImage.At: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 color.Color
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockImage.Color: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(color.Color)

	return
}

func (m *mockImage) Bounds() (o0 image.Rectangle) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"Bounds",
		file,
		line,
		[]interface{}{})

	if len(retVals) != 1 {
		panic(fmt.Sprintf("mockImage.Bounds: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 image.Rectangle
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockImage.Rectangle: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(image.Rectangle)

	return
}

func (m *mockImage) ColorModel() (o0 color.Model) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"ColorModel",
		file,
		line,
		[]interface{}{})

	if len(retVals) != 1 {
		panic(fmt.Sprintf("mockImage.ColorModel: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 color.Model
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockImage.Model: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(color.Model)

	return
}

type mockPalettedImage struct {
	controller	oglemock.Controller
	description	string
}

func NewMockPalettedImage(c oglemock.Controller,
	desc string) *mockPalettedImage {
	return &mockPalettedImage{
		controller:	c,
		description:	desc,
	}
}

func (m *mockPalettedImage) Oglemock_Id() uintptr {
	return uintptr(unsafe.Pointer(m))
}

func (m *mockPalettedImage) Oglemock_Description() string {
	return m.description
}

func (m *mockPalettedImage) At(p0 int, p1 int) (o0 color.Color) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"At",
		file,
		line,
		[]interface{}{p0, p1})

	if len(retVals) != 1 {
		panic(fmt.Sprintf("mockPalettedImage.At: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 color.Color
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockPalettedImage.Color: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(color.Color)

	return
}

func (m *mockPalettedImage) Bounds() (o0 image.Rectangle) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"Bounds",
		file,
		line,
		[]interface{}{})

	if len(retVals) != 1 {
		panic(fmt.Sprintf("mockPalettedImage.Bounds: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 image.Rectangle
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockPalettedImage.Rectangle: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(image.Rectangle)

	return
}

func (m *mockPalettedImage) ColorIndexAt(p0 int, p1 int) (o0 uint8) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"ColorIndexAt",
		file,
		line,
		[]interface{}{p0, p1})

	if len(retVals) != 1 {
		panic(fmt.Sprintf("mockPalettedImage.ColorIndexAt: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 uint8
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockPalettedImage.uint8: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(uint8)

	return
}

func (m *mockPalettedImage) ColorModel() (o0 color.Model) {
	// Get a file name and line number for the caller.
	_, file, line, _ := runtime.Caller(1)

	// Hand the call off to the controller, which does most of the work.
	retVals := m.controller.HandleMethodCall(
		m,
		"ColorModel",
		file,
		line,
		[]interface{}{})

	if len(retVals) != 1 {
		panic(fmt.Sprintf("mockPalettedImage.ColorModel: invalid return values: %v", retVals))
	}

	var v reflect.Value

	// o0 color.Model
	v = reflect.ValueOf(retVals[0])
	if v.Type() != reflect.TypeOf(o0) {
		panic(fmt.Sprintf("mockPalettedImage.Model: invalid return value 0: %v", v))
	}
	o0 = v.Interface().(color.Model)

	return
}
