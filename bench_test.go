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

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/bmizerany/pat"
	"github.com/dimfeld/httptreemux"
	"github.com/go-martini/martini"
	"github.com/gocraft/web"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/naoina/kocha-urlrouter"
	_ "github.com/naoina/kocha-urlrouter/doublearray"
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

var nullLogger *log.Logger

func init() {
	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
	nullLogger = log.New(new(mockResponseWriter), "", 0)
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

// go-json-rest/rest
func goJsonRestHandler(w rest.ResponseWriter, req *rest.Request) {}

func goJsonRestHandlerWrite(w rest.ResponseWriter, req *rest.Request) {
	io.WriteString(w.(io.Writer), req.PathParam("name"))
}

func loadGoJsonRest(routes []route) *rest.ResourceHandler {
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
		Logger:            nullLogger,
		ErrorLogger:       nullLogger,
		DisableXPoweredBy: true,
	}
	restRoutes := make([]*rest.Route, 0, len(routes))
	for _, route := range routes {
		restRoutes = append(restRoutes,
			&rest.Route{route.method, route.path, goJsonRestHandler},
		)
	}
	handler.SetRoutes(restRoutes...)
	return &handler
}

func loadGoJsonRestSingle(method, path string, hfunc rest.HandlerFunc) *rest.ResourceHandler {
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
		Logger:            nullLogger,
		ErrorLogger:       nullLogger,
		DisableXPoweredBy: true,
	}
	handler.SetRoutes(
		&rest.Route{method, path, hfunc},
	)
	return &handler
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
func httpRouterHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}

func httpRouterHandleWrite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	io.WriteString(w, ps.ByName("name"))
}

func loadHttpRouter(routes []route) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, httpRouterHandle)
	}
	return router
}

// httpTreeMux
func httpTreeMuxHandler(w http.ResponseWriter, r *http.Request, vars map[string]string) {}

func httpTreeMuxHandlerWrite(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	io.WriteString(w, vars["name"])
}

func loadHttpTreeMux(routes []route) *httptreemux.TreeMux {
	router := httptreemux.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, httpTreeMuxHandler)
	}
	return router
}

// Kocha-urlrouter
type kochaHandler struct {
	routerMap map[string]urlrouter.URLRouter
	params    []urlrouter.Param
}

func (h *kochaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	meth, params := h.routerMap[r.Method].Lookup(r.URL.Path)
	h.params = params
	meth.(http.HandlerFunc).ServeHTTP(w, r)
}

func (h *kochaHandler) Get(w http.ResponseWriter, r *http.Request)    {}
func (h *kochaHandler) Post(w http.ResponseWriter, r *http.Request)   {}
func (h *kochaHandler) Put(w http.ResponseWriter, r *http.Request)    {}
func (h *kochaHandler) Patch(w http.ResponseWriter, r *http.Request)  {}
func (h *kochaHandler) Delete(w http.ResponseWriter, r *http.Request) {}
func (h *kochaHandler) kochaHandlerWrite(w http.ResponseWriter, r *http.Request) {
	var name string
	for _, param := range h.params {
		if param.Name == "name" {
			name = param.Value
			break
		}
	}
	io.WriteString(w, name)
}

