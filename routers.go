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
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/bmizerany/pat"
	"github.com/go-playground/lars"

	// "github.com/daryl/zeus"
	cloudykitrouter "github.com/cloudykit/router"
	"github.com/dimfeld/httptreemux"
	"github.com/emicklei/go-restful"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/go-martini/martini"
	"github.com/go-zoo/bone"
	"github.com/gocraft/web"
	"github.com/gorilla/mux"
	gowwwrouter "github.com/gowww/router"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo/v4"
	llog "github.com/lunny/log"
	"github.com/lunny/tango"
	vulcan "github.com/mailgun/route"
	"github.com/mikespook/possum"
	possumrouter "github.com/mikespook/possum/router"
	possumview "github.com/mikespook/possum/view"
	"github.com/naoina/denco"
	urlrouter "github.com/naoina/kocha-urlrouter"
	_ "github.com/naoina/kocha-urlrouter/doublearray"
	"github.com/pilu/traffic"
	"github.com/plimble/ace"
	"github.com/pratikdeoghare/lenrouter"
	"github.com/rcrowley/go-tigertonic"

	// "github.com/revel/pathtree"
	// "github.com/revel/revel"
	"github.com/aerogo/aero"
	"github.com/typepress/rivet"
	"github.com/ursiform/bear"
	"github.com/vanng822/r2router"
	goji "github.com/zenazn/goji/web"
	gojiv2 "goji.io"
	gojiv2pat "goji.io/pat"
	"gopkg.in/macaron.v1"
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
	initMartini()
	// initRevel()
	initTango()
	initTraffic()
}

// Common
func httpHandlerFunc(_ http.ResponseWriter, _ *http.Request) {}

func httpHandlerFuncTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

// Ace
func aceHandle(_ *ace.C) {}

func aceHandleWrite(c *ace.C) {
	io.WriteString(c.Writer, c.Param("name"))
}

func aceHandleTest(c *ace.C) {
	io.WriteString(c.Writer, c.Request.RequestURI)
}

