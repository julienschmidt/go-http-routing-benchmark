// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package suite

import "github.com/julienschmidt/go-http-routing-benchmark/driver"

func init() {
	p1 := driver.NewBenchmark(
		driver.ParameterReadWrite,
		"OneParam",
		"Request one parameterized endpoint with two segments, one of which is a parameter",
	)
	p1.AddRoute("GET", "/user/:name")
	p1.AddRequest("GET", "/user/gordon")
	p1.AddTest(TestResponseOkWithoutOutput)
	p1.AddTest(TestParameterizedRoutesContainsNParams(1))

	p2 := driver.NewBenchmark(
		driver.ParameterReadWrite,
		"FiveParams",
		"Request one parameterized endpoint with five segments, all of which are parameters",
	)
	p2.AddRoute("GET", "/:a/:b/:c/:d/:e")
	p2.AddRequest("GET", "/test/test/test/test/test")
	p2.AddTest(TestResponseOkWithoutOutput)
	p2.AddTest(TestParameterizedRoutesContainsNParams(5))

	p3 := driver.NewBenchmark(
		driver.ParameterReadWrite,
		"TwentyParams",
		"Request one parameterized endpoint with twenty segments, all of which are parameters",
	)
	p3.AddTest(TestResponseOkWithoutOutput)
	p3.AddTest(TestParameterizedRoutesContainsNParams(20))
	p3.AddRequest("GET", "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t")
	p3.AddRoute("GET", "/:a/:b/:c/:d/:e/:f/:g/:h/:i/:j/:k/:l/:m/:n/:o/:p/:q/:r/:s/:t")
}

func init() {
	rwp1 := driver.NewBenchmark(
		driver.ParameterReadWrite,
		"ReadWriteOneParam",
		"Request one parameterized endpoint with two segments, one of which is a "+
			"parameter and respond with the value of the last parameter",
	)
	rwp1.AddTest(TestParameterReadWriteName)
	rwp1.AddTest(TestResponseOkParameterReadWrite)
	rwp1.AddTest(TestParameterizedRoutesContainsNParams(1))
	rwp1.AddRoute("GET", "/user/:"+driver.ParamNameReadWrite)
	rwp1.AddRequest("GET", "/user/"+driver.ParamNameReadWrite)

	rwp2 := driver.NewBenchmark(
		driver.ParameterReadWrite,
		"ReadWriteFiveParams",
		"Request one parameterized endpoint with five segments, all of which "+
			"are parameters and respond with the value of the last parameter",
	)
	rwp2.AddTest(TestParameterReadWriteName)
	rwp2.AddTest(TestResponseOkParameterReadWrite)
	rwp2.AddTest(TestParameterizedRoutesContainsNParams(5))
	rwp2.AddRoute("GET", "/:a/:b/:c/:d/:"+driver.ParamNameReadWrite)
	rwp2.AddRequest("GET", "/test/test/test/test/"+driver.ParamNameReadWrite)

	rwp3 := driver.NewBenchmark(
		driver.ParameterReadWrite,
		"ReadWriteTwentyParams",
		"Request one parameterized endpoint with twenty segments, all of which "+
			"are parameters and respond with the value of the last parameter",
	)
	rwp3.AddTest(TestParameterReadWriteName)
	rwp3.AddTest(TestResponseOkParameterReadWrite)
	rwp3.AddTest(TestParameterizedRoutesContainsNParams(20))
	rwp3.AddRequest("GET", "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/"+driver.ParamNameReadWrite)
	rwp3.AddRoute("GET", "/:a/:b/:c/:d/:e/:f/:g/:h/:i/:j/:k/:l/:m/:n/:o/:p/:q/:r/:s/:"+driver.ParamNameReadWrite)
}
