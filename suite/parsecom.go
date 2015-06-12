// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package suite

import "github.com/julienschmidt/go-http-routing-benchmark/driver"

func init() {
	pc1 := driver.NewBenchmark(
		driver.Parameterized,
		"ParseComStatic",
		"Request one static parse.com endpoint with two segments",
	)
	pc1.AddRoutes(parsecomRoutes)
	pc1.AddTest(TestResponseOkWithoutOutput)
	pc1.AddRequest("GET", "/1/users")

	pc2 := driver.NewBenchmark(
		driver.Parameterized,
		"ParseComParam",
		"Request one parameterized parse.com endpoint with three segments, one of which is a parameter",
	)
	pc2.AddRoutes(parsecomRoutes)
	pc2.AddTest(TestResponseOkWithoutOutput)
	pc2.AddRequest("GET", "/1/classes/go")

	pc3 := driver.NewBenchmark(
		driver.Parameterized,
		"ParseCom2Params",
		"Request one parameterized parse.com endpoint with four segments, two of which are parameters",
	)
	pc3.AddRoutes(parsecomRoutes)
	pc3.AddTest(TestResponseOkWithoutOutput)
	pc3.AddRequest("GET", "/1/classes/go/123456789")

	pc4 := driver.NewBenchmark(
		driver.Parameterized,
		"ParseComNotFound",
		"Request one unavailable parse.com endpoint",
	)
	pc4.AddRoutes(parsecomRoutes)
	pc4.AddTest(TestResponseNotFound)
	pc4.AddRequest("GET", "/notfound/this-path-is-unavailable") // Without trailing slash!

	pc5 := driver.NewBenchmark(
		driver.Parameterized,
		"ParseComAll",
		"Request all parse.com endpoints",
	)
	pc5.AddRoutes(parsecomRoutes)
	pc5.AddRequests(parsecomRoutes)
	pc5.AddTest(TestResponseOkWithoutOutput)
}

// Parse
// https://parse.com/docs/rest#summary
var parsecomRoutes = driver.Fixtures{
	// Objects
	{"POST", "/1/classes/:className"},
	{"GET", "/1/classes/:className/:objectId"},
	{"PUT", "/1/classes/:className/:objectId"},
	{"GET", "/1/classes/:className"},
	{"DELETE", "/1/classes/:className/:objectId"},

	// Users
	{"POST", "/1/users"},
	{"GET", "/1/login"},
	{"GET", "/1/users/:objectId"},
	{"PUT", "/1/users/:objectId"},
	{"GET", "/1/users"},
	{"DELETE", "/1/users/:objectId"},
	{"POST", "/1/requestPasswordReset"},

	// Roles
	{"POST", "/1/roles"},
	{"GET", "/1/roles/:objectId"},
	{"PUT", "/1/roles/:objectId"},
	{"GET", "/1/roles"},
	{"DELETE", "/1/roles/:objectId"},

	// Files
	{"POST", "/1/files/:fileName"},

	// Analytics
	{"POST", "/1/events/:eventName"},

	// Push Notifications
	{"POST", "/1/push"},

	// Installations
	{"POST", "/1/installations"},
	{"GET", "/1/installations/:objectId"},
	{"PUT", "/1/installations/:objectId"},
	{"GET", "/1/installations"},
	{"DELETE", "/1/installations/:objectId"},

	// Cloud Functions
	{"POST", "/1/functions"},
}
