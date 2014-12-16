// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

var staticRoutes = []route{
	{"GET", "/"},
	{"GET", "/cmd.html"},
	{"GET", "/code.html"},
	{"GET", "/contrib.html"},
	{"GET", "/contribute.html"},
	{"GET", "/debugging_with_gdb.html"},
	{"GET", "/docs.html"},
	{"GET", "/effective_go.html"},
	{"GET", "/files.log"},
	{"GET", "/gccgo_contribute.html"},
	{"GET", "/gccgo_install.html"},
	{"GET", "/go-logo-black.png"},
	{"GET", "/go-logo-blue.png"},
	{"GET", "/go-logo-white.png"},
	{"GET", "/go1.1.html"},
	{"GET", "/go1.2.html"},
	{"GET", "/go1.html"},
	{"GET", "/go1compat.html"},
	{"GET", "/go_faq.html"},
	{"GET", "/go_mem.html"},
	{"GET", "/go_spec.html"},
	{"GET", "/help.html"},
	{"GET", "/ie.css"},
	{"GET", "/install-source.html"},
	{"GET", "/install.html"},
	{"GET", "/logo-153x55.png"},
	{"GET", "/Makefile"},
	{"GET", "/root.html"},
	{"GET", "/share.png"},
	{"GET", "/sieve.gif"},
	{"GET", "/tos.html"},
	{"GET", "/articles/"},
	{"GET", "/articles/go_command.html"},
	{"GET", "/articles/index.html"},
	{"GET", "/articles/wiki/"},
	{"GET", "/articles/wiki/edit.html"},
	{"GET", "/articles/wiki/final-noclosure.go"},
	{"GET", "/articles/wiki/final-noerror.go"},
	{"GET", "/articles/wiki/final-parsetemplate.go"},
	{"GET", "/articles/wiki/final-template.go"},
	{"GET", "/articles/wiki/final.go"},
	{"GET", "/articles/wiki/get.go"},
	{"GET", "/articles/wiki/http-sample.go"},
	{"GET", "/articles/wiki/index.html"},
	{"GET", "/articles/wiki/Makefile"},
	{"GET", "/articles/wiki/notemplate.go"},
	{"GET", "/articles/wiki/part1-noerror.go"},
	{"GET", "/articles/wiki/part1.go"},
	{"GET", "/articles/wiki/part2.go"},
	{"GET", "/articles/wiki/part3-errorhandling.go"},
	{"GET", "/articles/wiki/part3.go"},
	{"GET", "/articles/wiki/test.bash"},
	{"GET", "/articles/wiki/test_edit.good"},
	{"GET", "/articles/wiki/test_Test.txt.good"},
	{"GET", "/articles/wiki/test_view.good"},
	{"GET", "/articles/wiki/view.html"},
	{"GET", "/codewalk/"},
	{"GET", "/codewalk/codewalk.css"},
	{"GET", "/codewalk/codewalk.js"},
	{"GET", "/codewalk/codewalk.xml"},
	{"GET", "/codewalk/functions.xml"},
	{"GET", "/codewalk/markov.go"},
	{"GET", "/codewalk/markov.xml"},
	{"GET", "/codewalk/pig.go"},
	{"GET", "/codewalk/popout.png"},
	{"GET", "/codewalk/run"},
	{"GET", "/codewalk/sharemem.xml"},
	{"GET", "/codewalk/urlpoll.go"},
	{"GET", "/devel/"},
	{"GET", "/devel/release.html"},
	{"GET", "/devel/weekly.html"},
	{"GET", "/gopher/"},
	{"GET", "/gopher/appenginegopher.jpg"},
	{"GET", "/gopher/appenginegophercolor.jpg"},
	{"GET", "/gopher/appenginelogo.gif"},
	{"GET", "/gopher/bumper.png"},
	{"GET", "/gopher/bumper192x108.png"},
	{"GET", "/gopher/bumper320x180.png"},
	{"GET", "/gopher/bumper480x270.png"},
	{"GET", "/gopher/bumper640x360.png"},
	{"GET", "/gopher/doc.png"},
	{"GET", "/gopher/frontpage.png"},
	{"GET", "/gopher/gopherbw.png"},
	{"GET", "/gopher/gophercolor.png"},
	{"GET", "/gopher/gophercolor16x16.png"},
	{"GET", "/gopher/help.png"},
	{"GET", "/gopher/pkg.png"},
	{"GET", "/gopher/project.png"},
	{"GET", "/gopher/ref.png"},
	{"GET", "/gopher/run.png"},
	{"GET", "/gopher/talks.png"},
	{"GET", "/gopher/pencil/"},
	{"GET", "/gopher/pencil/gopherhat.jpg"},
	{"GET", "/gopher/pencil/gopherhelmet.jpg"},
	{"GET", "/gopher/pencil/gophermega.jpg"},
	{"GET", "/gopher/pencil/gopherrunning.jpg"},
	{"GET", "/gopher/pencil/gopherswim.jpg"},
	{"GET", "/gopher/pencil/gopherswrench.jpg"},
	{"GET", "/play/"},
	{"GET", "/play/fib.go"},
	{"GET", "/play/hello.go"},
	{"GET", "/play/life.go"},
	{"GET", "/play/peano.go"},
	{"GET", "/play/pi.go"},
	{"GET", "/play/sieve.go"},
	{"GET", "/play/solitaire.go"},
	{"GET", "/play/tree.go"},
	{"GET", "/progs/"},
	{"GET", "/progs/cgo1.go"},
	{"GET", "/progs/cgo2.go"},
	{"GET", "/progs/cgo3.go"},
	{"GET", "/progs/cgo4.go"},
	{"GET", "/progs/defer.go"},
	{"GET", "/progs/defer.out"},
	{"GET", "/progs/defer2.go"},
	{"GET", "/progs/defer2.out"},
	{"GET", "/progs/eff_bytesize.go"},
	{"GET", "/progs/eff_bytesize.out"},
	{"GET", "/progs/eff_qr.go"},
	{"GET", "/progs/eff_sequence.go"},
	{"GET", "/progs/eff_sequence.out"},
	{"GET", "/progs/eff_unused1.go"},
	{"GET", "/progs/eff_unused2.go"},
	{"GET", "/progs/error.go"},
	{"GET", "/progs/error2.go"},
	{"GET", "/progs/error3.go"},
	{"GET", "/progs/error4.go"},
	{"GET", "/progs/go1.go"},
	{"GET", "/progs/gobs1.go"},
	{"GET", "/progs/gobs2.go"},
	{"GET", "/progs/image_draw.go"},
	{"GET", "/progs/image_package1.go"},
	{"GET", "/progs/image_package1.out"},
	{"GET", "/progs/image_package2.go"},
	{"GET", "/progs/image_package2.out"},
	{"GET", "/progs/image_package3.go"},
	{"GET", "/progs/image_package3.out"},
	{"GET", "/progs/image_package4.go"},
	{"GET", "/progs/image_package4.out"},
	{"GET", "/progs/image_package5.go"},
	{"GET", "/progs/image_package5.out"},
	{"GET", "/progs/image_package6.go"},
	{"GET", "/progs/image_package6.out"},
	{"GET", "/progs/interface.go"},
	{"GET", "/progs/interface2.go"},
	{"GET", "/progs/interface2.out"},
	{"GET", "/progs/json1.go"},
	{"GET", "/progs/json2.go"},
	{"GET", "/progs/json2.out"},
	{"GET", "/progs/json3.go"},
	{"GET", "/progs/json4.go"},
	{"GET", "/progs/json5.go"},
	{"GET", "/progs/run"},
	{"GET", "/progs/slices.go"},
	{"GET", "/progs/timeout1.go"},
	{"GET", "/progs/timeout2.go"},
	{"GET", "/progs/update.bash"},
}

