// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package benchmark

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"testing"

	"github.com/bmizerany/pat"
	"github.com/codegangsta/martini"
	"github.com/gocraft/web"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/pilu/traffic"
	"github.com/rcrowley/go-tigertonic"
	goji "github.com/zenazn/goji/web"
)

type route struct {
	method string
	path   string
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

func init() {
	log.SetOutput(new(mockResponseWriter))
}

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := new(mockResponseWriter)
	u := r.URL
	rq := u.RawQuery

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
			u.Path = route.path
			u.RawQuery = rq
			router.ServeHTTP(w, r)
		}
	}
}

// Common
func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {}

// gocraft/web
type gocraftWebContext struct{}

func gocraftWebHandler(w web.ResponseWriter, r *web.Request) {}

func gocraftWebHandlerWrite(w web.ResponseWriter, r *web.Request) {
	io.WriteString(w, r.PathParams["name"])
}

func loadGocraftWeb(routes []route) *web.Router {
	router := web.New(gocraftWebContext{})
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, gocraftWebHandler)
		case "POST":
			router.Post(route.path, gocraftWebHandler)
		case "PUT":
			router.Put(route.path, gocraftWebHandler)
		case "PATCH":
			router.Patch(route.path, gocraftWebHandler)
		case "DELETE":
			router.Delete(route.path, gocraftWebHandler)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return router
}

// goji
func gojiFuncWrite(c goji.C, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, c.URLParams["name"])
}

func loadGoji(routes []route) *goji.Mux {
	router := goji.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, httpHandlerFunc)
		case "POST":
			router.Post(route.path, httpHandlerFunc)
		case "PUT":
			router.Put(route.path, httpHandlerFunc)
		case "PATCH":
			router.Patch(route.path, httpHandlerFunc)
		case "DELETE":
			router.Delete(route.path, httpHandlerFunc)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return router
}

// gorilla/mux
func gorillaHandlerWrite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	io.WriteString(w, params["name"])
}

func loadGorillaMux(routes []route) *mux.Router {
	re := regexp.MustCompile(":([^/]*)")
	m := mux.NewRouter()
	for _, route := range routes {
		m.HandleFunc(re.ReplaceAllString(route.path, "{$1}"), httpHandlerFunc).Methods(route.method)
	}
	return m
}

// HttpRouter
func httpRouterHandle(w http.ResponseWriter, r *http.Request, _ map[string]string) {}

func httpRouterHandleWrite(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	io.WriteString(w, vars["name"])
}

func loadHttpRouter(routes []route) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, httpRouterHandle)
	}
	return router
}

// Martini
func martiniHandler() {}

func martiniHandlerWrite(params martini.Params) string {
	return params["name"]
}

func loadMartini(routes []route) *martini.Martini {
	router := martini.NewRouter()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, martiniHandler)
		case "POST":
			router.Post(route.path, martiniHandler)
		case "PUT":
			router.Put(route.path, martiniHandler)
		case "PATCH":
			router.Patch(route.path, martiniHandler)
		case "DELETE":
			router.Delete(route.path, martiniHandler)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	martini := martini.New()
	martini.Action(router.Handle)
	return martini
}

// pat
func patHandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get(":name"))
}

func loadPat(routes []route) *pat.PatternServeMux {
	m := pat.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			m.Get(route.path, http.HandlerFunc(httpHandlerFunc))
		case "POST":
			m.Post(route.path, http.HandlerFunc(httpHandlerFunc))
		case "PUT":
			m.Put(route.path, http.HandlerFunc(httpHandlerFunc))
		case "DELETE":
			m.Del(route.path, http.HandlerFunc(httpHandlerFunc))
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return m
}

// Tiger Tonic
func tigerTonicHandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get("name"))
}

func loadTigerTonic(routes []route) *tigertonic.TrieServeMux {
	re := regexp.MustCompile(":([^/]*)")
	mux := tigertonic.NewTrieServeMux()
	for _, route := range routes {
		mux.HandleFunc(route.method, re.ReplaceAllString(route.path, "{$1}"), httpHandlerFunc)
	}
	return mux
}

// Traffic
func trafficHandlerWrite(w traffic.ResponseWriter, r *traffic.Request) {
	io.WriteString(w, r.URL.Query().Get("name"))
}
func trafficHandler(w traffic.ResponseWriter, r *traffic.Request) {}

func loadTraffic(routes []route) *traffic.Router {
	traffic.SetVar("env", "bench")
	router := traffic.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, trafficHandler)
		case "POST":
			router.Post(route.path, trafficHandler)
		case "PUT":
			router.Put(route.path, trafficHandler)
		case "PATCH":
			router.Patch(route.path, trafficHandler)
		case "DELETE":
			router.Delete(route.path, trafficHandler)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return router
}