func loadAce(routes []route) http.Handler {
	h := []ace.HandlerFunc{aceHandle}
	if loadTestHandler {
		h = []ace.HandlerFunc{aceHandleTest}
	}

	router := ace.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadAceSingle(method, path string, handle ace.HandlerFunc) http.Handler {
	router := ace.New()
	router.Handle(method, path, []ace.HandlerFunc{handle})
	return router
}

// Aero
func aeroHandler(c aero.Context) error {
	return nil
}

func aeroHandlerWrite(ctx aero.Context) error {
	io.WriteString(ctx.Response().Internal(), ctx.Get("name"))
	return nil
}
func aeroHandlerTest(ctx aero.Context) error {
	io.WriteString(ctx.Response().Internal(), ctx.Request().Path())
	return nil
}
func loadAero(routes []route) http.Handler {
	var h aero.Handler = aeroHandler
	if loadTestHandler {
		h = aeroHandlerTest
	}
	app := aero.New()
	for _, r := range routes {
		switch r.method {
		case "GET":
			app.Get(r.path, h)
		case "POST":
			app.Post(r.path, h)
		case "PUT":
			app.Put(r.path, h)
		case "PATCH":
			app.Router().Add(http.MethodPatch, r.path, h)
		case "DELETE":
			app.Delete(r.path, h)
		default:
			panic("Unknow HTTP method: " + r.method)
		}
	}
	return app
}
func loadAeroSingle(method, path string, h aero.Handler) http.Handler {
	app := aero.New()
	switch method {
	case "GET":
		app.Get(path, h)
	case "POST":
		app.Post(path, h)
	case "PUT":
		app.Put(path, h)
	case "PATCH":
		app.Router().Add(http.MethodPatch, path, h)
	case "DELETE":
		app.Delete(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	// }
	return app
}

// bear
func bearHandler(_ http.ResponseWriter, _ *http.Request, _ *bear.Context) {}

func bearHandlerWrite(w http.ResponseWriter, _ *http.Request, ctx *bear.Context) {
	io.WriteString(w, ctx.Params["name"])
}

func bearHandlerTest(w http.ResponseWriter, r *http.Request, _ *bear.Context) {
	io.WriteString(w, r.RequestURI)
}

func loadBear(routes []route) http.Handler {
	h := bearHandler
	if loadTestHandler {
		h = bearHandlerTest
	}

	router := bear.New()
	re := regexp.MustCompile(":([^/]*)")
	for _, route := range routes {
		switch route.method {
		case "GET", "POST", "PUT", "PATCH", "DELETE":
			router.On(route.method, re.ReplaceAllString(route.path, "{$1}"), h)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return router
}

func loadBearSingle(method string, path string, handler bear.HandlerFunc) http.Handler {
	router := bear.New()
	switch method {
	case "GET", "POST", "PUT", "PATCH", "DELETE":
		router.On(method, path, handler)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return router
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

// bone
func boneHandlerWrite(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, bone.GetValue(req, "name"))
}

func loadBone(routes []route) http.Handler {
	h := http.HandlerFunc(httpHandlerFunc)
	if loadTestHandler {
		h = http.HandlerFunc(httpHandlerFuncTest)
	}

	router := bone.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, h)
		case "POST":
			router.Post(route.path, h)
		case "PUT":
			router.Put(route.path, h)
		case "PATCH":
			router.Patch(route.path, h)
		case "DELETE":
			router.Delete(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return router
}

func loadBoneSingle(method, path string, handler http.Handler) http.Handler {
	router := bone.New()
	switch method {
	case "GET":
		router.Get(path, handler)
	case "POST":
		router.Post(path, handler)
	case "PUT":
		router.Put(path, handler)
	case "PATCH":
		router.Patch(path, handler)
	case "DELETE":
		router.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return router
}

// chi
// chi
func chiHandleWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, chi.URLParam(r, "name"))
}

func loadChi(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")

	mux := chi.NewRouter()
	for _, route := range routes {
		path := re.ReplaceAllString(route.path, "{$1}")

		switch route.method {
		case "GET":
			mux.Get(path, h)
		case "POST":
			mux.Post(path, h)
		case "PUT":
			mux.Put(path, h)
		case "PATCH":
			mux.Patch(path, h)
		case "DELETE":
			mux.Delete(path, h)
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

// CloudyKit Router
func cloudyKitRouterHandler(_ http.ResponseWriter, _ *http.Request, _ cloudykitrouter.Parameter) {}

func cloudyKitRouterHandlerWrite(w http.ResponseWriter, _ *http.Request, ps cloudykitrouter.Parameter) {
	io.WriteString(w, ps.ByName("name"))
}

func cloudyKitRouterHandlerTest(w http.ResponseWriter, r *http.Request, _ cloudykitrouter.Parameter) {
	io.WriteString(w, r.RequestURI)
}

func loadCloudyKitRouter(routes []route) http.Handler {
	h := cloudyKitRouterHandler
	if loadTestHandler {
		h = cloudyKitRouterHandlerTest
	}

	router := cloudykitrouter.New()
	for _, route := range routes {
		router.AddRoute(route.method, route.path, h)
	}
	return router
}

func loadCloudyKitRouterSingle(method, path string, handler cloudykitrouter.Handler) http.Handler {
	router := cloudykitrouter.New()
	router.AddRoute(method, path, handler)
	return router
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

// Echo
func echoHandler(c echo.Context) error {
	return nil
}

func echoHandlerWrite(c echo.Context) error {
	io.WriteString(c.Response(), c.Param("name"))
	return nil
}

func echoHandlerTest(c echo.Context) error {
	io.WriteString(c.Response(), c.Request().RequestURI)
	return nil
}

func loadEcho(routes []route) http.Handler {
	var h echo.HandlerFunc = echoHandler
	if loadTestHandler {
		h = echoHandlerTest
	}

	e := echo.New()
	for _, r := range routes {
		switch r.method {
		case "GET":
			e.GET(r.path, h)
		case "POST":
			e.POST(r.path, h)
		case "PUT":
			e.PUT(r.path, h)
		case "PATCH":
			e.PATCH(r.path, h)
		case "DELETE":
			e.DELETE(r.path, h)
		default:
			panic("Unknow HTTP method: " + r.method)
		}
	}
	return e
}

func loadEchoSingle(method, path string, h echo.HandlerFunc) http.Handler {
	e := echo.New()
	switch method {
	case "GET":
		e.GET(path, h)
	case "POST":
		e.POST(path, h)
	case "PUT":
		e.PUT(path, h)
	case "PATCH":
		e.PATCH(path, h)
	case "DELETE":
		e.DELETE(path, h)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return e
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

// gocraft/web
type gocraftWebContext struct{}

func gocraftWebHandler(w web.ResponseWriter, r *web.Request) {}

func gocraftWebHandlerWrite(w web.ResponseWriter, r *web.Request) {
	io.WriteString(w, r.PathParams["name"])
}

func gocraftWebHandlerTest(w web.ResponseWriter, r *web.Request) {
	io.WriteString(w, r.RequestURI)
}

func loadGocraftWeb(routes []route) http.Handler {
	h := gocraftWebHandler
	if loadTestHandler {
		h = gocraftWebHandlerTest
	}

	router := web.New(gocraftWebContext{})
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, h)
		case "POST":
			router.Post(route.path, h)
		case "PUT":
			router.Put(route.path, h)
		case "PATCH":
			router.Patch(route.path, h)
		case "DELETE":
			router.Delete(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return router
}

func loadGocraftWebSingle(method, path string, handler interface{}) http.Handler {
	router := web.New(gocraftWebContext{})
	switch method {
	case "GET":
		router.Get(path, handler)
	case "POST":
		router.Post(path, handler)
	case "PUT":
		router.Put(path, handler)
	case "PATCH":
		router.Patch(path, handler)
	case "DELETE":
		router.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return router
}

// goji
func gojiFuncWrite(c goji.C, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, c.URLParams["name"])
}

func loadGoji(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	mux := goji.New()
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

func loadGojiSingle(method, path string, handler interface{}) http.Handler {
	mux := goji.New()
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
		panic("Unknow HTTP method: " + method)
	}
	return mux
}

// goji v2 (github.com/goji/goji)
func gojiv2Handler(w http.ResponseWriter, r *http.Request) {}

func gojiv2HandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, gojiv2pat.Param(r, "name"))
}

func gojiv2HandlerTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

func loadGojiv2(routes []route) http.Handler {
	h := gojiv2Handler
	if loadTestHandler {
		h = gojiv2HandlerTest
	}

	mux := gojiv2.NewMux()
	for _, route := range routes {
		switch route.method {
		case "GET":
			mux.HandleFunc(gojiv2pat.Get(route.path), h)
		case "POST":
			mux.HandleFunc(gojiv2pat.Post(route.path), h)
		case "PUT":
			mux.HandleFunc(gojiv2pat.Put(route.path), h)
		case "PATCH":
			mux.HandleFunc(gojiv2pat.Patch(route.path), h)
		case "DELETE":
			mux.HandleFunc(gojiv2pat.Delete(route.path), h)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return mux
}

func loadGojiv2Single(method, path string, handler func(http.ResponseWriter, *http.Request)) http.Handler {
	mux := gojiv2.NewMux()
	switch method {
	case "GET":
		mux.HandleFunc(gojiv2pat.Get(path), handler)
	case "POST":
		mux.HandleFunc(gojiv2pat.Post(path), handler)
	case "PUT":
		mux.HandleFunc(gojiv2pat.Put(path), handler)
	case "PATCH":
		mux.HandleFunc(gojiv2pat.Patch(path), handler)
	case "DELETE":
		mux.HandleFunc(gojiv2pat.Delete(path), handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return mux
}

// go-json-rest/rest
func goJsonRestHandler(w rest.ResponseWriter, req *rest.Request) {}

func goJsonRestHandlerWrite(w rest.ResponseWriter, req *rest.Request) {
	io.WriteString(w.(io.Writer), req.PathParam("name"))
}

func goJsonRestHandlerTest(w rest.ResponseWriter, req *rest.Request) {
	io.WriteString(w.(io.Writer), req.RequestURI)
}

func loadGoJsonRest(routes []route) http.Handler {
	h := goJsonRestHandler
	if loadTestHandler {
		h = goJsonRestHandlerTest
	}

	api := rest.NewApi()
	restRoutes := make([]*rest.Route, 0, len(routes))
	for _, route := range routes {
		restRoutes = append(restRoutes,
			&rest.Route{route.method, route.path, h},
		)
	}
	router, err := rest.MakeRouter(restRoutes...)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	return api.MakeHandler()
}

func loadGoJsonRestSingle(method, path string, hfunc rest.HandlerFunc) http.Handler {
	api := rest.NewApi()
	router, err := rest.MakeRouter(
		&rest.Route{method, path, hfunc},
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	return api.MakeHandler()
}

// go-restful
func goRestfulHandler(r *restful.Request, w *restful.Response) {}

func goRestfulHandlerWrite(r *restful.Request, w *restful.Response) {
	io.WriteString(w, r.PathParameter("name"))
}

func goRestfulHandlerTest(r *restful.Request, w *restful.Response) {
	io.WriteString(w, r.Request.RequestURI)
}

func loadGoRestful(routes []route) http.Handler {
	h := goRestfulHandler
	if loadTestHandler {
		h = goRestfulHandlerTest
	}

	re := regexp.MustCompile(":([^/]*)")

	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)

	for _, route := range routes {
		path := re.ReplaceAllString(route.path, "{$1}")

		switch route.method {
		case "GET":
			ws.Route(ws.GET(path).To(h))
		case "POST":
			ws.Route(ws.POST(path).To(h))
		case "PUT":
			ws.Route(ws.PUT(path).To(h))
		case "PATCH":
			ws.Route(ws.PATCH(path).To(h))
		case "DELETE":
			ws.Route(ws.DELETE(path).To(h))
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	wsContainer.Add(ws)
	return wsContainer
}

func loadGoRestfulSingle(method, path string, handler restful.RouteFunction) http.Handler {
	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)
	switch method {
	case "GET":
		ws.Route(ws.GET(path).To(handler))
	case "POST":
		ws.Route(ws.POST(path).To(handler))
	case "PUT":
		ws.Route(ws.PUT(path).To(handler))
	case "PATCH":
		ws.Route(ws.PATCH(path).To(handler))
	case "DELETE":
		ws.Route(ws.DELETE(path).To(handler))
	default:
		panic("Unknow HTTP method: " + method)
	}
	wsContainer.Add(ws)
	return wsContainer
}

// gorilla/mux
func gorillaHandlerWrite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	io.WriteString(w, params["name"])
}

func loadGorillaMux(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")
	m := mux.NewRouter()
	for _, route := range routes {
		m.HandleFunc(
			re.ReplaceAllString(route.path, "{$1}"),
			h,
		).Methods(route.method)
	}
	return m
}

func loadGorillaMuxSingle(method, path string, handler http.HandlerFunc) http.Handler {
	m := mux.NewRouter()
	m.HandleFunc(path, handler).Methods(method)
	return m
}

// gowww/router
func gowwwRouterHandleWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, gowwwrouter.Parameter(r, "name"))
}

func loadGowwwRouter(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	router := gowwwrouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, http.HandlerFunc(h))
	}
	return router
}

func loadGowwwRouterSingle(method, path string, handler http.Handler) http.Handler {
	router := gowwwrouter.New()
	router.Handle(method, path, handler)
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

// httpTreeMux
func httpTreeMuxHandler(_ http.ResponseWriter, _ *http.Request, _ map[string]string) {}

func httpTreeMuxHandlerWrite(w http.ResponseWriter, _ *http.Request, vars map[string]string) {
	io.WriteString(w, vars["name"])
}

func httpTreeMuxHandlerTest(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	io.WriteString(w, r.RequestURI)
}

func loadHttpTreeMux(routes []route) http.Handler {
	h := httpTreeMuxHandler
	if loadTestHandler {
		h = httpTreeMuxHandlerTest
	}

	router := httptreemux.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, h)
	}
	return router
}

func loadHttpTreeMuxSingle(method, path string, handler httptreemux.HandlerFunc) http.Handler {
	router := httptreemux.New()
	router.Handle(method, path, handler)
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

func loadKocha(routes []route) http.Handler {
	/*h := httpRouterHandle
	if loadTestHandler {
		h = httpRouterHandleTest
	}*/

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
		recordMap[route.method] = append(
			recordMap[route.method],
			urlrouter.NewRecord(route.path, f),
		)
	}
	for method, records := range recordMap {
		if err := handler.routerMap[method].Build(records); err != nil {
			panic(err)
		}
	}
	return handler
}

func loadKochaSingle(method, path string, handler *kochaHandler, hfunc http.HandlerFunc) http.Handler {
	handler.routerMap = map[string]urlrouter.URLRouter{
		method: urlrouter.NewURLRouter("doublearray"),
	}

	if err := handler.routerMap[method].Build([]urlrouter.Record{
		urlrouter.NewRecord(path, hfunc),
	}); err != nil {
		panic(err)
	}
	return handler
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

func larsNativeHandlerTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
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

// lenrouter
func lenrouterHandle(_ http.ResponseWriter, _ *http.Request, _ lenrouter.Params) {}

func lenrouterHandleWrite(w http.ResponseWriter, _ *http.Request, ps lenrouter.Params) {
	io.WriteString(w, ps.ByName("name"))
}

func lenrouterHandleTest(w http.ResponseWriter, r *http.Request, _ lenrouter.Params) {
	io.WriteString(w, r.RequestURI)
}

func loadLenrouter(routes []route) http.Handler {
	h := lenrouterHandle
	if loadTestHandler {
		h = lenrouterHandleTest
	}
	var endpoints []lenrouter.Endpoint
	for _, route := range routes {
		endpoints = append(endpoints, lenrouter.Endpoint{
			Method:  route.method,
			Pattern: route.path,
			Handler: h,
		})
	}

	return lenrouter.New(200, 20, endpoints...)
}

func loadLenrouterSingle(method, path string, handle lenrouter.Handle) http.Handler {
	return lenrouter.New(200, 20, lenrouter.Endpoint{
		Method:  method,
		Pattern: path,
		Handler: handle,
	})
}

// Macaron
func macaronHandler() {}

func macaronHandlerWrite(c *macaron.Context) string {
	return c.Params("name")
}

func macaronHandlerTest(c *macaron.Context) string {
	return c.Req.RequestURI
}

func loadMacaron(routes []route) http.Handler {
	var h = []macaron.Handler{macaronHandler}
	if loadTestHandler {
		h[0] = macaronHandlerTest
	}

	m := macaron.New()
	for _, route := range routes {
		m.Handle(route.method, route.path, h)
	}
	return m
}

func loadMacaronSingle(method, path string, handler interface{}) http.Handler {
	m := macaron.New()
	m.Handle(method, path, []macaron.Handler{handler})
	return m
}

// Martini
func martiniHandler() {}

func martiniHandlerWrite(params martini.Params) string {
	return params["name"]
}

func initMartini() {
	martini.Env = martini.Prod
}

func loadMartini(routes []route) http.Handler {
	var h interface{} = martiniHandler
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	router := martini.NewRouter()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, h)
		case "POST":
			router.Post(route.path, h)
		case "PUT":
			router.Put(route.path, h)
		case "PATCH":
			router.Patch(route.path, h)
		case "DELETE":
			router.Delete(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	martini := martini.New()
	martini.Action(router.Handle)
	return martini
}

func loadMartiniSingle(method, path string, handler interface{}) http.Handler {
	router := martini.NewRouter()
	switch method {
	case "GET":
		router.Get(path, handler)
	case "POST":
		router.Post(path, handler)
	case "PUT":
		router.Put(path, handler)
	case "PATCH":
		router.Patch(path, handler)
	case "DELETE":
		router.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}

	martini := martini.New()
	martini.Action(router.Handle)
	return martini
}

// pat
func patHandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get(":name"))
}

func loadPat(routes []route) http.Handler {
	h := http.HandlerFunc(httpHandlerFunc)
	if loadTestHandler {
		h = http.HandlerFunc(httpHandlerFuncTest)
	}

	m := pat.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			m.Get(route.path, h)
		case "POST":
			m.Post(route.path, h)
		case "PUT":
			m.Put(route.path, h)
		case "DELETE":
			m.Del(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return m
}

func loadPatSingle(method, path string, handler http.Handler) http.Handler {
	m := pat.New()
	switch method {
	case "GET":
		m.Get(path, handler)
	case "POST":
		m.Post(path, handler)
	case "PUT":
		m.Put(path, handler)
	case "DELETE":
		m.Del(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return m
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

// R2router
func r2routerHandler(w http.ResponseWriter, req *http.Request, _ r2router.Params) {}

func r2routerHandleWrite(w http.ResponseWriter, req *http.Request, params r2router.Params) {
	io.WriteString(w, params.Get("name"))
}

func r2routerHandleTest(w http.ResponseWriter, req *http.Request, _ r2router.Params) {
	io.WriteString(w, req.RequestURI)
}

func loadR2router(routes []route) http.Handler {
	h := r2routerHandler
	if loadTestHandler {
		h = r2routerHandleTest
	}

	router := r2router.NewRouter()
	for _, r := range routes {
		router.AddHandler(r.method, r.path, h)
	}
	return router
}

func loadR2routerSingle(method, path string, handler r2router.HandlerFunc) http.Handler {
	router := r2router.NewRouter()
	router.AddHandler(method, path, handler)
	return router
}

// Revel (Router only)
// In the following code some Revel internals are modeled.
// The original revel code is copyrighted by Rob Figueiredo.
// See https://github.com/revel/revel/blob/master/LICENSE
// type RevelController struct {
// 	*revel.Controller
// 	router *revel.Router
// }

// func (rc *RevelController) Handle() revel.Result {
// 	return revelResult{}
// }

// func (rc *RevelController) HandleWrite() revel.Result {
// 	return rc.RenderText(rc.Params.Get("name"))
// }

// func (rc *RevelController) HandleTest() revel.Result {
// 	return rc.RenderText(rc.Request.GetRequestURI())
// }

// type revelResult struct{}

// func (rr revelResult) Apply(req *revel.Request, resp *revel.Response) {}

// func (rc *RevelController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// Dirty hacks, do NOT copy!
// 	revel.MainRouter = rc.router

// 	upgrade := r.Header.Get("Upgrade")
// 	if upgrade == "websocket" || upgrade == "Websocket" {
// 		panic("Not implemented")
// 	} else {
// 		var (
// 			req  = revel.NewRequest(r)
// 			resp = revel.NewResponse(w)
// 			c    = revel.NewController(req, resp)
// 		)
// 		req.Websocket = nil
// 		revel.Filters[0](c, revel.Filters[1:])
// 		if c.Result != nil {
// 			c.Result.Apply(req, resp)
// 		} else if c.Response.Status != 0 {
// 			panic("Not implemented")
// 		}
// 		// Close the Writer if we can
// 		if w, ok := resp.Out.(io.Closer); ok {
// 			w.Close()
// 		}
// 	}
// }

// func initRevel() {
// 	// Only use the Revel filters required for this benchmark
// 	revel.Filters = []revel.Filter{
// 		revel.RouterFilter,
// 		revel.ParamsFilter,
// 		revel.ActionInvoker,
// 	}

// 	revel.RegisterController((*RevelController)(nil),
// 		[]*revel.MethodType{
// 			{
// 				Name: "Handle",
// 			},
// 			{
// 				Name: "HandleWrite",
// 			},
// 			{
// 				Name: "HandleTest",
// 			},
// 		})
// }

// func loadRevel(routes []route) http.Handler {
// 	h := "RevelController.Handle"
// 	if loadTestHandler {
// 		h = "RevelController.HandleTest"
// 	}

// 	router := revel.NewRouter("")

// 	// parseRoutes
// 	var rs []*revel.Route
// 	for _, r := range routes {
// 		rs = append(rs, revel.NewRoute(r.method, r.path, h, "", "", 0))
// 	}
// 	router.Routes = rs

// 	// updateTree
// 	router.Tree = pathtree.New()
// 	for _, r := range router.Routes {
// 		err := router.Tree.Add(r.TreePath, r)
// 		// Allow GETs to respond to HEAD requests.
// 		if err == nil && r.Method == "GET" {
// 			err = router.Tree.Add("/HEAD"+r.Path, r)
// 		}
// 		// Error adding a route to the pathtree.
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	rc := new(RevelController)
// 	rc.router = router
// 	return rc
// }

// func loadRevelSingle(method, path, action string) http.Handler {
// 	router := revel.NewRouter("")

// 	route := revel.NewRoute(method, path, action, "", "", 0)
// 	if err := router.Tree.Add(route.TreePath, route); err != nil {
// 		panic(err)
// 	}

// 	rc := new(RevelController)
// 	rc.router = router
// 	return rc
// }

// Rivet
func rivetHandler() {}

func rivetHandlerWrite(c *rivet.Context) {
	c.WriteString(c.Get("name"))
}

func rivetHandlerTest(c *rivet.Context) {
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

func loadTangoSingle(method, path string, handler func(*tango.Context)) http.Handler {
	tg := tango.NewWithLog(llog.Std)
	tg.Route(method, path, handler)
	return tg
}

// Tiger Tonic
func tigerTonicHandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get("name"))
}

func loadTigerTonic(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")
	mux := tigertonic.NewTrieServeMux()
	for _, route := range routes {
		mux.HandleFunc(route.method, re.ReplaceAllString(route.path, "{$1}"), h)
	}
	return mux
}

func loadTigerTonicSingle(method, path string, handler http.HandlerFunc) http.Handler {
	mux := tigertonic.NewTrieServeMux()
	mux.HandleFunc(method, path, handler)
	return mux
}

// Traffic
func trafficHandler(w traffic.ResponseWriter, r *traffic.Request) {}

func trafficHandlerWrite(w traffic.ResponseWriter, r *traffic.Request) {
	io.WriteString(w, r.URL.Query().Get("name"))
}

func trafficHandlerTest(w traffic.ResponseWriter, r *traffic.Request) {
	io.WriteString(w, r.RequestURI)
}

func initTraffic() {
	traffic.SetVar("env", "bench")
}

func loadTraffic(routes []route) http.Handler {
	h := trafficHandler
	if loadTestHandler {
		h = trafficHandlerTest
	}

	router := traffic.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, h)
		case "POST":
			router.Post(route.path, h)
		case "PUT":
			router.Put(route.path, h)
		case "PATCH":
			router.Patch(route.path, h)
		case "DELETE":
			router.Delete(route.path, h)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return router
}

func loadTrafficSingle(method, path string, handler traffic.HttpHandleFunc) http.Handler {
	router := traffic.New()
	switch method {
	case "GET":
		router.Get(path, handler)
	case "POST":
		router.Post(path, handler)
	case "PUT":
		router.Put(path, handler)
	case "PATCH":
		router.Patch(path, handler)
	case "DELETE":
		router.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return router
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

// Zeus
// func zeusHandlerWrite(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, zeus.Var(r, "name"))
// }

// func loadZeus(routes []route) http.Handler {
// 	h := http.HandlerFunc(httpHandlerFunc)
// 	if loadTestHandler {
// 		h = http.HandlerFunc(httpHandlerFuncTest)
// 	}

// 	m := zeus.New()
// 	for _, route := range routes {
// 		switch route.method {
// 		case "GET":
// 			m.GET(route.path, h)
// 		case "POST":
// 			m.POST(route.path, h)
// 		case "PUT":
// 			m.PUT(route.path, h)
// 		case "DELETE":
// 			m.DELETE(route.path, h)
// 		default:
// 			panic("Unknow HTTP method: " + route.method)
// 		}
// 	}
// 	return m
// }

// func loadZeusSingle(method, path string, handler http.HandlerFunc) http.Handler {
// 	m := zeus.New()
// 	switch method {
// 	case "GET":
// 		m.GET(path, handler)
// 	case "POST":
// 		m.POST(path, handler)
// 	case "PUT":
// 		m.PUT(path, handler)
// 	case "DELETE":
// 		m.DELETE(path, handler)
// 	default:
// 		panic("Unknow HTTP method: " + method)
// 	}
// 	return m
// }

// Usage notice
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}
