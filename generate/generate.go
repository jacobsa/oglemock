// Copyright 2012 Aaron Jacobs. All Rights Reserved.
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

// Package generate implements code generation for mock classes. This is an
// implementation detail of the createmock command, which you probably want to
// use directly instead.
package generate

import (
	"bytes"
	"errors"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"path"
	"reflect"
	"text/template"
)

const tmplStr =
`package {{.Pkg}}

import (
	{{range $identifier, $import := .Imports}}
		{{$identifier}} "{{$import}}"
	{{end}}
)

{{range .Interfaces}}
	{{$structName := printf "mock%s" .Name}}

	type {{$structName}} struct {
		controller oglemock.Controller
		description string
	}
	
	func New{{printf "Mock%s" .Name}}(
		c oglemock.Controller,
		desc string) *{{$structName}} {
	  return &{{$structName}}{
			controller: c,
			description: desc,
		}
	}
	
	func (m *{{$structName}}) Oglemock_Id() uintptr {
		return uintptr(unsafe.Pointer(m))
	}
	
	func (m *{{$structName}}) Oglemock_Description() string {
		return m.description
	}

	{{range getMethods .}}
	  {{$inputTypes := getInputs .Type}}
	  {{$outputTypes := getOutputs .Type}}

		func (m *{{$structName}}) {{.Name}}({{range $i, $type := $inputTypes}}p{{$i}} {{getTypeString $type}}, {{end}}
		) ({{range $i, $type := $outputTypes}}o{{$i}} {{getTypeString $type}}, {{end}}
		){
			// Hand the call off to the controller, which does most of the work.
			retVals := m.controller.HandleMethodCall(m, "{{.Name}}")
			if len(retVals != {{len $outputTypes}}) {
				panic(fmt.Sprintf("{{$structName}}.{{.Name}}: invalid return values: %v", retVals))
			}

			var v reflect.Value
			{{range $i, $type := $outputTypes}}
				// o{{$i}} {{getTypeString $type}}
				v = reflect.ValueOf(retVals[{{$i}}])
				if v.Type() != reflect.TypeOf(o{{$i}}) {
					panic(fmt.Sprintf("{{$structName}}.{{.Name}}: invalid return value {{$i}}: %v", v))
				}
				o{{$i}} = v.Interface().({{getTypeString $type}})
			{{end}}

			return
		}
	{{end}}
{{end}}
`

type tmplArg struct {
	// The package of the generated code.
	Pkg string

	// Imports needed by the interfaces.
	Imports importMap

	// The set of interfaces to mock.
	Interfaces []reflect.Type
}

var tmpl *template.Template

func init() {
	extraFuncs := make(template.FuncMap)
	extraFuncs["getMethods"] = getMethods
	extraFuncs["getInputs"] = getInputs
	extraFuncs["getOutputs"] = getOutputs
	extraFuncs["getTypeString"] = getTypeString

	tmpl = template.New("code")
	tmpl.Funcs(extraFuncs)
	tmpl.Parse(tmplStr)
}

func getTypeString(t reflect.Type) string {
	return t.String()
}

func getMethods(it reflect.Type) []reflect.Method {
	numMethods := it.NumMethod()
	methods := make([]reflect.Method, numMethods)

	for i := 0; i < numMethods; i++ {
		methods[i] = it.Method(i)
	}

	return methods
}

func getInputs(ft reflect.Type) []reflect.Type {
	numIn := ft.NumIn()
	inputs := make([]reflect.Type, numIn)

	for i := 0; i < numIn; i++ {
		inputs[i] = ft.In(i)
	}

	return inputs
}

func getOutputs(ft reflect.Type) []reflect.Type {
	numOut := ft.NumOut()
	outputs := make([]reflect.Type, numOut)

	for i := 0; i < numOut; i++ {
		outputs[i] = ft.Out(i)
	}

	return outputs
}

// A map from import identifier to package to use that identifier for,
// containing elements for each import needed by a set of mocked interfaces.
type importMap map[string]string

// Add an import for the supplied type, without recursing.
func addImportForType(imports importMap, t reflect.Type) {
	// If there is no package path, this is a built-in type and we don't need an
	// import.
	pkgPath := t.PkgPath()
	if pkgPath == "" {
		return
	}

	// Work around a bug in Go:
	//
	//     http://code.google.com/p/go/issues/detail?id=2660
	//
	var errorPtr *error
	if t == reflect.TypeOf(errorPtr).Elem() {
		return
	}

	// Use the package's base name as an identifier.
	imports[path.Base(pkgPath)] = pkgPath
}

// Add all necessary imports for the type, recursing as appropriate.
func addImportsForType(imports importMap, t reflect.Type) {
	// Add any import needed for the type itself.
	addImportForType(imports, t)

	// Handle special cases where recursion is needed.
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Ptr, reflect.Slice:
		addImportsForType(imports, t.Elem())

	case reflect.Func:
		// Input parameters.
		for i := 0; i < t.NumIn(); i++ {
			addImportsForType(imports, t.In(i))
		}

		// Return values.
		for i := 0; i < t.NumOut(); i++ {
			addImportsForType(imports, t.Out(i))
		}

	case reflect.Map:
		addImportsForType(imports, t.Key())
		addImportsForType(imports, t.Elem())
	}
}

// Add imports for each of the methods of the interface, but not the interface
// itself.
func addImportsForInterfaceMethods(imports importMap, it reflect.Type) {
	// Handle each method.
	for i := 0; i < it.NumMethod(); i++ {
		m := it.Method(i)
		addImportsForType(imports, m.Type)
	}
}

// Given a set of interfaces, return a map from import identifier to package to
// use that identifier for, containing elements for each import needed by the
// mock versions of those interfaces.
func getImports(interfaces []reflect.Type) importMap {
	imports := make(importMap)
	for _, it := range interfaces {
		addImportsForInterfaceMethods(imports, it)
	}

	return imports
}

// Given a set of interfaces to mock, write out source code for a package named
// `pkg` that contains mock implementations of those interfaces.
func GenerateMockSource(w io.Writer, pkg string, interfaces []reflect.Type) error {
	// Sanity-check arguments.
	if pkg == "" {
		return errors.New("Package name must be non-empty.")
	}

	if len(interfaces) == 0 {
		return errors.New("List of interfaces must be non-empty.")
	}

	// Make sure each type is indeed an interface.
	for _, it := range interfaces {
		if it.Kind() != reflect.Interface {
			return errors.New("Invalid type: " + it.String())
		}
	}

	// Create an appropriate template arg, then execute the template. Write the
	// raw output into a buffer.
	var arg tmplArg
	arg.Pkg = pkg
	arg.Imports = getImports(interfaces)
	arg.Interfaces = interfaces

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, arg); err != nil {
		return err
	}

	// Pretty-print the output.
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, pkg + ".go", buf, parser.ParseComments)
	if err != nil {
		return errors.New("Error parsing generated code: " + err.Error())
	}

	if err = printer.Fprint(w, fset, astFile); err != nil {
		return errors.New("Error pretty printing: " + err.Error())
	}

	return nil
}