// Micro Benchmarks

// Route with Param (no write)
func BenchmarkGocraftWeb_Param(b *testing.B) {
	router := web.New(gocraftWebContext{})
	router.Get("/user/:name", gocraftWebHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param(b *testing.B) {
	router := goji.New()
	router.Get("/user/:name", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param(b *testing.B) {
	router := mux.NewRouter()
	router.HandleFunc("/user/{name}", httpHandlerFunc).Methods("GET")

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param(b *testing.B) {
	router := httprouter.New()
	router.GET("/user/:name", httpRouterHandle)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkMartini_Param(b *testing.B) {
	router := martini.NewRouter()
	router.Get("/user/:name", martiniHandler)
	martini := martini.New()
	martini.Action(router.Handle)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, martini, r)
}
func BenchmarkPat_Param(b *testing.B) {
	router := pat.New()
	router.Get("/user/:name", http.HandlerFunc(httpHandlerFunc))

	w := new(mockResponseWriter)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r, _ := http.NewRequest("GET", "/user/gordon", nil)
		router.ServeHTTP(w, r)
	}

	//benchRequest(b, router, r)
}
func BenchmarkTigerTonic_Param(b *testing.B) {
	router := tigertonic.NewTrieServeMux()
	router.HandleFunc("GET", "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_Param(b *testing.B) {
	traffic.SetVar("env", "bench")
	router := traffic.New()
	router.Get("/user/:name", trafficHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}

var twentyPat = "/:a/:b/:c/:d/:e/:f/:g/:h/:i/:j/:k/:l/:m/:n/:o/:p/:q/:r/:s/:t"
var twentyBrace = "/{a}/{b}/{c}/{d}/{e}/{f}/{g}/{h}/{i}/{j}/{k}/{l}/{m}/{n}/{o}/{p}/{q}/{r}/{s}/{t}"
var twentyRoute = "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t"

// Route with 20 Params (no write)
func BenchmarkGocraftWeb_Param20(b *testing.B) {
	router := web.New(gocraftWebContext{})
	router.Get(twentyPat, gocraftWebHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param20(b *testing.B) {
	router := goji.New()
	router.Get(twentyPat, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param20(b *testing.B) {
	router := mux.NewRouter()
	router.HandleFunc(twentyBrace, httpHandlerFunc).Methods("GET")

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param20(b *testing.B) {
	router := httprouter.New()
	router.GET(twentyPat, httpRouterHandle)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkMartini_Param20(b *testing.B) {
	router := martini.NewRouter()
	router.Get(twentyPat, martiniHandler)
	martini := martini.New()
	martini.Action(router.Handle)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, martini, r)
}
func BenchmarkPat_Param20(b *testing.B) {
	router := pat.New()
	router.Get(twentyPat, http.HandlerFunc(httpHandlerFunc))

	w := new(mockResponseWriter)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r, _ := http.NewRequest("GET", twentyRoute, nil)
		router.ServeHTTP(w, r)
	}

	//benchRequest(b, router, r)
}
func BenchmarkTigerTonic_Param20(b *testing.B) {
	router := tigertonic.NewTrieServeMux()
	router.HandleFunc("GET", twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_Param20(b *testing.B) {
	traffic.SetVar("env", "bench")
	router := traffic.New()
	router.Get(twentyPat, trafficHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}

// Route with Param and write
func BenchmarkGocraftWeb_ParamWrite(b *testing.B) {
	router := web.New(gocraftWebContext{})
	router.Get("/user/:name", gocraftWebHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_ParamWrite(b *testing.B) {
	router := goji.New()
	router.Get("/user/:name", gojiFuncWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_ParamWrite(b *testing.B) {
	router := mux.NewRouter()
	router.HandleFunc("/user/{name}", gorillaHandlerWrite).Methods("GET")

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_ParamWrite(b *testing.B) {
	router := httprouter.New()
	router.GET("/user/:name", httpRouterHandleWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkMartini_ParamWrite(b *testing.B) {
	router := martini.NewRouter()
	router.Get("/user/:name", martiniHandlerWrite)
	martini := martini.New()
	martini.Action(router.Handle)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, martini, r)
}
func BenchmarkPat_ParamWrite(b *testing.B) {
	router := pat.New()
	router.Get("/user/:name", http.HandlerFunc(patHandlerWrite))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTigerTonic_ParamWrite(b *testing.B) {
	router := tigertonic.NewTrieServeMux()
	router.Handle("GET", "/user/{name}", http.HandlerFunc(tigerTonicHandlerWrite))
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_ParamWrite(b *testing.B) {
	traffic.SetVar("env", "bench")
	router := traffic.New()
	router.Get("/user/:name", trafficHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
