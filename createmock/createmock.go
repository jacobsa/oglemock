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
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"text/template"

	// Ensure that the generate package, which is used by the generated code, is
	// installed by goinstall.
	_ "github.com/jacobsa/oglemock/generate"
)

// A template for generated code that is used to print the result.
const tmplStr =
`
{{$inputPkg := .InputPkg}}
{{$outputPkg := .OutputPkg}}

package main

import (
	"{{$inputPkg}}"
	"github.com/jacobsa/oglemock/generate"
	"os"
	"reflect"
)

func getTypeForPtr(ptr interface{}) reflect.Type {
	return reflect.TypeOf(ptr).Elem()
}

func main() {
	// Reduce noise in logging output.
	log.SetFlags(0)

	interfaces := []reflect.Type{
		{{range $typeName := .TypeNames}}
			getTypeForPtr((*{{$inputPkg}}.{{$typeName}})(nil)),
		{{end}}
	}

	err := generate.GenerateMockSource(os.Stdout, "{{$outputPkg}}", interfaces)
	if err != nil {
		log.Fatalf("Error generating mock source: %v", err)
	}
}
`

type tmplArg struct {
	InputPkg string
	OutputPkg string

	// Types to be mocked, relative to their package's name.
	TypeNames []string
}

func main() {
	// Reduce noise in logging output.
	log.SetFlags(0)

	// Check the command-line arguments.
	flag.Parse()

	cmdLineArgs := flag.Args()
	if len(cmdLineArgs) < 2 {
		fmt.Println("Usage: createmock [package] [interface ...]")
		os.Exit(1)
	}

	// Create a temporary file to hold generated code.
	codeFile, err := ioutil.TempFile("", "createmock-")
	if err != nil {
		log.Fatalf("Couldn't create a temporary file: %v", err)
	}

	codePath := codeFile.Name()
	defer os.Remove(codePath)

	// Create another temporary file to hold a compiled binary.
	binaryFile, err := ioutil.TempFile("", "createmock-")
	if err != nil {
		log.Fatalf("Couldn't create a temporary file: %v", err)
	}

	binaryPath := binaryFile.Name()
	binaryFile.Close()
	defer os.Remove(binaryPath)

	// Create an appropriate template argument.
	var arg tmplArg
	arg.InputPkg = cmdLineArgs[0]
	arg.OutputPkg = "mock_" + path.Base(arg.InputPkg)
	arg.TypeNames = cmdLineArgs[1:]

	// Execute the template to generate code that will itself generate the mock
	// code. Write the code to the temp file.
	tmpl := template.Must(template.New("code").Parse(tmplStr))
	if err := tmpl.Execute(codeFile, arg); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	codeFile.Close()

	// Attempt to build the code.
	cmd := exec.Command("go", "build", "-o", binaryPath, codePath)
	buildOutput, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf(
			"%s\n\nError building generated code:\n\n" +
				"    %v\n\n. Please report this oglemock bug.",
			buildOutput,
		err)
	}
}
