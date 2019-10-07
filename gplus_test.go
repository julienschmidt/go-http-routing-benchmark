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
	gplusAce             http.Handler
	gplusBear            http.Handler
	gplusBeego           http.Handler
	gplusBone            http.Handler
	gplusChi             http.Handler
	gplusCloudyKitRouter http.Handler
	gplusDenco           http.Handler
	gplusEcho            http.Handler
	gplusGin             http.Handler
	gplusGocraftWeb      http.Handler
	gplusGoji            http.Handler
	gplusGojiv2          http.Handler
	gplusGoJsonRest      http.Handler
	gplusGoRestful       http.Handler
	gplusGorillaMux      http.Handler
	gplusGowwwRouter     http.Handler
	gplusHttpRouter      http.Handler
	gplusHttpTreeMux     http.Handler
	gplusKocha           http.Handler
	gplusLARS            http.Handler
	gplusMacaron         http.Handler
	gplusMartini         http.Handler
	gplusPat             http.Handler
	gplusPossum          http.Handler
	gplusR2router        http.Handler
	gplusRevel           http.Handler
	gplusRivet           http.Handler
	gplusTango           http.Handler
	gplusTigerTonic      http.Handler
	gplusTraffic         http.Handler
	gplusVulcan          http.Handler
	// gplusZeus        http.Handler
)

func init() {
	println("#GPlusAPI Routes:", len(gplusAPI))

	calcMem("Ace", func() {
		gplusAce = loadAce(gplusAPI)
	})
	calcMem("Bear", func() {
		gplusBear = loadBear(gplusAPI)
	})
	calcMem("Beego", func() {
		gplusBeego = loadBeego(gplusAPI)
	})
	calcMem("Bone", func() {
		gplusBone = loadBone(gplusAPI)
	})
	calcMem("Chi", func() {
		gplusChi = loadChi(gplusAPI)
	})
	calcMem("CloudyKitRouter", func() {
		gplusCloudyKitRouter = loadCloudyKitRouter(gplusAPI)
	})
	calcMem("Denco", func() {
		gplusDenco = loadDenco(gplusAPI)
	})
	calcMem("Echo", func() {
		gplusEcho = loadEcho(gplusAPI)
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
	calcMem("Gojiv2", func() {
		gplusGojiv2 = loadGojiv2(gplusAPI)
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
	calcMem("GowwwRouter", func() {
		gplusGowwwRouter = loadGowwwRouter(gplusAPI)
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
	calcMem("LARS", func() {
		gplusLARS = loadLARS(gplusAPI)
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
	calcMem("Possum", func() {
		gplusPossum = loadPossum(gplusAPI)
	})
	calcMem("R2router", func() {
		gplusR2router = loadR2router(gplusAPI)
	})
	// calcMem("Revel", func() {
	// 	gplusRevel = loadRevel(gplusAPI)
	// })
	calcMem("Rivet", func() {
		gplusRivet = loadRivet(gplusAPI)
	})
	calcMem("Tango", func() {
		gplusTango = loadTango(gplusAPI)
	})
	calcMem("TigerTonic", func() {
		gplusTigerTonic = loadTigerTonic(gplusAPI)
	})
	calcMem("Traffic", func() {
		gplusTraffic = loadTraffic(gplusAPI)
	})
	calcMem("Vulcan", func() {
		gplusVulcan = loadVulcan(gplusAPI)
	})
	// calcMem("Zeus", func() {
	// 	gplusZeus = loadZeus(gplusAPI)
	// })

	println()
}

// Static
func BenchmarkAce_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusAce, req)
}
func BenchmarkBear_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusBear, req)
}
func BenchmarkBeego_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkBone_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusBone, req)
}
func BenchmarkChi_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusChi, req)
}
func BenchmarkCloudyKitRouter_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusCloudyKitRouter, req)
}
func BenchmarkDenco_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkEcho_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusEcho, req)
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
func BenchmarkGojiv2_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGojiv2, req)
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
func BenchmarkGowwwRouter_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGowwwRouter, req)
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
func BenchmarkLARS_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusLARS, req)
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
func BenchmarkPossum_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusPossum, req)
}
func BenchmarkR2router_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusR2router, req)
}

