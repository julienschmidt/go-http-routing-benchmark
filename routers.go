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
	"github.com/Unknwon/macaron"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/bmizerany/pat"
	"github.com/daryl/zeus"
	"github.com/dimfeld/httptreemux"
	"github.com/emicklei/go-restful"
	"github.com/gin-gonic/gin"
	"github.com/go-martini/martini"
	"github.com/go-zoo/bone"
	"github.com/gocraft/web"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/lunny/tango"
	vulcan "github.com/mailgun/route"
	"github.com/naoina/denco"
	"github.com/naoina/kocha-urlrouter"
	_ "github.com/naoina/kocha-urlrouter/doublearray"
	"github.com/pilu/traffic"
	"github.com/plimble/ace"
	"github.com/rcrowley/go-tigertonic"
	"github.com/revel/revel"
	"github.com/robfig/pathtree"
	"github.com/typepress/rivet"
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
	// beego sets it to runtime.NumCPU()
	// Currently none of the contestors does concurrent routing
	runtime.GOMAXPROCS(1)

	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
	nullLogger = log.New(new(mockResponseWriter), "", 0)

	initBeego()
	initGin()
	initMartini()
	initRevel()
	initTango()
	initTraffic()
}

// Common
func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {}

// Ace
func aceHandle(_ *ace.C) {}

func aceHandleWrite(c *ace.C) {
	io.WriteString(c.Writer, c.Params.ByName("name"))
}

func loadAce(routes []route) http.Handler {
	router := ace.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, []ace.HandlerFunc{aceHandle})
	}
	return router
}

func loadAceSingle(method, path string, handle ace.HandlerFunc) http.Handler {
	router := ace.New()
	router.Handle(method, path, []ace.HandlerFunc{handle})
	return router
}

// beego
func beegoHandler(ctx *context.Context) {}

func beegoHandlerWrite(ctx *context.Context) {
	ctx.WriteString(ctx.Input.Param(":name"))
}

func initBeego() {
	beego.RunMode = "prod"
	beego.BeeLogger.Close()
}

func loadBeego(routes []route) http.Handler {
	re := regexp.MustCompile(":([^/]*)")
	app := beego.NewControllerRegister()
	for _, route := range routes {
		route.path = re.ReplaceAllString(route.path, ":$1")
		switch route.method {
		case "GET":
			app.Get(route.path, beegoHandler)
		case "POST":
			app.Post(route.path, beegoHandler)
		case "PUT":
			app.Put(route.path, beegoHandler)
		case "PATCH":
			app.Patch(route.path, beegoHandler)
		case "DELETE":
			app.Delete(route.path, beegoHandler)
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
func loadBone(routes []route) http.Handler {
	router := bone.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, http.HandlerFunc(httpHandlerFunc))
		case "POST":
			router.Post(route.path, http.HandlerFunc(httpHandlerFunc))
		case "PUT":
			router.Put(route.path, http.HandlerFunc(httpHandlerFunc))
		case "PATCH":
			router.Patch(route.path, http.HandlerFunc(httpHandlerFunc))
		case "DELETE":
			router.Delete(route.path, http.HandlerFunc(httpHandlerFunc))
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

// Denco
func dencoHandler(w http.ResponseWriter, r *http.Request, params denco.Params) {}

func dencoHandlerWrite(w http.ResponseWriter, r *http.Request, params denco.Params) {
	var name string
	for _, param := range params {
		if param.Name == "name" {
			name = param.Value
			break
		}
	}
	io.WriteString(w, name)
}

func loadDenco(routes []route) http.Handler {
	mux := denco.NewMux()
	handlers := make([]denco.Handler, 0, len(routes))
	for _, route := range routes {
		h := mux.Handler(route.method, route.path, dencoHandler)
		handlers = append(handlers, h)
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

func initGin() {
	gin.SetMode("release")
}

func loadGin(routes []route) http.Handler {
	router := gin.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, []gin.HandlerFunc{ginHandle})
	}
	return router
}

func loadGinSingle(method, path string, handle gin.HandlerFunc) http.Handler {
	router := gin.New()
	router.Handle(method, path, []gin.HandlerFunc{handle})
	return router
}

// gocraft/web
type gocraftWebContext struct{}

func gocraftWebHandler(w web.ResponseWriter, r *web.Request) {}

func gocraftWebHandlerWrite(w web.ResponseWriter, r *web.Request) {
	io.WriteString(w, r.PathParams["name"])
}

func loadGocraftWeb(routes []route) http.Handler {
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
	mux := goji.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			mux.Get(route.path, httpHandlerFunc)
		case "POST":
			mux.Post(route.path, httpHandlerFunc)
		case "PUT":
			mux.Put(route.path, httpHandlerFunc)
		case "PATCH":
			mux.Patch(route.path, httpHandlerFunc)
		case "DELETE":
			mux.Delete(route.path, httpHandlerFunc)
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

// go-json-rest/rest
func goJsonRestHandler(w rest.ResponseWriter, req *rest.Request) {}

func goJsonRestHandlerWrite(w rest.ResponseWriter, req *rest.Request) {
	io.WriteString(w.(io.Writer), req.PathParam("name"))
}

func newGoJsonRestResourceHandler() *rest.ResourceHandler {
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
		Logger:            nullLogger,
		ErrorLogger:       nullLogger,
		DisableXPoweredBy: true,
	}
	return &handler
}

func loadGoJsonRest(routes []route) http.Handler {
	handler := newGoJsonRestResourceHandler()
	restRoutes := make([]*rest.Route, 0, len(routes))
	for _, route := range routes {
		restRoutes = append(restRoutes,
			&rest.Route{route.method, route.path, goJsonRestHandler},
		)
	}
	handler.SetRoutes(restRoutes...)
	return handler
}

func loadGoJsonRestSingle(method, path string, hfunc rest.HandlerFunc) http.Handler {
	handler := newGoJsonRestResourceHandler()
	handler.SetRoutes(
		&rest.Route{method, path, hfunc},
	)
	return handler
}

// go-restful
func goRestfulHandlerWrite(r *restful.Request, w *restful.Response) {
	io.WriteString(w, r.Request.URL.Query().Get("name"))
}

func goRestfulHandler(r *restful.Request, w *restful.Response) {}

func loadGoRestful(routes []route) http.Handler {
	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)

	for _, route := range routes {
		switch route.method {
		case "GET":
			ws.Route(ws.GET(route.path).To(goRestfulHandler))
		case "POST":
			ws.Route(ws.POST(route.path).To(goRestfulHandler))
		case "PUT":
			ws.Route(ws.PUT(route.path).To(goRestfulHandler))
		case "PATCH":
			ws.Route(ws.PATCH(route.path).To(goRestfulHandler))
		case "DELETE":
			ws.Route(ws.DELETE(route.path).To(goRestfulHandler))
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
	re := regexp.MustCompile(":([^/]*)")
	m := mux.NewRouter()
	for _, route := range routes {
		m.HandleFunc(
			re.ReplaceAllString(route.path, "{$1}"),
			httpHandlerFunc,
		).Methods(route.method)
	}
	return m
}

func loadGorillaMuxSingle(method, path string, handler http.HandlerFunc) http.Handler {
	m := mux.NewRouter()
	m.HandleFunc(path, handler).Methods(method)
	return m
}

// HttpRouter
func httpRouterHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}

func httpRouterHandleWrite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	io.WriteString(w, ps.ByName("name"))
}

func loadHttpRouter(routes []route) http.Handler {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, httpRouterHandle)
	}
	return router
}

