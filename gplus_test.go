// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// Google+
// https://developers.google.com/+/api/latest/
// (in reality this is just a subset of a much larger API)
var gplusAPI = []route{
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

var (
	gplusBeego       http.Handler
	gplusDenco       http.Handler
	gplusGocraftWeb  http.Handler
	gplusGoji        http.Handler
	gplusGoJsonRest  http.Handler
	gplusGorillaMux  http.Handler
	gplusGin         http.Handler
	gplusHttpRouter  http.Handler
	gplusHttpTreeMux http.Handler
	gplusKocha       http.Handler
	gplusMartini     http.Handler
	gplusPat         http.Handler
	gplusTigerTonic  http.Handler
	gplusTraffic     http.Handler
)

func init() {
	println("#GPlusAPI Routes:", len(gplusAPI))

	calcMem("Beego", func() {
		gplusBeego = loadBeego(gplusAPI)
	})
	calcMem("Denco", func() {
		gplusDenco = loadDenco(gplusAPI)
	})
	calcMem("GocraftWeb", func() {
		gplusGocraftWeb = loadGocraftWeb(gplusAPI)
	})
	calcMem("Goji", func() {
		gplusGoji = loadGoji(gplusAPI)
	})
	calcMem("GoJsonRest", func() {
		gplusGoJsonRest = loadGoJsonRest(gplusAPI)
	})
	calcMem("GorillaMux", func() {
		gplusGorillaMux = loadGorillaMux(gplusAPI)
	})
	calcMem("Gin", func() {
		gplusGin = loadGin(gplusAPI)
	})
	calcMem("HttpRouter", func() {
		gplusHttpRouter = loadHttpRouter(gplusAPI)
	})
	calcMem("HttpTreeMux", func() {
		gplusHttpTreeMux = loadHttpTreeMux(gplusAPI)
	})
	calcMem("Kocha", func() {
		gplusKocha = loadKocha(gplusAPI)
	})
	calcMem("Martini", func() {
		gplusMartini = loadMartini(gplusAPI)
	})
	calcMem("Pat", func() {
		gplusPat = loadPat(gplusAPI)
	})
	calcMem("TigerTonic", func() {
		gplusTigerTonic = loadTigerTonic(gplusAPI)
	})
	calcMem("Traffic", func() {
		gplusTraffic = loadTraffic(gplusAPI)
	})

	println()
}

// Static
func BenchmarkBeego_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkDenco_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkGocraftWeb_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGocraftWeb, req)
}
func BenchmarkGoji_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGoji, req)
}
func BenchmarkGoJsonRest_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGoJsonRest, req)
}
func BenchmarkGorillaMux_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkGin_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGin, req)
}
func BenchmarkHttpRouter_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkKocha_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusKocha, req)
}
func BenchmarkMartini_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPat_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusPat, req)
}
func BenchmarkTigerTonic_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusTraffic, req)
}

// One Param
func BenchmarkBeego_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkDenco_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkGocraftWeb_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGocraftWeb, req)
}
func BenchmarkGoji_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGoji, req)
}
func BenchmarkGoJsonRest_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGoJsonRest, req)
}
func BenchmarkGorillaMux_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkGin_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGin, req)
}
func BenchmarkHttpRouter_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkKocha_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusKocha, req)
}
func BenchmarkMartini_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPat_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusPat, req)
}
func BenchmarkTigerTonic_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusTraffic, req)
}

// Two Params
func BenchmarkBeego_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkDenco_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkGocraftWeb_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGocraftWeb, req)
}
func BenchmarkGoji_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGoji, req)
}
func BenchmarkGoJsonRest_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGoJsonRest, req)
}
func BenchmarkGorillaMux_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkGin_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGin, req)
}
func BenchmarkHttpRouter_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkKocha_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusKocha, req)
}
func BenchmarkMartini_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPat_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusPat, req)
}
func BenchmarkTigerTonic_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusTraffic, req)
}

// All Routes
func BenchmarkBeego_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusBeego, gplusAPI)
}
func BenchmarkDenco_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusDenco, gplusAPI)
}
func BenchmarkGocraftWeb_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGocraftWeb, gplusAPI)
}
func BenchmarkGoji_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGoji, gplusAPI)
}
func BenchmarkGoJsonRest_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGoJsonRest, gplusAPI)
}
func BenchmarkGorillaMux_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGorillaMux, gplusAPI)
}
func BenchmarkGin_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGin, gplusAPI)
}
func BenchmarkHttpRouter_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusHttpRouter, gplusAPI)
}
func BenchmarkHttpTreeMux_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusHttpTreeMux, gplusAPI)
}
func BenchmarkKocha_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusKocha, gplusAPI)
}
func BenchmarkMartini_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusMartini, gplusAPI)
}
func BenchmarkPat_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusPat, gplusAPI)
}
func BenchmarkTigerTonic_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusTigerTonic, gplusAPI)
}
func BenchmarkTraffic_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusTraffic, gplusAPI)
}
