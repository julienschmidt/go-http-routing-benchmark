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
	parseAce         http.Handler
	parseBeego       http.Handler
	parseBone        http.Handler
	parseDenco       http.Handler
	parseGin         http.Handler
	parseGocraftWeb  http.Handler
	parseGoji        http.Handler
	parseGoJsonRest  http.Handler
	parseGoRestful   http.Handler
	parseGorillaMux  http.Handler
	parseHttpRouter  http.Handler
	parseHttpTreeMux http.Handler
	parseKocha       http.Handler
	parseMacaron     http.Handler
	parseMartini     http.Handler
	parsePat         http.Handler
	parseRevel       http.Handler
	parseRivet       http.Handler
	parseTigerTonic  http.Handler
	parseTraffic     http.Handler
	parseZeus        http.Handler
)

func init() {
	println("#ParseAPI Routes:", len(parseAPI))

	calcMem("Ace", func() {
		parseAce = loadAce(parseAPI)
	})
	calcMem("Beego", func() {
		parseBeego = loadBeego(parseAPI)
	})
	calcMem("Bone", func() {
		parseBone = loadBone(parseAPI)
	})
	calcMem("Denco", func() {
		parseDenco = loadDenco(parseAPI)
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
	calcMem("GoJsonRest", func() {
		parseGoJsonRest = loadGoJsonRest(parseAPI)
	})
	calcMem("GoRestful", func() {
		parseGoRestful = loadGoRestful(parseAPI)
	})
	calcMem("GorillaMux", func() {
		parseGorillaMux = loadGorillaMux(parseAPI)
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
	calcMem("Macaron", func() {
		parseMacaron = loadMacaron(parseAPI)
	})
	calcMem("Martini", func() {
		parseMartini = loadMartini(parseAPI)
	})
	calcMem("Pat", func() {
		parsePat = loadPat(parseAPI)
	})
	calcMem("Revel", func() {
		parseRevel = loadRevel(parseAPI)
	})
	calcMem("Rivet", func() {
		parseRivet = loadRivet(parseAPI)
	})
	calcMem("TigerTonic", func() {
		parseTigerTonic = loadTigerTonic(parseAPI)
	})
	calcMem("Traffic", func() {
		parseTraffic = loadTraffic(parseAPI)
	})
	calcMem("Zeus", func() {
		parseZeus = loadZeus(parseAPI)
	})

	println()
}

// Static
func BenchmarkAce_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseAce, req)
}
func BenchmarkBeego_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseBeego, req)
}
func BenchmarkBone_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseBone, req)
}
func BenchmarkDenco_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseDenco, req)
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
func BenchmarkRevel_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseRevel, req)
}
func BenchmarkRivet_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseRivet, req)
}
func BenchmarkTigerTonic_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseTigerTonic, req)
}
func BenchmarkTraffic_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseTraffic, req)
}
func BenchmarkZeus_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseZeus, req)
}

// One Param
func BenchmarkAce_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseAce, req)
}
func BenchmarkBeego_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseBeego, req)
}
func BenchmarkBone_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseBone, req)
}
func BenchmarkDenco_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseDenco, req)
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
func BenchmarkMacaron_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseMacaron, req)
}
func BenchmarkMartini_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseMartini, req)
}
func BenchmarkRevel_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseRevel, req)
}
func BenchmarkRivet_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseRivet, req)
}
func BenchmarkPat_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parsePat, req)
}
func BenchmarkTigerTonic_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseTigerTonic, req)
}
func BenchmarkTraffic_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseTraffic, req)
}
func BenchmarkZeus_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseZeus, req)
}

// Two Params
func BenchmarkAce_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseAce, req)
}
func BenchmarkBeego_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseBeego, req)
}
func BenchmarkBone_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseBone, req)
}
func BenchmarkDenco_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseDenco, req)
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
func BenchmarkMacaron_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseMacaron, req)
}
func BenchmarkMartini_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseMartini, req)
}
func BenchmarkRevel_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseRevel, req)
}
func BenchmarkRivet_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseRivet, req)
}
func BenchmarkPat_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parsePat, req)
}
func BenchmarkTigerTonic_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseTigerTonic, req)
}
func BenchmarkTraffic_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseTraffic, req)
}
func BenchmarkZeus_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseZeus, req)
}

// All Routes
func BenchmarkAce_ParseAll(b *testing.B) {
	benchRoutes(b, parseAce, parseAPI)
}
func BenchmarkBeego_ParseAll(b *testing.B) {
	benchRoutes(b, parseBeego, parseAPI)
}
func BenchmarkBone_ParseAll(b *testing.B) {
	benchRoutes(b, parseBone, parseAPI)
}
func BenchmarkDenco_ParseAll(b *testing.B) {
	benchRoutes(b, parseDenco, parseAPI)
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
func BenchmarkGoJsonRest_ParseAll(b *testing.B) {
	benchRoutes(b, parseGoJsonRest, parseAPI)
}
func BenchmarkGoRestful_ParseAll(b *testing.B) {
	benchRoutes(b, parseGoRestful, parseAPI)
}
func BenchmarkGorillaMux_ParseAll(b *testing.B) {
	benchRoutes(b, parseGorillaMux, parseAPI)
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
func BenchmarkMacaron_ParseAll(b *testing.B) {
	benchRoutes(b, parseMacaron, parseAPI)
}
func BenchmarkMartini_ParseAll(b *testing.B) {
	benchRoutes(b, parseMartini, parseAPI)
}
func BenchmarkRevel_ParseAll(b *testing.B) {
	benchRoutes(b, parseRevel, parseAPI)
}
func BenchmarkRivet_ParseAll(b *testing.B) {
	benchRoutes(b, parseRivet, parseAPI)
}
func BenchmarkPat_ParseAll(b *testing.B) {
	benchRoutes(b, parsePat, parseAPI)
}
func BenchmarkTigerTonic_ParseAll(b *testing.B) {
	benchRoutes(b, parseTigerTonic, parseAPI)
}
func BenchmarkTraffic_ParseAll(b *testing.B) {
	benchRoutes(b, parseTraffic, parseAPI)
}
func BenchmarkZeus_ParseAll(b *testing.B) {
	benchRoutes(b, parseZeus, parseAPI)
}
