// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// Parse
// https://parse.com/docs/rest#summary
var parseAPI = []route{
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

var (
	parseAce             http.Handler
	parseAero            http.Handler
	parseBear            http.Handler
	parseBeego           http.Handler
	parseBone            http.Handler
	parseChi             http.Handler
	parseCloudyKitRouter http.Handler
	parseDenco           http.Handler
	parseEcho            http.Handler
	parseGin             http.Handler
	parseGocraftWeb      http.Handler
	parseGoji            http.Handler
	parseGojiv2          http.Handler
	parseGoJsonRest      http.Handler
	parseGoRestful       http.Handler
	parseGorillaMux      http.Handler
	parseGowwwRouter     http.Handler
	parseHttpRouter      http.Handler
	parseHttpTreeMux     http.Handler
	parseKocha           http.Handler
	parseLARS            http.Handler
	parseLenrouter       http.Handler
	parseMacaron         http.Handler
	parseMartini         http.Handler
	parsePat             http.Handler
	parsePossum          http.Handler
	parseR2router        http.Handler
	parseRevel           http.Handler
	parseRivet           http.Handler
	parseTango           http.Handler
	parseTigerTonic      http.Handler
	parseTraffic         http.Handler
	parseVulcan          http.Handler
	// parseZeus        http.Handler
)

func init() {
	println("#ParseAPI Routes:", len(parseAPI))

	calcMem("Ace", func() {
		parseAce = loadAce(parseAPI)
	})
	calcMem("Aero", func() {
		parseAero = loadAero(parseAPI)
	})
	calcMem("Bear", func() {
		parseBear = loadBear(parseAPI)
	})
	calcMem("Beego", func() {
		parseBeego = loadBeego(parseAPI)
	})
	calcMem("Bone", func() {
		parseBone = loadBone(parseAPI)
	})
	calcMem("Chi", func() {
		parseChi = loadChi(parseAPI)
	})
	calcMem("CloudyKitRouter", func() {
		parseCloudyKitRouter = loadCloudyKitRouter(parseAPI)
	})
	calcMem("Denco", func() {
		parseDenco = loadDenco(parseAPI)
	})
	calcMem("Echo", func() {
		parseEcho = loadEcho(parseAPI)
	})
	calcMem("Gin", func() {
		parseGin = loadGin(parseAPI)
	})
	calcMem("GocraftWeb", func() {
		parseGocraftWeb = loadGocraftWeb(parseAPI)
	})
	calcMem("Goji", func() {
		parseGoji = loadGoji(parseAPI)
	})
	calcMem("Gojiv2", func() {
		parseGojiv2 = loadGojiv2(parseAPI)
	})
	calcMem("GoJsonRest", func() {
		parseGoJsonRest = loadGoJsonRest(parseAPI)
	})
	calcMem("GoRestful", func() {
		parseGoRestful = loadGoRestful(parseAPI)
	})
	calcMem("GorillaMux", func() {
		parseGorillaMux = loadGorillaMux(parseAPI)
	})
	calcMem("GowwwRouter", func() {
		parseGowwwRouter = loadGowwwRouter(parseAPI)
	})
	calcMem("HttpRouter", func() {
		parseHttpRouter = loadHttpRouter(parseAPI)
	})
	calcMem("HttpTreeMux", func() {
		parseHttpTreeMux = loadHttpTreeMux(parseAPI)
	})
	calcMem("Kocha", func() {
		parseKocha = loadKocha(parseAPI)
	})
	calcMem("LARS", func() {
		parseLARS = loadLARS(parseAPI)
	})
	calcMem("Lenrouter", func() {
		parseLenrouter = loadLenrouter(parseAPI)
	})
	calcMem("Macaron", func() {
		parseMacaron = loadMacaron(parseAPI)
	})
	calcMem("Martini", func() {
		parseMartini = loadMartini(parseAPI)
	})
	calcMem("Pat", func() {
		parsePat = loadPat(parseAPI)
	})
	calcMem("Possum", func() {
		parsePossum = loadPossum(parseAPI)
	})
	calcMem("R2router", func() {
		parseR2router = loadR2router(parseAPI)
	})
	// calcMem("Revel", func() {
	// 	parseRevel = loadRevel(parseAPI)
	// })
	calcMem("Rivet", func() {
		parseRivet = loadRivet(parseAPI)
	})
	calcMem("Tango", func() {
		parseTango = loadTango(parseAPI)
	})
	calcMem("TigerTonic", func() {
		parseTigerTonic = loadTigerTonic(parseAPI)
	})
	calcMem("Traffic", func() {
		parseTraffic = loadTraffic(parseAPI)
	})
	calcMem("Vulcan", func() {
		parseVulcan = loadVulcan(parseAPI)
	})
	// calcMem("Zeus", func() {
	// 	parseZeus = loadZeus(parseAPI)
	// })

	println()
}

// Static
func BenchmarkAce_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseAce, req)
}
func BenchmarkAero_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseAero, req)
}
func BenchmarkBear_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseBear, req)
}
func BenchmarkBeego_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseBeego, req)
}
func BenchmarkBone_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseBone, req)
}
func BenchmarkChi_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseChi, req)
}
func BenchmarkCloudyKitRouter_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseCloudyKitRouter, req)
}
func BenchmarkDenco_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseDenco, req)
}
func BenchmarkEcho_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseEcho, req)
}
func BenchmarkGin_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGin, req)
}
func BenchmarkGocraftWeb_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGocraftWeb, req)
}
func BenchmarkGoji_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGoji, req)
}
func BenchmarkGojiv2_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGojiv2, req)
}
func BenchmarkGoJsonRest_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGoJsonRest, req)
}
func BenchmarkGoRestful_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGoRestful, req)
}
func BenchmarkGorillaMux_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkGowwwRouter_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseGowwwRouter, req)
}
func BenchmarkHttpRouter_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkHttpTreeMux_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseHttpTreeMux, req)
}
func BenchmarkKocha_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseKocha, req)
}
func BenchmarkLARS_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseLARS, req)
}
func BenchmarkLenrouter_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseLenrouter, req)
}
func BenchmarkMacaron_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseMacaron, req)
}
func BenchmarkMartini_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseMartini, req)
}
func BenchmarkPat_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parsePat, req)
}
func BenchmarkPossum_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parsePossum, req)
}
func BenchmarkR2router_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseR2router, req)
}

