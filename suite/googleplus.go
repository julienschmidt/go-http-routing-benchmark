// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package suite

import "github.com/julienschmidt/go-http-routing-benchmark/driver"

func init() {
	gp1 := driver.NewBenchmark(
		driver.Parameterized,
		"GooglePlusStatic",
		"Request one static Google+ endpoint with one segment",
	)
	gp1.AddRoutes(googleplusRoutes)
	gp1.AddTest(TestResponseOkWithoutOutput)
	gp1.AddRequest("GET", "/people")

	gp2 := driver.NewBenchmark(
		driver.Parameterized,
		"GooglePlusParam",
		"Request one parameterized Google+ endpoint with two segments, one of which is a parameter",
	)
	gp2.AddRoutes(googleplusRoutes)
	gp2.AddTest(TestResponseOkWithoutOutput)
	gp2.AddRequest("GET", "/people/118051310819094153327")

	gp3 := driver.NewBenchmark(
		driver.Parameterized,
		"GooglePlus2Params",
		"Request one parameterized Google+ endpoint with four segments, two of which are parameters",
	)
	gp3.AddRoutes(googleplusRoutes)
	gp3.AddTest(TestResponseOkWithoutOutput)
	gp3.AddRequest("GET", "/people/118051310819094153327/activities/123456789")

	gp4 := driver.NewBenchmark(
		driver.Parameterized,
		"GooglePlusNotFound",
		"Request one unavailable Google+ endpoint",
	)
	gp4.AddRoutes(googleplusRoutes)
	gp4.AddTest(TestResponseNotFound)
	gp4.AddRequest("GET", "/notfound/this-path-is-unavailable") // Without trailing slash!

	gp5 := driver.NewBenchmark(
		driver.Parameterized,
		"GooglePlusAll",
		"Request all Google+ endpoints",
	)
	gp5.AddRoutes(googleplusRoutes)
	gp5.AddRequests(googleplusRoutes)
	gp5.AddTest(TestResponseOkWithoutOutput)
}

// Google+
// https://developers.google.com/+/api/latest/
// (in reality this is just a subset of a much larger API)
var googleplusRoutes = driver.Fixtures{
	// People
	{"GET", "/people/:userId"},
	{"GET", "/people"},
	{"GET", "/activities/:activityId/people/:collection"},
	{"GET", "/people/:userId/people/:collection"},
	{"GET", "/people/:userId/openIdConnect"},

	// Activities
	{"GET", "/people/:userId/activities/:collection"},
	{"GET", "/activities/:activityId"},
	{"GET", "/activities"},

	// Comments
	{"GET", "/activities/:activityId/comments"},
	{"GET", "/comments/:commentId"},

	// Moments
	{"POST", "/people/:userId/moments/:collection"},
	{"GET", "/people/:userId/moments/:collection"},
	{"DELETE", "/moments/:id"},
}
