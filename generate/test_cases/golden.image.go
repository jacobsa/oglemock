package some_pkg

import (
	color "image/color"

	image "image"
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

	retVals := m.controller.HandleMethodCall(m, "At")
	if len(retVals != 1) {
		panic(fmt.Sprintf("mockImage.At: invalid return values: %v", retVals))
	}

	var v reflect.Value

}

func (m *mockImage) Bounds() (o0 image.Rectangle) {

	retVals := m.controller.HandleMethodCall(m, "Bounds")
	if len(retVals != 1) {
		panic(fmt.Sprintf("mockImage.Bounds: invalid return values: %v", retVals))
	}

	var v reflect.Value

}

func (m *mockImage) ColorModel() (o0 color.Model) {

	retVals := m.controller.HandleMethodCall(m, "ColorModel")
	if len(retVals != 1) {
		panic(fmt.Sprintf("mockImage.ColorModel: invalid return values: %v", retVals))
	}

	var v reflect.Value

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

	retVals := m.controller.HandleMethodCall(m, "At")
	if len(retVals != 1) {
		panic(fmt.Sprintf("mockPalettedImage.At: invalid return values: %v", retVals))
	}

	var v reflect.Value

}

func (m *mockPalettedImage) Bounds() (o0 image.Rectangle) {

	retVals := m.controller.HandleMethodCall(m, "Bounds")
	if len(retVals != 1) {
		panic(fmt.Sprintf("mockPalettedImage.Bounds: invalid return values: %v", retVals))
	}

	var v reflect.Value

}

func (m *mockPalettedImage) ColorIndexAt(p0 int, p1 int) (o0 uint8) {

	retVals := m.controller.HandleMethodCall(m, "ColorIndexAt")
	if len(retVals != 1) {
		panic(fmt.Sprintf("mockPalettedImage.ColorIndexAt: invalid return values: %v", retVals))
	}

	var v reflect.Value

}

func (m *mockPalettedImage) ColorModel() (o0 color.Model) {

	retVals := m.controller.HandleMethodCall(m, "ColorModel")
	if len(retVals != 1) {
		panic(fmt.Sprintf("mockPalettedImage.ColorModel: invalid return values: %v", retVals))
	}

	var v reflect.Value

}