// func BenchmarkRevel_ParseStatic(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/1/users", nil)
// 	benchRequest(b, parseRevel, req)
// }
func BenchmarkRivet_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseRivet, req)
}
func BenchmarkTango_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseTango, req)
}
func BenchmarkTigerTonic_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseTigerTonic, req)
}
func BenchmarkTraffic_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseTraffic, req)
}
func BenchmarkVulcan_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseVulcan, req)
}

// func BenchmarkZeus_ParseStatic(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/1/users", nil)
// 	benchRequest(b, parseZeus, req)
// }

// One Param
func BenchmarkAce_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseAce, req)
}
func BenchmarkAero_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseAero, req)
}
func BenchmarkBear_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseBear, req)
}
func BenchmarkBeego_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseBeego, req)
}
func BenchmarkBone_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseBone, req)
}
func BenchmarkChi_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseChi, req)
}
func BenchmarkCloudyKitRouter_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseCloudyKitRouter, req)
}
func BenchmarkDenco_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseDenco, req)
}
func BenchmarkEcho_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseEcho, req)
}
func BenchmarkGin_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGin, req)
}
func BenchmarkGocraftWeb_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGocraftWeb, req)
}
func BenchmarkGoji_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGoji, req)
}
func BenchmarkGojiv2_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGojiv2, req)
}
func BenchmarkGoJsonRest_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGoJsonRest, req)
}
func BenchmarkGoRestful_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGoRestful, req)
}
func BenchmarkGorillaMux_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkGowwwRouter_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseGowwwRouter, req)
}
func BenchmarkHttpRouter_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkHttpTreeMux_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseHttpTreeMux, req)
}
func BenchmarkKocha_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseKocha, req)
}
func BenchmarkLARS_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseLARS, req)
}
func BenchmarkLenrouter_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseLenrouter, req)
}
func BenchmarkMacaron_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseMacaron, req)
}
func BenchmarkMartini_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseMartini, req)
}
func BenchmarkPat_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parsePat, req)
}
func BenchmarkPossum_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parsePossum, req)
}
func BenchmarkR2router_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseR2router, req)
}

