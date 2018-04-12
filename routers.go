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
	"runtime"

	// If you add new routers please:
	// - Keep the benchmark functions etc. alphabetically sorted
	// - Make a pull request (without benchmark results) at
	//   https://github.com/julienschmidt/go-http-routing-benchmark

	"github.com/go-chi/chi"
	"github.com/nissy/bon"
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

// flag indicating if the normal or the test handler should be loaded
var loadTestHandler = false

func init() {
	// beego sets it to runtime.NumCPU()
	// Currently none of the contesters does concurrent routing
	runtime.GOMAXPROCS(1)

	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
}

// Common
func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {
}

func httpHandlerFuncTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

// Bon
func bonHandleWrite(w http.ResponseWriter, r *http.Request) {
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

// Chi
func chiHandleWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, chi.URLParam(r, "name"))
}

func loadChi(routes []route) http.Handler {
	h := httpHandlerFunc

	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	r := chi.NewRouter()

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
			panic("Unknown HTTP method: " + route.method)
		}
	}

	return r
}

func loadChiSingle(method, path string, handler http.HandlerFunc) http.Handler {
	r := chi.NewRouter()

	switch method {
	case "GET":
		r.Get(path, handler)
	case "POST":
		r.Post(path, handler)
	case "PUT":
		r.Put(path, handler)
	case "PATCH":
		r.Patch(path, handler)
	case "DELETE":
		r.Delete(path, handler)
	default:
		panic("Unknown HTTP method: " + method)
	}

	return r
}

// Usage notice
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}
