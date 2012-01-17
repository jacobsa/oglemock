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

func (m *mockImage) At() {
}

func (m *mockImage) Bounds() {
}

func (m *mockImage) ColorModel() {
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

func (m *mockPalettedImage) At() {
}

func (m *mockPalettedImage) Bounds() {
}

func (m *mockPalettedImage) ColorIndexAt() {
}

func (m *mockPalettedImage) ColorModel() {
}