var (
	staticHttpServeMux http.Handler

	staticAce         http.Handler
	staticBeego       http.Handler
	staticBone        http.Handler
	staticDenco       http.Handler
	staticGin         http.Handler
	staticGocraftWeb  http.Handler
	staticGoji        http.Handler
	staticGoJsonRest  http.Handler
	staticGoRestful   http.Handler
	staticGorillaMux  http.Handler
	staticHttpRouter  http.Handler
	staticHttpTreeMux http.Handler
	staticKocha       http.Handler
	staticMacaron     http.Handler
	staticMartini     http.Handler
	staticPat         http.Handler
	staticRevel       http.Handler
	staticRivet       http.Handler
	staticTigerTonic  http.Handler
	staticTraffic     http.Handler
	staticZeus        http.Handler
)

func init() {
	println("#Static Routes:", len(staticRoutes))

	calcMem("HttpServeMux", func() {
		serveMux := http.NewServeMux()
		for _, route := range staticRoutes {
			serveMux.HandleFunc(route.path, httpHandlerFunc)
		}
		staticHttpServeMux = serveMux
	})

	calcMem("Ace", func() {
		staticAce = loadAce(staticRoutes)
	})
	calcMem("Beego", func() {
		staticBeego = loadBeego(staticRoutes)
	})
	calcMem("Bone", func() {
		staticBone = loadBone(staticRoutes)
	})
	calcMem("Denco", func() {
		staticDenco = loadDenco(staticRoutes)
	})
	calcMem("Gin", func() {
		staticGin = loadGin(staticRoutes)
	})
	calcMem("GocraftWeb", func() {
		staticGocraftWeb = loadGocraftWeb(staticRoutes)
	})
	calcMem("Goji", func() {
		staticGoji = loadGoji(staticRoutes)
	})
	calcMem("GoJsonRest", func() {
		staticGoJsonRest = loadGoJsonRest(staticRoutes)
	})
	calcMem("GoRestful", func() {
		staticGoRestful = loadGoRestful(staticRoutes)
	})
	calcMem("GorillaMux", func() {
		staticGorillaMux = loadGorillaMux(staticRoutes)
	})
	calcMem("HttpRouter", func() {
		staticHttpRouter = loadHttpRouter(staticRoutes)
	})
	calcMem("HttpTreeMux", func() {
		staticHttpTreeMux = loadHttpTreeMux(staticRoutes)
	})
	calcMem("Kocha", func() {
		staticKocha = loadKocha(staticRoutes)
	})
	calcMem("Macaron", func() {
		staticMacaron = loadMacaron(staticRoutes)
	})
	calcMem("Martini", func() {
		staticMartini = loadMartini(staticRoutes)
	})
	calcMem("Pat", func() {
		staticPat = loadPat(staticRoutes)
	})
	calcMem("Revel", func() {
		staticRevel = loadRevel(staticRoutes)
	})
	calcMem("Rivet", func() {
		staticRivet = loadRivet(staticRoutes)
	})
	calcMem("TigerTonic", func() {
		staticTigerTonic = loadTigerTonic(staticRoutes)
	})
	calcMem("Traffic", func() {
		staticTraffic = loadTraffic(staticRoutes)
	})
	calcMem("Zeus", func() {
		staticZeus = loadZeus(staticRoutes)
	})

	println()
}