func loadKocha(routes []route) *kochaHandler {
	handler := &kochaHandler{routerMap: map[string]urlrouter.URLRouter{
		"GET":    urlrouter.NewURLRouter("doublearray"),
		"POST":   urlrouter.NewURLRouter("doublearray"),
		"PUT":    urlrouter.NewURLRouter("doublearray"),
		"PATCH":  urlrouter.NewURLRouter("doublearray"),
		"DELETE": urlrouter.NewURLRouter("doublearray"),
	}}
	recordMap := make(map[string][]urlrouter.Record)
	for _, route := range routes {
		var f http.HandlerFunc
		switch route.method {
		case "GET":
			f = handler.Get
		case "POST":
			f = handler.Post
		case "PUT":
			f = handler.Put
		case "PATCH":
			f = handler.Patch
		case "DELETE":
			f = handler.Delete
		}
		recordMap[route.method] = append(recordMap[route.method], urlrouter.NewRecord(route.path, f))
	}
	for method, records := range recordMap {
		if err := handler.routerMap[method].Build(records); err != nil {
			panic(err)
		}
	}
	return handler
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
func BenchmarkGoJsonRest_Param(b *testing.B) {
	handler := loadGoJsonRestSingle("GET", "/user/:name", goJsonRestHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, handler, r)
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
func BenchmarkHttpTreeMux_Param(b *testing.B) {
	router := httptreemux.New()
	router.GET("/user/:name", httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_Param(b *testing.B) {
	handler := &kochaHandler{routerMap: map[string]urlrouter.URLRouter{
		"GET": urlrouter.NewURLRouter("doublearray"),
	}}
	if err := handler.routerMap["GET"].Build([]urlrouter.Record{
		urlrouter.NewRecord("/user/:name", http.HandlerFunc(handler.Get)),
	}); err != nil {
		panic(err)
	}
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, handler, r)
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

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
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

// Route with 5 Params (no write)
var fiveColon = "/:a/:b/:c/:d/:e"
var fiveBrace = "/{a}/{b}/{c}/{d}/{e}"
var fiveRoute = "/test/test/test/test/test"

func BenchmarkGocraftWeb_Param5(b *testing.B) {
	router := web.New(gocraftWebContext{})
	router.Get(fiveColon, gocraftWebHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param5(b *testing.B) {
	router := goji.New()
	router.Get(fiveColon, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoJsonRest_Param5(b *testing.B) {
	handler := loadGoJsonRestSingle("GET", fiveColon, goJsonRestHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkGorillaMux_Param5(b *testing.B) {
	router := mux.NewRouter()
	router.HandleFunc(fiveBrace, httpHandlerFunc).Methods("GET")

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param5(b *testing.B) {
	router := httprouter.New()
	router.GET(fiveColon, httpRouterHandle)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param5(b *testing.B) {
	router := httptreemux.New()
	router.GET(fiveColon, httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_Param5(b *testing.B) {
	handler := &kochaHandler{routerMap: map[string]urlrouter.URLRouter{
		"GET": urlrouter.NewURLRouter("doublearray"),
	}}
	if err := handler.routerMap["GET"].Build([]urlrouter.Record{
		urlrouter.NewRecord(fiveColon, http.HandlerFunc(handler.Get)),
	}); err != nil {
		panic(err)
	}
	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkMartini_Param5(b *testing.B) {
	router := martini.NewRouter()
	router.Get(fiveColon, martiniHandler)
	martini := martini.New()
	martini.Action(router.Handle)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, martini, r)
}
func BenchmarkPat_Param5(b *testing.B) {
	router := pat.New()
	router.Get(fiveColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTigerTonic_Param5(b *testing.B) {
	router := tigertonic.NewTrieServeMux()
	router.HandleFunc("GET", fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_Param5(b *testing.B) {
	traffic.SetVar("env", "bench")
	router := traffic.New()
	router.Get(fiveColon, trafficHandler)

	r, _ := http.NewRequest("GET", fiveRoute, nil)
	benchRequest(b, router, r)
}

// Route with 20 Params (no write)
var twentyColon = "/:a/:b/:c/:d/:e/:f/:g/:h/:i/:j/:k/:l/:m/:n/:o/:p/:q/:r/:s/:t"
var twentyBrace = "/{a}/{b}/{c}/{d}/{e}/{f}/{g}/{h}/{i}/{j}/{k}/{l}/{m}/{n}/{o}/{p}/{q}/{r}/{s}/{t}"
var twentyRoute = "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t"

func BenchmarkGocraftWeb_Param20(b *testing.B) {
	router := web.New(gocraftWebContext{})
	router.Get(twentyColon, gocraftWebHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param20(b *testing.B) {
	router := goji.New()
	router.Get(twentyColon, httpHandlerFunc)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoJsonRest_Param20(b *testing.B) {
	handler := loadGoJsonRestSingle("GET", twentyColon, goJsonRestHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkGorillaMux_Param20(b *testing.B) {
	router := mux.NewRouter()
	router.HandleFunc(twentyBrace, httpHandlerFunc).Methods("GET")

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param20(b *testing.B) {
	router := httprouter.New()
	router.GET(twentyColon, httpRouterHandle)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param20(b *testing.B) {
	router := httptreemux.New()
	router.GET(twentyColon, httpTreeMuxHandler)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_Param20(b *testing.B) {
	handler := &kochaHandler{routerMap: map[string]urlrouter.URLRouter{
		"GET": urlrouter.NewURLRouter("doublearray"),
	}}
	if err := handler.routerMap["GET"].Build([]urlrouter.Record{
		urlrouter.NewRecord(twentyColon, http.HandlerFunc(handler.Get)),
	}); err != nil {
		panic(err)
	}
	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkMartini_Param20(b *testing.B) {
	router := martini.NewRouter()
	router.Get(twentyColon, martiniHandler)
	martini := martini.New()
	martini.Action(router.Handle)

	r, _ := http.NewRequest("GET", twentyRoute, nil)
	benchRequest(b, martini, r)
}
func BenchmarkPat_Param20(b *testing.B) {
	router := pat.New()
	router.Get(twentyColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
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
	router.Get(twentyColon, trafficHandler)

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
func BenchmarkGoJsonRest_ParamWrite(b *testing.B) {
	handler := loadGoJsonRestSingle("GET", "/user/:name", goJsonRestHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, handler, r)
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
func BenchmarkHttpTreeMux_ParamWrite(b *testing.B) {
	router := httptreemux.New()
	router.GET("/user/:name", httpTreeMuxHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_ParamWrite(b *testing.B) {
	handler := &kochaHandler{routerMap: map[string]urlrouter.URLRouter{
		"GET": urlrouter.NewURLRouter("doublearray"),
	}}
	if err := handler.routerMap["GET"].Build([]urlrouter.Record{
		urlrouter.NewRecord("/user/:name", http.HandlerFunc(handler.kochaHandlerWrite)),
	}); err != nil {
		panic(err)
	}
	r, _ := http.NewRequest("GET", "/user/gordon", nil)
	benchRequest(b, handler, r)
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
