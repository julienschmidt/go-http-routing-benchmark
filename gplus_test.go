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
	gplusAce         http.Handler
	gplusBeego       http.Handler
	gplusBone        http.Handler
	gplusDenco       http.Handler
	gplusGin         http.Handler
	gplusGocraftWeb  http.Handler
	gplusGoji        http.Handler
	gplusGoJsonRest  http.Handler
	gplusGoRestful   http.Handler
	gplusGorillaMux  http.Handler
	gplusHttpRouter  http.Handler
	gplusHttpTreeMux http.Handler
	gplusKocha       http.Handler
	gplusMacaron     http.Handler
	gplusMartini     http.Handler
	gplusPat         http.Handler
	gplusRevel       http.Handler
	gplusRivet       http.Handler
	gplusTigerTonic  http.Handler
	gplusTraffic     http.Handler
	gplusZeus        http.Handler
)

func init() {
	println("#GPlusAPI Routes:", len(gplusAPI))

	calcMem("Ace", func() {
		gplusAce = loadAce(gplusAPI)
	})
	calcMem("Beego", func() {
		gplusBeego = loadBeego(gplusAPI)
	})
	calcMem("Bone", func() {
		gplusBone = loadBone(gplusAPI)
	})
	calcMem("Denco", func() {
		gplusDenco = loadDenco(gplusAPI)
	})
	calcMem("Gin", func() {
		gplusGin = loadGin(gplusAPI)
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
	calcMem("GoRestful", func() {
		gplusGoRestful = loadGoRestful(gplusAPI)
	})
	calcMem("GorillaMux", func() {
		gplusGorillaMux = loadGorillaMux(gplusAPI)
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
	calcMem("Macaron", func() {
		gplusMacaron = loadMacaron(gplusAPI)
	})
	calcMem("Martini", func() {
		gplusMartini = loadMartini(gplusAPI)
	})
	calcMem("Pat", func() {
		gplusPat = loadPat(gplusAPI)
	})
	calcMem("Revel", func() {
		gplusRevel = loadRevel(gplusAPI)
	})
	calcMem("Rivet", func() {
		gplusRivet = loadRivet(gplusAPI)
	})
	calcMem("TigerTonic", func() {
		gplusTigerTonic = loadTigerTonic(gplusAPI)
	})
	calcMem("Traffic", func() {
		gplusTraffic = loadTraffic(gplusAPI)
	})
	calcMem("Zeus", func() {
		gplusZeus = loadZeus(gplusAPI)
	})

	println()
}

// Static
func BenchmarkAce_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusAce, req)
}
func BenchmarkBeego_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkBone_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusBone, req)
}
func BenchmarkDenco_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkGin_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGin, req)
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
func BenchmarkGoRestful_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGoRestful, req)
}
func BenchmarkGorillaMux_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGorillaMux, req)
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
func BenchmarkMacaron_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusMacaron, req)
}
func BenchmarkMartini_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPat_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusPat, req)
}
func BenchmarkRevel_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusRevel, req)
}
func BenchmarkRivet_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusRivet, req)
}
func BenchmarkTigerTonic_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusTraffic, req)
}
func BenchmarkZeus_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusZeus, req)
}

// One Param
func BenchmarkAce_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusAce, req)
}
func BenchmarkBeego_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkBone_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusBone, req)
}
func BenchmarkDenco_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkGin_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGin, req)
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
func BenchmarkGoRestful_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGoRestful, req)
}
func BenchmarkGorillaMux_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGorillaMux, req)
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
func BenchmarkMacaron_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusMacaron, req)
}
func BenchmarkMartini_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPat_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusPat, req)
}
func BenchmarkRevel_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusRevel, req)
}
func BenchmarkRivet_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusRivet, req)
}
func BenchmarkTigerTonic_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusTraffic, req)
}
func BenchmarkZeus_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusZeus, req)
}

// Two Params
func BenchmarkAce_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusAce, req)
}
func BenchmarkBeego_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkBone_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusBone, req)
}
func BenchmarkDenco_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkGin_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGin, req)
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
func BenchmarkGoRestful_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGoRestful, req)
}
func BenchmarkGorillaMux_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGorillaMux, req)
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
func BenchmarkMacaron_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusMacaron, req)
}
func BenchmarkMartini_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPat_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusPat, req)
}
func BenchmarkRevel_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusRevel, req)
}
func BenchmarkRivet_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusRivet, req)
}
func BenchmarkTigerTonic_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusTraffic, req)
}
func BenchmarkZeus_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusZeus, req)
}

// All Routes
func BenchmarkAce_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusAce, gplusAPI)
}
func BenchmarkBeego_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusBeego, gplusAPI)
}
func BenchmarkBone_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusBone, gplusAPI)
}
func BenchmarkDenco_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusDenco, gplusAPI)
}
func BenchmarkGin_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGin, gplusAPI)
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
func BenchmarkGoRestful_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGoRestful, gplusAPI)
}
func BenchmarkGorillaMux_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGorillaMux, gplusAPI)
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
func BenchmarkMacaron_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusMacaron, gplusAPI)
}
func BenchmarkMartini_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusMartini, gplusAPI)
}
func BenchmarkPat_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusPat, gplusAPI)
}
func BenchmarkRevel_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusRevel, gplusAPI)
}
func BenchmarkRivet_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusRivet, gplusAPI)
}
func BenchmarkTigerTonic_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusTigerTonic, gplusAPI)
}
func BenchmarkTraffic_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusTraffic, gplusAPI)
}
func BenchmarkZeus_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusZeus, gplusAPI)
}
