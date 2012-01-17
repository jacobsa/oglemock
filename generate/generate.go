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
	"errors"
	"io"
	"path"
	"reflect"
	"text/template"
)

const tmplStr =
`package {{.Pkg}}

import ({{range $identifier, $import := .Imports}}
	{{$identifier}} "{{$import}}"{{end}}
)

{{range .Interfaces}}
type mock{{.Name}} struct {
	controller oglemock.Controller
}
{{end}}
`

var tmpl = template.Must(template.New("code").Parse(tmplStr))

// A map from import identifier to package to use that identifier for,
// containing elements for each import needed by a set of mocked interfaces.
type importMap map[string]string

type tmplArg struct {
	// The package of the generated code.
	Pkg string

	// Imports needed by the interfaces.
	Imports importMap

	// The set of interfaces to mock.
	Interfaces []reflect.Type
}

// Add an import for the supplied type, without recursing.
func addImportForType(imports importMap, t reflect.Type) {
	// If there is no package path, this is a built-in type and we don't need an
	// import.
	pkgPath := t.PkgPath()
	if pkgPath == "" {
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

	// Create an appropriate template arg, then execute the template.
	var arg tmplArg
	arg.Pkg = pkg
	arg.Imports = getImports(interfaces)
	arg.Interfaces = interfaces

	return tmpl.Execute(w, arg)
}