// func BenchmarkRevel_GPlusStatic(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/people", nil)
// 	benchRequest(b, gplusRevel, req)
// }
func BenchmarkRivet_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusRivet, req)
}
func BenchmarkTango_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusTango, req)
}
func BenchmarkTigerTonic_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusTraffic, req)
}
func BenchmarkVulcan_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusVulcan, req)
}

// func BenchmarkZeus_GPlusStatic(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/people", nil)
// 	benchRequest(b, gplusZeus, req)
// }

// One Param
func BenchmarkAce_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusAce, req)
}
func BenchmarkBear_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusBear, req)
}
func BenchmarkBeego_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkBone_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusBone, req)
}
func BenchmarkChi_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusChi, req)
}
func BenchmarkCloudyKitRouter_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusCloudyKitRouter, req)
}
func BenchmarkDenco_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkEcho_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusEcho, req)
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
func BenchmarkGojiv2_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGojiv2, req)
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
func BenchmarkGowwwRouter_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGowwwRouter, req)
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
func BenchmarkLARS_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusLARS, req)
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
func BenchmarkPossum_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusPossum, req)
}
func BenchmarkR2router_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusR2router, req)
}

// func BenchmarkRevel_GPlusParam(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
// 	benchRequest(b, gplusRevel, req)
// }
func BenchmarkRivet_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusRivet, req)
}
func BenchmarkTango_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusTango, req)
}
func BenchmarkTigerTonic_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusTraffic, req)
}
func BenchmarkVulcan_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusVulcan, req)
}

// func BenchmarkZeus_GPlusParam(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
// 	benchRequest(b, gplusZeus, req)
// }

// Two Params
func BenchmarkAce_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusAce, req)
}
func BenchmarkBear_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusBear, req)
}
func BenchmarkBeego_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusBeego, req)
}
func BenchmarkBone_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusBone, req)
}
func BenchmarkChi_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusChi, req)
}
func BenchmarkCloudyKitRouter_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusCloudyKitRouter, req)
}
func BenchmarkDenco_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkEcho_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusEcho, req)
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
func BenchmarkGojiv2_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGojiv2, req)
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
func BenchmarkGowwwRouter_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGowwwRouter, req)
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
func BenchmarkLARS_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusLARS, req)
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
func BenchmarkPossum_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusPossum, req)
}
func BenchmarkR2router_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusR2router, req)
}

// func BenchmarkRevel_GPlus2Params(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
// 	benchRequest(b, gplusRevel, req)
// }
func BenchmarkRivet_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusRivet, req)
}
func BenchmarkTango_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusTango, req)
}
func BenchmarkTigerTonic_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusTraffic, req)
}
func BenchmarkVulcan_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusVulcan, req)
}

// func BenchmarkZeus_GPlus2Params(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
// 	benchRequest(b, gplusZeus, req)
// }

// All Routes
func BenchmarkAce_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusAce, gplusAPI)
}
func BenchmarkBear_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusBear, gplusAPI)
}
func BenchmarkBeego_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusBeego, gplusAPI)
}
func BenchmarkBone_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusBone, gplusAPI)
}
func BenchmarkChi_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusChi, gplusAPI)
}
func BenchmarkCloudyKitRouter_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusCloudyKitRouter, gplusAPI)
}
func BenchmarkDenco_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusDenco, gplusAPI)
}
func BenchmarkEcho_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusEcho, gplusAPI)
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
func BenchmarkGojiv2_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGojiv2, gplusAPI)
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
func BenchmarkGowwwRouter_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGowwwRouter, gplusAPI)
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
func BenchmarkLARS_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusLARS, gplusAPI)
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
func BenchmarkPossum_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusPossum, gplusAPI)
}
func BenchmarkR2router_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusR2router, gplusAPI)
}

// func BenchmarkRevel_GPlusAll(b *testing.B) {
// 	benchRoutes(b, gplusRevel, gplusAPI)
// }
func BenchmarkRivet_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusRivet, gplusAPI)
}
func BenchmarkTango_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusTango, gplusAPI)
}
func BenchmarkTigerTonic_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusTigerTonic, gplusAPI)
}
func BenchmarkTraffic_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusTraffic, gplusAPI)
}
func BenchmarkVulcan_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusVulcan, gplusAPI)
}

// func BenchmarkZeus_GPlusAll(b *testing.B) {
// 	benchRoutes(b, gplusZeus, gplusAPI)
// }