func loadHttpRouterSingle(method, path string, handle httprouter.Handle) http.Handler {
	router := httprouter.New()
	router.Handle(method, path, handle)
	return router
}

// httpTreeMux
func httpTreeMuxHandler(w http.ResponseWriter, r *http.Request, vars map[string]string) {}

func httpTreeMuxHandlerWrite(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	io.WriteString(w, vars["name"])
}

func loadHttpTreeMux(routes []route) http.Handler {
	router := httptreemux.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, httpTreeMuxHandler)
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

// Macaron
func macaronHandler(_ *macaron.Context) {}

func macaronHandlerWrite(c *macaron.Context) string {
	return c.Params("name")
}

func loadMacaron(routes []route) http.Handler {
	m := macaron.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			m.Get(route.path, martiniHandler)
		case "POST":
			m.Post(route.path, martiniHandler)
		case "PUT":
			m.Put(route.path, martiniHandler)
		case "PATCH":
			m.Patch(route.path, martiniHandler)
		case "DELETE":
			m.Delete(route.path, martiniHandler)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return m
}

func loadMacaronSingle(method, path string, handler interface{}) http.Handler {
	m := macaron.New()
	switch method {
	case "GET":
		m.Get(path, handler)
	case "POST":
		m.Post(path, handler)
	case "PUT":
		m.Put(path, handler)
	case "PATCH":
		m.Patch(path, handler)
	case "DELETE":
		m.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
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

// Revel (Router only)
// In the following code some Revel internals are modelled.
// The original revel code is copyrighted by Rob Figueiredo.
// See https://github.com/revel/revel/blob/master/LICENSE
type RevelController struct {
	*revel.Controller
	router *revel.Router
}

func (rc *RevelController) Handle() revel.Result {
	return revelResult{}
}

func (rc *RevelController) HandleWrite() revel.Result {
	return rc.RenderText(rc.Params.Get("name"))
}

type revelResult struct{}

func (rr revelResult) Apply(req *revel.Request, resp *revel.Response) {}

func (rc *RevelController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Dirty hacks, do NOT copy!
	revel.MainRouter = rc.router

	upgrade := r.Header.Get("Upgrade")
	if upgrade == "websocket" || upgrade == "Websocket" {
		panic("Not implemented")
	} else {
		var (
			req  = revel.NewRequest(r)
			resp = revel.NewResponse(w)
			c    = revel.NewController(req, resp)
		)
		req.Websocket = nil
		revel.Filters[0](c, revel.Filters[1:])
		if c.Result != nil {
			c.Result.Apply(req, resp)
		} else if c.Response.Status != 0 {
			panic("Not implemented")
		}
		// Close the Writer if we can
		if w, ok := resp.Out.(io.Closer); ok {
			w.Close()
		}
	}
}

func initRevel() {
	// Only use the Revel filters required for this benchmark
	revel.Filters = []revel.Filter{
		revel.RouterFilter,
		revel.ParamsFilter,
		revel.ActionInvoker,
	}

	revel.RegisterController((*RevelController)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Handle",
			},
			&revel.MethodType{
				Name: "HandleWrite",
			},
		})
}