// func BenchmarkRevel_ParseParam(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
// 	benchRequest(b, parseRevel, req)
// }
func BenchmarkRivet_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseRivet, req)
}
func BenchmarkTango_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseTango, req)
}
func BenchmarkTigerTonic_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseTigerTonic, req)
}
func BenchmarkTraffic_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseTraffic, req)
}
func BenchmarkVulcan_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseVulcan, req)
}

// func BenchmarkZeus_ParseParam(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
// 	benchRequest(b, parseZeus, req)
// }

// Two Params
func BenchmarkAce_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseAce, req)
}
func BenchmarkAero_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseAero, req)
}
func BenchmarkBear_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseBear, req)
}
func BenchmarkBeego_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseBeego, req)
}
func BenchmarkBone_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseBone, req)
}
func BenchmarkChi_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseChi, req)
}
func BenchmarkCloudyKitRouter_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseCloudyKitRouter, req)
}
func BenchmarkDenco_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseDenco, req)
}
func BenchmarkEcho_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseEcho, req)
}
func BenchmarkGin_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGin, req)
}
func BenchmarkGocraftWeb_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGocraftWeb, req)
}
func BenchmarkGoji_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGoji, req)
}
func BenchmarkGojiv2_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGojiv2, req)
}
func BenchmarkGoJsonRest_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGoJsonRest, req)
}
func BenchmarkGoRestful_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGoRestful, req)
}
func BenchmarkGorillaMux_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGorillaMux, req)
}
func BenchmarkGowwwRouter_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseGowwwRouter, req)
}
func BenchmarkHttpRouter_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseHttpRouter, req)
}
func BenchmarkHttpTreeMux_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseHttpTreeMux, req)
}
func BenchmarkKocha_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseKocha, req)
}
func BenchmarkLARS_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseLARS, req)
}
func BenchmarkLenrouter_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseLenrouter, req)
}
func BenchmarkMacaron_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseMacaron, req)
}
func BenchmarkMartini_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseMartini, req)
}
func BenchmarkPat_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parsePat, req)
}
func BenchmarkPossum_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parsePossum, req)
}
func BenchmarkR2router_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseR2router, req)
}

// func BenchmarkRevel_Parse2Params(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
// 	benchRequest(b, parseRevel, req)
// }
func BenchmarkRivet_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseRivet, req)
}
func BenchmarkTango_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseTango, req)
}
func BenchmarkTigerTonic_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseTigerTonic, req)
}
func BenchmarkTraffic_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseTraffic, req)
}
func BenchmarkVulcan_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseVulcan, req)
}

// func BenchmarkZeus_Parse2Params(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
// 	benchRequest(b, parseZeus, req)
// }

