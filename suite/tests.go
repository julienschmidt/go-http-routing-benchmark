// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package suite

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

// TestStaticRoutesAreStatic ensures that static routes contain no parameterized segment.
var TestStaticRoutesAreStatic = &driver.Test{
	Description: "Test that the read/write parameter exists",
	Route: func(f *driver.Fixture) *driver.Assertion {
		x := strings.Count(f.Path, "/:")
		if x > 0 {
			return &driver.Assertion{
				Error: "Static routes must not contain parameters: " + f.Path,
			}
		}

		return nil
	},
}

// TestParameterizedRoutesContainsNParams ensures that parameterized routes
// contain n parameterized segments.
func TestParameterizedRoutesContainsNParams(n int) *driver.Test {
	return &driver.Test{
		Description: "Test that the read/write parameter exists",
		Route: func(f *driver.Fixture) *driver.Assertion {
			x := strings.Count(f.Path, "/:")
			if x != n {
				return &driver.Assertion{
					Error: fmt.Sprintf("Parameterized route must contain %d parameter(s), has %d", n, x),
				}
			}

			return nil
		},
	}
}

// TestParameterReadWriteName tests that the route contains the parameter
// name used to benchmark ParameterReadWrite.
var TestParameterReadWriteName = &driver.Test{
	Description: "Test that the read/write parameter exists",
	Route: func(f *driver.Fixture) *driver.Assertion {
		x := strings.Count(f.Path, "/:"+driver.ParamNameReadWrite)
		if x == 0 {
			return &driver.Assertion{
				Error: "The route contains no parameter named " + driver.ParamNameReadWrite,
			}
		} else if x > 1 {
			return &driver.Assertion{
				Error: "The route contains too many parameters named " + driver.ParamNameReadWrite,
			}
		}

		return nil
	},
}

// TestResponseOkParameterReadWrite is a test case that tests that the response status
// code is 200 and the response body is equal to the value of ParameterReadWrite name.
var TestResponseOkParameterReadWrite = &driver.Test{
	Description: "Test that the response status is 200 ok without output written",
	Request: func(r *http.Request, w *driver.Response) *driver.Assertion {
		return &driver.Assertion{
			Expect: fmt.Sprintf("Status code: 200, Output: \"%s\"", driver.ParamNameReadWrite),
			Actual: fmt.Sprintf("Status code: %d, Output: \"%s\"", w.Status(), string(w.Output())),
		}
	},
}

// TestResponseNotFound is a test case that tests that the response status
// code is 404 with or without output.
var TestResponseNotFound = &driver.Test{
	Description: "Test that the response status is 404 not found",
	Request: func(r *http.Request, w *driver.Response) *driver.Assertion {
		return &driver.Assertion{
			Expect: "Status code: 404",
			Actual: fmt.Sprintf("Status code: %d", w.Status()),
		}
	},
}

// TestResponseOkWithoutOutput is a test case that tests that the response status
// code is 0 and the response body is empty.
var TestResponseOkWithoutOutput = &driver.Test{
	Description: "Test that the response status is 200 without output written",
	Request: func(r *http.Request, w *driver.Response) *driver.Assertion {
		return &driver.Assertion{
			Expect: "Status code: 200, Output: \"\"",
			Actual: fmt.Sprintf("Status code: %d, Output: \"%s\"", w.Status(), string(w.Output())),
		}
	},
}