// All routes

func BenchmarkAce_StaticAll(b *testing.B) {
	benchRoutes(b, staticAce, staticRoutes)
}
func BenchmarkHttpServeMux_StaticAll(b *testing.B) {
	benchRoutes(b, staticHttpServeMux, staticRoutes)
}
func BenchmarkBeego_StaticAll(b *testing.B) {
	benchRoutes(b, staticBeego, staticRoutes)
}
func BenchmarkBone_StaticAll(b *testing.B) {
	benchRoutes(b, staticBone, staticRoutes)
}
func BenchmarkDenco_StaticAll(b *testing.B) {
	benchRoutes(b, staticDenco, staticRoutes)
}
func BenchmarkGin_StaticAll(b *testing.B) {
	benchRoutes(b, staticGin, staticRoutes)
}
func BenchmarkGocraftWeb_StaticAll(b *testing.B) {
	benchRoutes(b, staticGocraftWeb, staticRoutes)
}
func BenchmarkGoji_StaticAll(b *testing.B) {
	benchRoutes(b, staticGoji, staticRoutes)
}
func BenchmarkGoJsonRest_StaticAll(b *testing.B) {
	benchRoutes(b, staticGoJsonRest, staticRoutes)
}
func BenchmarkGoRestful_StaticAll(b *testing.B) {
	benchRoutes(b, staticGoRestful, staticRoutes)
}
func BenchmarkGorillaMux_StaticAll(b *testing.B) {
	benchRoutes(b, staticGorillaMux, staticRoutes)
}
func BenchmarkHttpRouter_StaticAll(b *testing.B) {
	benchRoutes(b, staticHttpRouter, staticRoutes)
}
func BenchmarkHttpTreeMux_StaticAll(b *testing.B) {
	benchRoutes(b, staticHttpRouter, staticRoutes)
}
func BenchmarkKocha_StaticAll(b *testing.B) {
	benchRoutes(b, staticKocha, staticRoutes)
}
func BenchmarkMacaron_StaticAll(b *testing.B) {
	benchRoutes(b, staticMacaron, staticRoutes)
}
func BenchmarkMartini_StaticAll(b *testing.B) {
	benchRoutes(b, staticMartini, staticRoutes)
}
func BenchmarkPat_StaticAll(b *testing.B) {
	benchRoutes(b, staticPat, staticRoutes)
}
func BenchmarkRevel_StaticAll(b *testing.B) {
	benchRoutes(b, staticRevel, staticRoutes)
}
func BenchmarkRivet_StaticAll(b *testing.B) {
	benchRoutes(b, staticRivet, staticRoutes)
}
func BenchmarkTigerTonic_StaticAll(b *testing.B) {
	benchRoutes(b, staticTigerTonic, staticRoutes)
}
func BenchmarkTraffic_StaticAll(b *testing.B) {
	benchRoutes(b, staticTraffic, staticRoutes)
}
func BenchmarkZeus_StaticAll(b *testing.B) {
	benchRoutes(b, staticZeus, staticRoutes)
}
