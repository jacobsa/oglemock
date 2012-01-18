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
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
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
	"log"
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

var unknownPackageRegexp =
	regexp.MustCompile(`(?s): import ([\pL_0-9/]+)\s+package could not be found`)

var undefinedInterfaceRegexp =
	regexp.MustCompile(`tool\.go:\d+: undefined: [\pL_0-9]+\.([\pL_0-9]+)`)

// Does the 'go build' output indicate that a package wasn't found? If so,
// return the name of the package.
func findUnknownPackage(output []byte) *string {
	if match := unknownPackageRegexp.FindSubmatch(output); match != nil {
		res := string(match[1])
		return &res
	}

	return nil
}

// Does the 'go build' output indicate that an interface wasn't found? If so,
// return the name of the interface.
func findUndefinedInterface(output []byte) *string {
	if match := undefinedInterfaceRegexp.FindSubmatch(output); match != nil {
		res := string(match[1])
		return &res
	}

	return nil
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

	// Create a temporary directory inside of $GOPATH to hold generated code.
	tree, _, err := build.FindTree("github.com/jacobsa/oglemock")
	if err != nil {
		log.Fatalf("Couldn't find oglemock in $GOPATH: %v", err)
	}

	tmpDir, err := ioutil.TempDir(path.Join(tree.Path, "src"), "tmp-createmock-")
	if err != nil {
		log.Fatalf("Couldn't create a temporary directory: %v", err)
	}

	defer os.RemoveAll(tmpDir)

	// Create a file to hold generated code.
	codeFile, err := os.Create(path.Join(tmpDir, "tool.go"))
	if err != nil {
		log.Fatalf("Couldn't create a file to hold code: %v", err)
	}

	// Create an appropriate path for the built binary.
	binaryPath := path.Join(tmpDir, "tool")

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
	cmd := exec.Command("go", "build", "-o", binaryPath)
	cmd.Dir = tmpDir
	buildOutput, err := cmd.CombinedOutput()

	if err != nil {
		// Did the compilation fail due to the user-specified package not being found?
		if pkg := findUnknownPackage(buildOutput); pkg != nil && *pkg == arg.InputPkg {
			log.Fatalf("Unknown package: %s", *pkg)
		}

		// Did the compilation fail due to an unknown interface?
		if in := findUndefinedInterface(buildOutput); in != nil {
			log.Fatalf("Unknown interface: %s", *in)
		}

		// Otherwise print a generic error.
		log.Fatalf(
			"%s\n\nError building generated code:\n\n" +
				"    %v\n\n Please report this oglemock bug.",
			buildOutput,
		err)
	}

	// Run the binary.
	cmd = exec.Command(binaryPath)
	binaryOutput, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf(
			"%s\n\nError running generated code:\n\n" +
				"    %v\n\n Please report this oglemock bug.",
			binaryOutput,
		err)
	}

	// Copy its output.
	_, err = os.Stdout.Write(binaryOutput)
	if err != nil {
		log.Fatalf("Error copying binary output: %v", err)
	}
}
