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

package main

import (
	. "github.com/jacobsa/ogletest"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"testing"
)

var dumpNew = flag.Bool("dump_new", false, "Dump new golden files.")

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

var createmockPath string

type CreateMockTest struct {
}

func TestOgletest(t *testing.T) { RunTests(t) }
func init() { RegisterTestSuite(&CreateMockTest{}) }

func (t *CreateMockTest) SetUpTestSuite() {
	// Create a temporary file to hold the built createmock binary.
	f, err := ioutil.TempFile("", "createmock-")
	if err != nil {
		panic("Creating temporary file: " + err.Error())
	}

	createmockPath = f.Name()
	f.Close()

	// Build the createmock tool so that it can be used in the tests below.
	cmd := exec.Command("go", "build", "-o", createmockPath, "github.com/jacobsa/oglemock/createmock")
	if output, err := cmd.CombinedOutput(); err != nil {
		panic(fmt.Sprintf("Error building createmock: %v\n\n%s", err, output))
	}
}

func (t *CreateMockTest) TearDownTestSuite() {
	// Delete the createmock binary we built above.
	os.Remove(createmockPath)
	createmockPath = ""
}

func (t *CreateMockTest) runGoldenTest(
	caseName string,
	expectedReturnCode int,
	createmockArgs ...string) {
  // Run createmock.
	cmd := exec.Command(createmockPath, createmockArgs...)
	output, err := cmd.CombinedOutput()

	// Make sure the process actually exited.
	exitError, ok := err.(*exec.ExitError)
	if err != nil && (!ok || !exitError.Exited()) {
		panic("exec.Command.CombinedOutput: " + err.Error())
	}

	// Extract a return code.
	var actualReturnCode int
	if exitError != nil {
		actualReturnCode = exitError.ExitStatus()
	}

	// Make sure the return code is correct.
	ExpectEq(expectedReturnCode, actualReturnCode)

	// Read the golden file.
	goldenPath := path.Join("test_cases", "golden." + caseName + ".go")
	goldenData := readFileOrDie(goldenPath)

	// Compare the two.
	identical := (string(output) == string(goldenData))
	ExpectTrue(identical, "Output doesn't match for case '%s'.", caseName)

	// Write out a new golden file if requested.
	if !identical && *dumpNew {
		writeContentsToFileOrDie(output, goldenPath)
	}
}

func writeContentsToFileOrDie(contents []byte, path string) {
	if err := ioutil.WriteFile(path, contents, 0600); err != nil {
		panic("ioutil.WriteFile: " + err.Error())
	}
}

func readFileOrDie(path string) []byte {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic("ioutil.ReadFile: " + err.Error())
	}

	return contents
}

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *CreateMockTest) IoPartial() {
	t.runGoldenTest(
		"io_partial",
		0,
		"io",
		"Reader",
		"Writer")
}
