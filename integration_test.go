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

package oglemock_test

import (
	. "github.com/jacobsa/ogletest"
	"github.com/jacobsa/oglemock"
	"github.com/jacobsa/oglemock/sample/mock_io"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type IntegrationTest struct {
	reporter fakeErrorReporter
	controller oglemock.Controller

	reader mock_io.MockReader
}

func init() { RegisterTestSuite(&IntegrationTest{}) }

func (t *IntegrationTest) SetUp(c *TestInfo) {
	t.reporter.errorsReported = make([]errorReport, 0)
	t.reporter.fatalErrorsReported = make([]errorReport, 0)
	t.controller = oglemock.NewController(&t.reporter)

	t.reader = mock_io.NewMockReader(t.controller, "")
}

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *IntegrationTest) ZeroValuesForScalars() {
	// Make an unexpected call.
	n, err := t.reader.Read([]uint8{})

	// Check the return values.
	ExpectEq(0, n)
	ExpectEq(nil, err)
}