// All Routes
func BenchmarkAce_ParseAll(b *testing.B) {
	benchRoutes(b, parseAce, parseAPI)
}
func BenchmarkAero_ParseAll(b *testing.B) {
	benchRoutes(b, parseAero, parseAPI)
}
func BenchmarkBear_ParseAll(b *testing.B) {
	benchRoutes(b, parseBear, parseAPI)
}
func BenchmarkBeego_ParseAll(b *testing.B) {
	benchRoutes(b, parseBeego, parseAPI)
}
func BenchmarkBone_ParseAll(b *testing.B) {
	benchRoutes(b, parseBone, parseAPI)
}
func BenchmarkChi_ParseAll(b *testing.B) {
	benchRoutes(b, parseChi, parseAPI)
}
func BenchmarkCloudyKitRouter_ParseAll(b *testing.B) {
	benchRoutes(b, parseCloudyKitRouter, parseAPI)
}
func BenchmarkDenco_ParseAll(b *testing.B) {
	benchRoutes(b, parseDenco, parseAPI)
}
func BenchmarkEcho_ParseAll(b *testing.B) {
	benchRoutes(b, parseEcho, parseAPI)
}
func BenchmarkGin_ParseAll(b *testing.B) {
	benchRoutes(b, parseGin, parseAPI)
}
func BenchmarkGocraftWeb_ParseAll(b *testing.B) {
	benchRoutes(b, parseGocraftWeb, parseAPI)
}
func BenchmarkGoji_ParseAll(b *testing.B) {
	benchRoutes(b, parseGoji, parseAPI)
}
func BenchmarkGojiv2_ParseAll(b *testing.B) {
	benchRoutes(b, parseGojiv2, parseAPI)
}
func BenchmarkGoJsonRest_ParseAll(b *testing.B) {
	benchRoutes(b, parseGoJsonRest, parseAPI)
}
func BenchmarkGoRestful_ParseAll(b *testing.B) {
	benchRoutes(b, parseGoRestful, parseAPI)
}
func BenchmarkGorillaMux_ParseAll(b *testing.B) {
	benchRoutes(b, parseGorillaMux, parseAPI)
}
func BenchmarkGowwwRouter_ParseAll(b *testing.B) {
	benchRoutes(b, parseGowwwRouter, parseAPI)
}
func BenchmarkHttpRouter_ParseAll(b *testing.B) {
	benchRoutes(b, parseHttpRouter, parseAPI)
}
func BenchmarkHttpTreeMux_ParseAll(b *testing.B) {
	benchRoutes(b, parseHttpTreeMux, parseAPI)
}
func BenchmarkKocha_ParseAll(b *testing.B) {
	benchRoutes(b, parseKocha, parseAPI)
}
func BenchmarkLARS_ParseAll(b *testing.B) {
	benchRoutes(b, parseLARS, parseAPI)
}
func BenchmarkLenrouter_ParseAll(b *testing.B) {
	benchRoutes(b, parseLenrouter, parseAPI)
}
func BenchmarkMacaron_ParseAll(b *testing.B) {
	benchRoutes(b, parseMacaron, parseAPI)
}
func BenchmarkMartini_ParseAll(b *testing.B) {
	benchRoutes(b, parseMartini, parseAPI)
}
func BenchmarkPat_ParseAll(b *testing.B) {
	benchRoutes(b, parsePat, parseAPI)
}
func BenchmarkPossum_ParseAll(b *testing.B) {
	benchRoutes(b, parsePossum, parseAPI)
}
func BenchmarkR2router_ParseAll(b *testing.B) {
	benchRoutes(b, parseR2router, parseAPI)
}

// func BenchmarkRevel_ParseAll(b *testing.B) {
// 	benchRoutes(b, parseRevel, parseAPI)
// }
func BenchmarkRivet_ParseAll(b *testing.B) {
	benchRoutes(b, parseRivet, parseAPI)
}
func BenchmarkTango_ParseAll(b *testing.B) {
	benchRoutes(b, parseTango, parseAPI)
}
func BenchmarkTigerTonic_ParseAll(b *testing.B) {
	benchRoutes(b, parseTigerTonic, parseAPI)
}
func BenchmarkTraffic_ParseAll(b *testing.B) {
	benchRoutes(b, parseTraffic, parseAPI)
}
func BenchmarkVulcan_ParseAll(b *testing.B) {
	benchRoutes(b, parseVulcan, parseAPI)
}

// func BenchmarkZeus_ParseAll(b *testing.B) {
// 	benchRoutes(b, parseZeus, parseAPI)
// }
