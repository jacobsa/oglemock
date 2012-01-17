package some_pkg

import ()

type mockReader struct {
	controller	oglemock.Controller
	description	string
}

func NewMockReader(c oglemock.Controller,
	desc string) *mockReader {
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

	retVals := m.controller.HandleMethodCall(m, "Read")
	if len(retVals != 2) {
		panic(fmt.Sprintf("mockReader.Read: invalid return values: %v", retVals))
	}

	var v reflect.Value

}

type mockWriter struct {
	controller	oglemock.Controller
	description	string
}

func NewMockWriter(c oglemock.Controller,
	desc string) *mockWriter {
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

	retVals := m.controller.HandleMethodCall(m, "Write")
	if len(retVals != 2) {
		panic(fmt.Sprintf("mockWriter.Write: invalid return values: %v", retVals))
	}

	var v reflect.Value

}
