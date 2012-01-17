package some_pkg

import (
	runtime "runtime"
)



type mockReader struct {
	controller oglemock.Controller
	description string
}

func NewMockReader(
	c oglemock.Controller,
	desc string) *mockReader {
  return &mockReader{
		controller: c,
		description: desc,
	}
}

func (m *mockReader) Oglemock_Id() uintptr {
	return uintptr(unsafe.Pointer(m))
}

func (m *mockReader) Oglemock_Description() string {
	return m.description
}


type mockWriter struct {
	controller oglemock.Controller
	description string
}

func NewMockWriter(
	c oglemock.Controller,
	desc string) *mockWriter {
  return &mockWriter{
		controller: c,
		description: desc,
	}
}

func (m *mockWriter) Oglemock_Id() uintptr {
	return uintptr(unsafe.Pointer(m))
}

func (m *mockWriter) Oglemock_Description() string {
	return m.description
}

