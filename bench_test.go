// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

var benchRe *regexp.Regexp

func isTested(name string) bool {
	if benchRe == nil {
		// Get -test.bench flag value (not accessible via flag package)
		bench := ""
		for _, arg := range os.Args {
			if strings.HasPrefix(arg, "-test.bench=") {
				// ignore the benchmark name after an underscore
				bench = strings.SplitN(arg[12:], "_", 2)[0]
				break
			}
		}

		// Compile RegExp to match Benchmark names
		var err error
		benchRe, err = regexp.Compile(bench)
		if err != nil {
			panic(err.Error())
		}
	}
	return benchRe.MatchString(name)
}

func calcMem(name string, load func()) {
	if !isTested(name) {
		return
	}

	m := new(runtime.MemStats)

	// before
	runtime.GC()
	runtime.ReadMemStats(m)
	before := m.HeapAlloc

	load()

	// after
	runtime.GC()
	runtime.ReadMemStats(m)
	after := m.HeapAlloc
	println("   "+name+":", after-before, "Bytes")
}

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := new(mockResponseWriter)
	u := r.URL
	rq := u.RawQuery
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.RawQuery = rq
		router.ServeHTTP(w, r)
	}
}

func benchRoutes(b *testing.B, router http.Handler, routes []route) {
	w := new(mockResponseWriter)
	r, _ := http.NewRequest("GET", "/", nil)
	u := r.URL
	rq := u.RawQuery

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, route := range routes {
			r.Method = route.method
			r.RequestURI = route.path
			u.Path = route.path
			u.RawQuery = rq
			router.ServeHTTP(w, r)
		}
	}
}

// Micro Benchmarks
// Route with Param (no write)

func BenchmarkMy_Param(b *testing.B) {
	router := loadMySingle("GET", "/user/:name", httpHandlerFunc)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_Param(b *testing.B) {
	router := loadBeegoSingle("GET", "/user/:name", beegoHandler)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_Param(b *testing.B) {
	router := loadChiSingle("GET", "/user/:name", httpHandlerFunc)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_Param(b *testing.B) {
	router := loadDencoSingle("GET", "/user/:name", dencoHandler)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param(b *testing.B) {
	router := loadGinSingle("GET", "/user/:name", ginHandle)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param(b *testing.B) {
	router := loadHttpRouterSingle("GET", "/user/:name", httpRouterHandle)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkLARS_Param(b *testing.B) {
	router := loadLARSSingle("GET", "/user/:name", larsHandler)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkPossum_Param(b *testing.B) {
	router := loadPossumSingle("GET", "/user/:name", possumHandler)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkRivet_Param(b *testing.B) {
	router := loadRivetSingle("GET", "/user/:name", rivetHandler)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkVulcan_Param(b *testing.B) {
	router := loadVulcanSingle("GET", "/user/:name", vulcanHandler)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

// Route with 5 Params (no write)
const fiveColon = "/:a/:b/:c/:d/:e"
const fiveBrace = "/{a}/{b}/{c}/{d}/{e}"
const fiveRoute = "/test/test/test/test/test"

func BenchmarkMy_Param5(b *testing.B) {
	router := loadMySingle("GET", fiveColon, httpHandlerFunc)
	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_Param5(b *testing.B) {
	router := loadBeegoSingle("GET", fiveColon, beegoHandler)
	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_Param5(b *testing.B) {
	router := loadChiSingle("GET", fiveColon, httpHandlerFunc)
	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_Param5(b *testing.B) {
	router := loadDencoSingle("GET", fiveColon, dencoHandler)
	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param5(b *testing.B) {
	router := loadGinSingle("GET", fiveColon, ginHandle)
	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param5(b *testing.B) {
	router := loadHttpRouterSingle("GET", fiveColon, httpRouterHandle)
	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkLARS_Param5(b *testing.B) {
	router := loadLARSSingle("GET", fiveColon, larsHandler)
	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkPossum_Param5(b *testing.B) {
	router := loadPossumSingle("GET", fiveColon, possumHandler)
	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkRivet_Param5(b *testing.B) {
	router := loadRivetSingle("GET", fiveColon, rivetHandler)
	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkVulcan_Param5(b *testing.B) {
	router := loadVulcanSingle("GET", fiveColon, vulcanHandler)
	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}

// Route with 20 Params (no write)
const twentyColon = "/:a/:b/:c/:d/:e/:f/:g/:h/:i/:j/:k/:l/:m/:n/:o/:p/:q/:r/:s/:t"
const twentyBrace = "/{a}/{b}/{c}/{d}/{e}/{f}/{g}/{h}/{i}/{j}/{k}/{l}/{m}/{n}/{o}/{p}/{q}/{r}/{s}/{t}"
const twentyRoute = "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t"

func BenchmarkMy_Param20(b *testing.B) {
	router := loadMySingle("GET", twentyColon, httpHandlerFunc)
	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_Param20(b *testing.B) {
	router := loadBeegoSingle("GET", twentyColon, beegoHandler)
	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_Param20(b *testing.B) {
	router := loadChiSingle("GET", twentyColon, httpHandlerFunc)
	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_Param20(b *testing.B) {
	router := loadDencoSingle("GET", twentyColon, dencoHandler)
	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param20(b *testing.B) {
	router := loadGinSingle("GET", twentyColon, ginHandle)
	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param20(b *testing.B) {
	router := loadHttpRouterSingle("GET", twentyColon, httpRouterHandle)
	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkLARS_Param20(b *testing.B) {
	router := loadLARSSingle("GET", twentyColon, larsHandler)
	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkPossum_Param20(b *testing.B) {
	router := loadPossumSingle("GET", twentyColon, possumHandler)
	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkRivet_Param20(b *testing.B) {
	router := loadRivetSingle("GET", twentyColon, rivetHandler)
	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkVulcan_Param20(b *testing.B) {
	router := loadVulcanSingle("GET", twentyColon, vulcanHandler)
	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}

// Route with Param and write

func BenchmarkMy_ParamWrite(b *testing.B) {
	router := loadMySingle("GET", "/user/:name", myHandleWrite)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_ParamWrite(b *testing.B) {
	router := loadBeegoSingle("GET", "/user/:name", beegoHandlerWrite)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_ParamWrite(b *testing.B) {
	router := loadChiSingle("GET", "/user/:name", chiHandleWrite)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_ParamWrite(b *testing.B) {
	router := loadDencoSingle("GET", "/user/:name", dencoHandlerWrite)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_ParamWrite(b *testing.B) {
	router := loadGinSingle("GET", "/user/:name", ginHandleWrite)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_ParamWrite(b *testing.B) {
	router := loadHttpRouterSingle("GET", "/user/:name", httpRouterHandleWrite)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkLARS_ParamWrite(b *testing.B) {
	router := loadLARSSingle("GET", "/user/:name", larsHandlerWrite)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkPossum_ParamWrite(b *testing.B) {
	router := loadPossumSingle("GET", "/user/:name", possumHandlerWrite)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkRivet_ParamWrite(b *testing.B) {
	router := loadRivetSingle("GET", "/user/:name", rivetHandlerWrite)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkVulcan_ParamWrite(b *testing.B) {
	router := loadVulcanSingle("GET", "/user/:name", vulcanHandlerWrite)
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
