// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"runtime"
	"testing"
)

func calcMem(name string, load func()) {
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
func BenchmarkBeego_Param(b *testing.B) {
	router := loadBeegoSingle("GET", "/user/:name", beegoHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkBone_Param(b *testing.B) {
	router := loadBoneSingle("GET", "/user/:name", http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_Param(b *testing.B) {
	handler := new(dencoHandler)
	router := loadDencoSingle(
		"GET", "/user/:name",
		handler, http.HandlerFunc(handler.Get),
	)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param(b *testing.B) {
	router := loadGinSingle("GET", "/user/:name", ginHandle)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_Param(b *testing.B) {
	router := loadGocraftWebSingle("GET", "/user/:name", gocraftWebHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param(b *testing.B) {
	router := loadGojiSingle("GET", "/user/:name", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoJsonRest_Param(b *testing.B) {
	router := loadGoJsonRestSingle("GET", "/user/:name", goJsonRestHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoRestful_Param(b *testing.B) {
	router := loadGoRestfulSingle("GET", "/user/:name", goRestfulHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param(b *testing.B) {
	router := loadGorillaMuxSingle("GET", "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param(b *testing.B) {
	router := loadHttpRouterSingle("GET", "/user/:name", httpRouterHandle)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", "/user/:name", httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_Param(b *testing.B) {
	handler := new(kochaHandler)
	router := loadKochaSingle(
		"GET", "/user/:name",
		handler, http.HandlerFunc(handler.Get),
	)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkMartini_Param(b *testing.B) {
	router := loadMartiniSingle("GET", "/user/:name", martiniHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_Param(b *testing.B) {
	router := loadPatSingle("GET", "/user/:name", http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkRevel_Param(b *testing.B) {
	router := loadRevelSingle("GET", "/user/:name", "RevelController.Handle")

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkRivet_Param(b *testing.B) {
	router := loadRivetSingle("GET", "/user/:name", rivetHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTango_Param(b *testing.B) {
	router := loadTangoSingle("GET", "/user/:name", tangoHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTigerTonic_Param(b *testing.B) {
	router := loadTigerTonicSingle("GET", "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_Param(b *testing.B) {
	router := loadTrafficSingle("GET", "/user/:name", trafficHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

// Route with 5 Params (no write)
const fiveColon = "/:a/:b/:c/:d/:e"
const fiveBrace = "/{a}/{b}/{c}/{d}/{e}"
const fiveRoute = "/test/test/test/test/test"

func BenchmarkBeego_Param5(b *testing.B) {
	router := loadBeegoSingle("GET", fiveColon, beegoHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkBone_Param5(b *testing.B) {
	router := loadBoneSingle("GET", fiveColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_Param5(b *testing.B) {
	handler := new(dencoHandler)
	router := loadDencoSingle(
		"GET", fiveColon,
		handler, http.HandlerFunc(handler.Get),
	)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param5(b *testing.B) {
	router := loadGinSingle("GET", fiveColon, ginHandle)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_Param5(b *testing.B) {
	router := loadGocraftWebSingle("GET", fiveColon, gocraftWebHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param5(b *testing.B) {
	router := loadGojiSingle("GET", fiveColon, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoJsonRest_Param5(b *testing.B) {
	handler := loadGoJsonRestSingle("GET", fiveColon, goJsonRestHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkGoRestful_Param5(b *testing.B) {
	router := loadGoRestfulSingle("GET", fiveColon, goRestfulHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param5(b *testing.B) {
	router := loadGorillaMuxSingle("GET", fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param5(b *testing.B) {
	router := loadHttpRouterSingle("GET", fiveColon, httpRouterHandle)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param5(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", fiveColon, httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_Param5(b *testing.B) {
	handler := new(kochaHandler)
	router := loadKochaSingle(
		"GET", fiveColon,
		handler, http.HandlerFunc(handler.Get),
	)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkMartini_Param5(b *testing.B) {
	router := loadMartiniSingle("GET", fiveColon, martiniHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_Param5(b *testing.B) {
	router := loadPatSingle("GET", fiveColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkRevel_Param5(b *testing.B) {
	router := loadRevelSingle("GET", fiveColon, "RevelController.Handle")

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkRivet_Param5(b *testing.B) {
	router := loadRivetSingle("GET", fiveColon, rivetHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkTango_Param5(b *testing.B) {
	router := loadTangoSingle("GET", fiveColon, tangoHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkTigerTonic_Param5(b *testing.B) {
	router := loadTigerTonicSingle("GET", fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_Param5(b *testing.B) {
	router := loadTrafficSingle("GET", fiveColon, trafficHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}

// Route with 20 Params (no write)
const twentyColon = "/:a/:b/:c/:d/:e/:f/:g/:h/:i/:j/:k/:l/:m/:n/:o/:p/:q/:r/:s/:t"
const twentyBrace = "/{a}/{b}/{c}/{d}/{e}/{f}/{g}/{h}/{i}/{j}/{k}/{l}/{m}/{n}/{o}/{p}/{q}/{r}/{s}/{t}"
const twentyRoute = "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t"

func BenchmarkBeego_Param20(b *testing.B) {
	router := loadBeegoSingle("GET", twentyColon, beegoHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkBone_Param20(b *testing.B) {
	router := loadBoneSingle("GET", twentyColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_Param20(b *testing.B) {
	handler := new(dencoHandler)
	router := loadDencoSingle(
		"GET", twentyColon,
		handler, http.HandlerFunc(handler.Get),
	)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param20(b *testing.B) {
	router := loadGinSingle("GET", twentyColon, ginHandle)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_Param20(b *testing.B) {
	router := loadGocraftWebSingle("GET", twentyColon, gocraftWebHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param20(b *testing.B) {
	router := loadGojiSingle("GET", twentyColon, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoJsonRest_Param20(b *testing.B) {
	handler := loadGoJsonRestSingle("GET", twentyColon, goJsonRestHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkGoRestful_Param20(b *testing.B) {
	handler := loadGoRestfulSingle("GET", twentyColon, goRestfulHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkGorillaMux_Param20(b *testing.B) {
	router := loadGorillaMuxSingle("GET", twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param20(b *testing.B) {
	router := loadHttpRouterSingle("GET", twentyColon, httpRouterHandle)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param20(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", twentyColon, httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_Param20(b *testing.B) {
	handler := new(kochaHandler)
	router := loadKochaSingle(
		"GET", twentyColon,
		handler, http.HandlerFunc(handler.Get),
	)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkMartini_Param20(b *testing.B) {
	router := loadMartiniSingle("GET", twentyColon, martiniHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_Param20(b *testing.B) {
	router := loadPatSingle("GET", twentyColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkRevel_Param20(b *testing.B) {
	router := loadRevelSingle("GET", twentyColon, "RevelController.Handle")

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkRivet_Param20(b *testing.B) {
	router := loadRivetSingle("GET", twentyColon, rivetHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkTango_Param20(b *testing.B) {
	router := loadTangoSingle("GET", twentyColon, tangoHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkTigerTonic_Param20(b *testing.B) {
	router := loadTigerTonicSingle("GET", twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_Param20(b *testing.B) {
	router := loadTrafficSingle("GET", twentyColon, trafficHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}

// Route with Param and write
func BenchmarkBeego_ParamWrite(b *testing.B) {
	router := loadBeegoSingle("GET", "/user/:name", beegoHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkBone_ParamWrite(b *testing.B) {
	router := loadBoneSingle("GET", "/user/:name", http.HandlerFunc(patHandlerWrite))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkDenco_ParamWrite(b *testing.B) {
	handler := new(dencoHandler)
	router := loadDencoSingle(
		"GET", "/user/:name",
		handler, http.HandlerFunc(handler.dencoHandlerWrite),
	)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_ParamWrite(b *testing.B) {
	router := loadGinSingle("GET", "/user/:name", ginHandleWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_ParamWrite(b *testing.B) {
	router := loadGocraftWebSingle("GET", "/user/:name", gocraftWebHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_ParamWrite(b *testing.B) {
	router := loadGojiSingle("GET", "/user/:name", gojiFuncWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoJsonRest_ParamWrite(b *testing.B) {
	handler := loadGoJsonRestSingle("GET", "/user/:name", goJsonRestHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, handler, r)
}
func BenchmarkGoRestful_ParamWrite(b *testing.B) {
	handler := loadGoRestfulSingle("GET", "/user/:name", goRestfulHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, handler, r)
}
func BenchmarkGorillaMux_ParamWrite(b *testing.B) {
	router := loadGorillaMuxSingle("GET", "/user/{name}", gorillaHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_ParamWrite(b *testing.B) {
	router := loadHttpRouterSingle("GET", "/user/:name", httpRouterHandleWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_ParamWrite(b *testing.B) {
	router := loadHttpTreeMuxSingle("GET", "/user/:name", httpTreeMuxHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_ParamWrite(b *testing.B) {
	handler := new(kochaHandler)
	router := loadKochaSingle(
		"GET", "/user/:name",
		handler, http.HandlerFunc(handler.kochaHandlerWrite),
	)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkMartini_ParamWrite(b *testing.B) {
	router := loadMartiniSingle("GET", "/user/:name", martiniHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_ParamWrite(b *testing.B) {
	router := loadPatSingle("GET", "/user/:name", http.HandlerFunc(patHandlerWrite))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkRevel_ParamWrite(b *testing.B) {
	router := loadRevelSingle("GET", "/user/:name", "RevelController.HandleWrite")

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkRivet_ParamWrite(b *testing.B) {
	router := loadRivetSingle("GET", "/user/:name", rivetHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTango_ParamWrite(b *testing.B) {
	router := loadTangoSingle("GET", "/user/:name", tangoHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTigerTonic_ParamWrite(b *testing.B) {
	router := loadTigerTonicSingle(
		"GET", "/user/{name}",
		http.HandlerFunc(tigerTonicHandlerWrite),
	)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_ParamWrite(b *testing.B) {
	router := loadTrafficSingle("GET", "/user/:name", trafficHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