func loadRevel(routes []route) http.Handler {
	router := revel.NewRouter("")

	// parseRoutes
	var rs []*revel.Route
	for _, r := range routes {
		rs = append(rs, revel.NewRoute(r.method, r.path, "RevelController.Handle", "", "", 0))
	}
	router.Routes = rs

	// updateTree
	router.Tree = pathtree.New()
	for _, r := range router.Routes {
		err := router.Tree.Add(r.TreePath, r)
		// Allow GETs to respond to HEAD requests.
		if err == nil && r.Method == "GET" {
			err = router.Tree.Add("/HEAD"+r.Path, r)
		}
		// Error adding a route to the pathtree.
		if err != nil {
			panic(err)
		}
	}

	rc := new(RevelController)
	rc.router = router
	return rc
}

func loadRevelSingle(method, path, action string) http.Handler {
	router := revel.NewRouter("")

	route := revel.NewRoute(method, path, action, "", "", 0)
	if err := router.Tree.Add(route.TreePath, route); err != nil {
		panic(err)
	}

	rc := new(RevelController)
	rc.router = router
	return rc
}

// Rivet
func rivetHandler() {}
func rivetHandlerWrite(c rivet.Context) {
	c.WriteString(c.GetParams().Get("name"))
}

func loadRivet(routes []route) http.Handler {
	router := rivet.NewRouter(nil)
	for _, route := range routes {
		router.Handle(route.method, route.path, rivetHandler)
	}
	return router
}

func loadRivetSingle(method, path string, handler interface{}) http.Handler {
	router := rivet.NewRouter(nil)

	router.Handle(method, path, handler)

	return router
}

// Tango
func tangoHandler(ctx *tango.Context) {
}
func tangoHandlerWrite(ctx *tango.Context) {
	ctx.Write([]byte(ctx.Params().Get(":name")))
}

func initTango() {
	tango.Env = tango.Prod
}

func loadTango(routes []route) http.Handler {
	tg := tango.New()
	for _, route := range routes {
		tg.Route([]string{route.method}, route.path, tangoHandler)
	}
	return tg
}

func loadTangoSingle(method, path string, handler func(*tango.Context)) http.Handler {
	tg := tango.New()
	tg.Route([]string{method}, path, handler)
	return tg
}

// Tiger Tonic
func tigerTonicHandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get("name"))
}

func loadTigerTonic(routes []route) http.Handler {
	re := regexp.MustCompile(":([^/]*)")
	mux := tigertonic.NewTrieServeMux()
	for _, route := range routes {
		mux.HandleFunc(route.method, re.ReplaceAllString(route.path, "{$1}"), httpHandlerFunc)
	}
	return mux
}

func loadTigerTonicSingle(method, path string, handler http.HandlerFunc) http.Handler {
	mux := tigertonic.NewTrieServeMux()
	mux.HandleFunc(method, path, handler)
	return mux
}

// Traffic
func trafficHandlerWrite(w traffic.ResponseWriter, r *traffic.Request) {
	io.WriteString(w, r.URL.Query().Get("name"))
}
func trafficHandler(w traffic.ResponseWriter, r *traffic.Request) {}

func initTraffic() {
	traffic.SetVar("env", "bench")
}

func loadTraffic(routes []route) http.Handler {
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
	re := regexp.MustCompile(":([^/]*)")
	mux := vulcan.NewMux()
	for _, route := range routes {
		path := re.ReplaceAllString(route.path, "<$1>")
		expr := fmt.Sprintf(`Method("%s") && Path("%s")`, route.method, path)
		if err := mux.HandleFunc(expr, httpHandlerFunc); err != nil {
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
func loadZeus(routes []route) http.Handler {
	m := zeus.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			m.GET(route.path, http.HandlerFunc(httpHandlerFunc))
		case "POST":
			m.POST(route.path, http.HandlerFunc(httpHandlerFunc))
		case "PUT":
			m.PUT(route.path, http.HandlerFunc(httpHandlerFunc))
		case "DELETE":
			m.DELETE(route.path, http.HandlerFunc(httpHandlerFunc))
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return m
}

func loadZeusSingle(method, path string, handler http.HandlerFunc) http.Handler {
	m := zeus.New()
	switch method {
	case "GET":
		m.GET(path, handler)
	case "POST":
		m.POST(path, handler)
	case "PUT":
		m.PUT(path, handler)
	case "DELETE":
		m.DELETE(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return m
}

// Usage notice
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}
