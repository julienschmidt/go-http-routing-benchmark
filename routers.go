// Copyright 2014 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"

	// If you add new routers please:
	// - Keep the benchmark functions etc. alphabetically sorted
	// - Make a pull request (without benchmark results) at
	//   https://github.com/julienschmidt/go-http-routing-benchmark

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/lars"
	"github.com/julienschmidt/httprouter"
	llog "github.com/lunny/log"
	"github.com/lunny/tango"
	vulcan "github.com/mailgun/route"
	"github.com/mikespook/possum"
	possumrouter "github.com/mikespook/possum/router"
	possumview "github.com/mikespook/possum/view"
	"github.com/naoina/denco"
	"github.com/nissy/bon"
	"github.com/pressly/chi"
	"github.com/typepress/rivet"
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

// flag indicating if the normal or the test handler should be loaded
var loadTestHandler = false

func init() {
	// beego sets it to runtime.NumCPU()
	// Currently none of the contesters does concurrent routing
	runtime.GOMAXPROCS(1)

	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
	nullLogger = log.New(new(mockResponseWriter), "", 0)

	initBeego()
	initGin()
	initTango()
}

// Common
func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {}

func httpHandlerFuncTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

// Bon
func myHandleWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, bon.URLParam(r, "name"))
}

func loadBon(routes []route) http.Handler {
	h := http.HandlerFunc(httpHandlerFunc)
	if loadTestHandler {
		h = http.HandlerFunc(httpHandlerFuncTest)
	}

	r := bon.NewRouter()
	for _, route := range routes {
		switch route.method {
		case "GET":
			r.Get(route.path, h)
		case "POST":
			r.Post(route.path, h)
		case "PUT":
			r.Put(route.path, h)
		case "PATCH":
			r.Patch(route.path, h)
		case "DELETE":
			r.Delete(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return r
}

func loadBonSingle(method, path string, h http.HandlerFunc) http.Handler {
	r := bon.NewRouter()
	switch method {
	case "GET":
		r.Get(path, h)
	case "POST":
		r.Post(path, h)
	case "PUT":
		r.Put(path, h)
	case "PATCH":
		r.Patch(path, h)
	case "DELETE":
		r.Delete(path, h)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return r
}

// beego
func beegoHandler(ctx *context.Context) {}

func beegoHandlerWrite(ctx *context.Context) {
	ctx.WriteString(ctx.Input.Param(":name"))
}

func beegoHandlerTest(ctx *context.Context) {
	ctx.WriteString(ctx.Request.RequestURI)
}

func initBeego() {
	beego.BConfig.RunMode = beego.PROD
	beego.BeeLogger.Close()
}

func loadBeego(routes []route) http.Handler {
	h := beegoHandler
	if loadTestHandler {
		h = beegoHandlerTest
	}

	re := regexp.MustCompile(":([^/]*)")
	app := beego.NewControllerRegister()
	for _, route := range routes {
		route.path = re.ReplaceAllString(route.path, ":$1")
		switch route.method {
		case "GET":
			app.Get(route.path, h)
		case "POST":
			app.Post(route.path, h)
		case "PUT":
			app.Put(route.path, h)
		case "PATCH":
			app.Patch(route.path, h)
		case "DELETE":
			app.Delete(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return app
}

func loadBeegoSingle(method, path string, handler beego.FilterFunc) http.Handler {
	app := beego.NewControllerRegister()
	switch method {
	case "GET":
		app.Get(path, handler)
	case "POST":
		app.Post(path, handler)
	case "PUT":
		app.Put(path, handler)
	case "PATCH":
		app.Patch(path, handler)
	case "DELETE":
		app.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return app
}

// Chi
func chiHandleWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, chi.URLParam(r, "name"))
}

func loadChi(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	mux := chi.NewRouter()
	for _, route := range routes {
		switch route.method {
		case "GET":
			mux.Get(route.path, h)
		case "POST":
			mux.Post(route.path, h)
		case "PUT":
			mux.Put(route.path, h)
		case "PATCH":
			mux.Patch(route.path, h)
		case "DELETE":
			mux.Delete(route.path, h)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return mux
}

func loadChiSingle(method, path string, handler http.HandlerFunc) http.Handler {
	mux := chi.NewRouter()
	switch method {
	case "GET":
		mux.Get(path, handler)
	case "POST":
		mux.Post(path, handler)
	case "PUT":
		mux.Put(path, handler)
	case "PATCH":
		mux.Patch(path, handler)
	case "DELETE":
		mux.Delete(path, handler)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return mux
}

// Denco
func dencoHandler(w http.ResponseWriter, r *http.Request, params denco.Params) {}

func dencoHandlerWrite(w http.ResponseWriter, r *http.Request, params denco.Params) {
	io.WriteString(w, params.Get("name"))
}

func dencoHandlerTest(w http.ResponseWriter, r *http.Request, params denco.Params) {
	io.WriteString(w, r.RequestURI)
}

func loadDenco(routes []route) http.Handler {
	h := dencoHandler
	if loadTestHandler {
		h = dencoHandlerTest
	}

	mux := denco.NewMux()
	handlers := make([]denco.Handler, 0, len(routes))
	for _, route := range routes {
		handler := mux.Handler(route.method, route.path, h)
		handlers = append(handlers, handler)
	}
	handler, err := mux.Build(handlers)
	if err != nil {
		panic(err)
	}
	return handler
}

func loadDencoSingle(method, path string, h denco.HandlerFunc) http.Handler {
	mux := denco.NewMux()
	handler, err := mux.Build([]denco.Handler{mux.Handler(method, path, h)})
	if err != nil {
		panic(err)
	}
	return handler
}

// Gin
func ginHandle(_ *gin.Context) {}

func ginHandleWrite(c *gin.Context) {
	io.WriteString(c.Writer, c.Params.ByName("name"))
}

func ginHandleTest(c *gin.Context) {
	io.WriteString(c.Writer, c.Request.RequestURI)
}

func initGin() {
	gin.SetMode(gin.ReleaseMode)
}

func loadGin(routes []route) http.Handler {
	h := ginHandle
	if loadTestHandler {
		h = ginHandleTest
	}

	router := gin.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadGinSingle(method, path string, handle gin.HandlerFunc) http.Handler {
	router := gin.New()
	router.Handle(method, path, handle)
	return router
}

// HttpRouter
func httpRouterHandle(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {}

func httpRouterHandleWrite(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	io.WriteString(w, ps.ByName("name"))
}

func httpRouterHandleTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, r.RequestURI)
}

func loadHttpRouter(routes []route) http.Handler {
	h := httpRouterHandle
	if loadTestHandler {
		h = httpRouterHandleTest
	}

	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadHttpRouterSingle(method, path string, handle httprouter.Handle) http.Handler {
	router := httprouter.New()
	router.Handle(method, path, handle)
	return router
}

// LARS
func larsHandler(c lars.Context) {
}

func larsHandlerWrite(c lars.Context) {
	io.WriteString(c.Response(), c.Param("name"))
}

func larsHandlerTest(c lars.Context) {
	io.WriteString(c.Response(), c.Request().RequestURI)
}

func loadLARS(routes []route) http.Handler {
	var h interface{} = larsHandler
	if loadTestHandler {
		h = larsHandlerTest
	}

	l := lars.New()

	for _, r := range routes {
		switch r.method {
		case "GET":
			l.Get(r.path, h)
		case "POST":
			l.Post(r.path, h)
		case "PUT":
			l.Put(r.path, h)
		case "PATCH":
			l.Patch(r.path, h)
		case "DELETE":
			l.Delete(r.path, h)
		default:
			panic("Unknow HTTP method: " + r.method)
		}
	}
	return l.Serve()
}

func loadLARSSingle(method, path string, h interface{}) http.Handler {
	l := lars.New()

	switch method {
	case "GET":
		l.Get(path, h)
	case "POST":
		l.Post(path, h)
	case "PUT":
		l.Put(path, h)
	case "PATCH":
		l.Patch(path, h)
	case "DELETE":
		l.Delete(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return l.Serve()
}

// Possum
func possumHandler(c *possum.Context) error {
	return nil
}

func possumHandlerWrite(c *possum.Context) error {
	io.WriteString(c.Response, c.Request.URL.Query().Get("name"))
	return nil
}

func possumHandlerTest(c *possum.Context) error {
	io.WriteString(c.Response, c.Request.RequestURI)
	return nil
}

func loadPossum(routes []route) http.Handler {
	h := possumHandler
	if loadTestHandler {
		h = possumHandlerTest
	}

	router := possum.NewServerMux()
	for _, route := range routes {
		router.HandleFunc(possumrouter.Simple(route.path), h, possumview.Simple("text/html", "utf-8"))
	}
	return router
}

func loadPossumSingle(method, path string, handler possum.HandlerFunc) http.Handler {
	router := possum.NewServerMux()
	router.HandleFunc(possumrouter.Simple(path), handler, possumview.Simple("text/html", "utf-8"))
	return router
}

// Rivet
func rivetHandler() {}

func rivetHandlerWrite(c rivet.Context) {
	c.WriteString(c.Get("name"))
}

func rivetHandlerTest(c rivet.Context) {
	c.WriteString(c.Req.RequestURI)
}

func loadRivet(routes []route) http.Handler {
	var h interface{} = rivetHandler
	if loadTestHandler {
		h = rivetHandlerTest
	}

	router := rivet.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadRivetSingle(method, path string, handler interface{}) http.Handler {
	router := rivet.New()

	router.Handle(method, path, handler)

	return router
}

// Tango
func tangoHandler(ctx *tango.Context) {}

func tangoHandlerWrite(ctx *tango.Context) {
	ctx.Write([]byte(ctx.Params().Get(":name")))
}

func tangoHandlerTest(ctx *tango.Context) {
	ctx.Write([]byte(ctx.Req().RequestURI))
}

func initTango() {
	llog.SetOutput(new(mockResponseWriter))
	llog.SetOutputLevel(llog.Lnone)
}

func loadTango(routes []route) http.Handler {
	h := tangoHandler
	if loadTestHandler {
		h = tangoHandlerTest
	}

	tg := tango.NewWithLog(llog.Std)
	for _, route := range routes {
		tg.Route(route.method, route.path, h)
	}
	return tg
}

// Mailgun Vulcan
func vulcanHandler(w http.ResponseWriter, r *http.Request) {}

func vulcanHandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get("name"))
}

func loadVulcan(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")
	mux := vulcan.NewMux()
	for _, route := range routes {
		path := re.ReplaceAllString(route.path, "<$1>")
		expr := fmt.Sprintf(`Method("%s") && Path("%s")`, route.method, path)
		if err := mux.HandleFunc(expr, h); err != nil {
			panic(err)
		}
	}
	return mux
}

func loadVulcanSingle(method, path string, handler http.HandlerFunc) http.Handler {
	re := regexp.MustCompile(":([^/]*)")
	mux := vulcan.NewMux()
	path = re.ReplaceAllString(path, "<$1>")
	expr := fmt.Sprintf(`Method("%s") && Path("%s")`, method, path)
	if err := mux.HandleFunc(expr, httpHandlerFunc); err != nil {
		panic(err)
	}
	return mux
}

// Usage notice
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}
