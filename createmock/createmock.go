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

// createmock is used to generate source code for mock versions of interfaces
// from installed packages.
package main

import (
	"flag"
	"fmt"
	"os"
)

// A template for generated code that is used to print the result.
const tmplStr =
`
package main

import (
	"{{.InputPkg}}"
	"github.com/jacobsa/oglemock/generate"
	"os"
	"reflect"
)

func getTypeForPtr(ptr interface{}) reflect.Type {
	return reflect.TypeOf(ptr).Elem()
}

func main() {
	interfaces := []reflect.Type{
		{{range $typeName := .TypeNames}}
			(*{{$typeName}})(nil),
		{{end}}
	}

	err := generate.GenerateMockSource(os.Stdout, "{{.OutputPkg}}", interfaces)
	if err != nil {
		log.Errorf("Error generating mock source: %v", err)
		os.Exit(1)
	}
}
`

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: createmock [package] [interface ...]")
		os.Exit(1)
	}

	fmt.Println("TODO: Implement me.")
}
