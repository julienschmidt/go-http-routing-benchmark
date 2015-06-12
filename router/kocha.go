// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !kocha

package router

import (
	"io"
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/naoina/kocha-urlrouter"
	_ "github.com/naoina/kocha-urlrouter/doublearray"
)

type kochaController interface {
	http.Handler
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Patch(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	RouterMap() map[string]urlrouter.URLRouter
}

type kochaHandler struct {
	routerMap map[string]urlrouter.URLRouter
	params    []urlrouter.Param
}

func (h *kochaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method, params := h.routerMap[r.Method].Lookup(r.URL.Path)
	h.params = params

	if method == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		method.(http.HandlerFunc).ServeHTTP(w, r)
	}
}

func (h *kochaHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *kochaHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *kochaHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *kochaHandler) Patch(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *kochaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *kochaHandler) RouterMap() map[string]urlrouter.URLRouter {
	return h.routerMap
}

type kochaHandlerReadWrite struct {
	routerMap map[string]urlrouter.URLRouter
	params    []urlrouter.Param
}

func (h *kochaHandlerReadWrite) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method, params := h.routerMap[r.Method].Lookup(r.URL.Path)
	h.params = params

	if method == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		method.(http.HandlerFunc).ServeHTTP(w, r)
	}
}

func (h *kochaHandlerReadWrite) Post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *kochaHandlerReadWrite) Put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *kochaHandlerReadWrite) Patch(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *kochaHandlerReadWrite) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *kochaHandlerReadWrite) RouterMap() map[string]urlrouter.URLRouter {
	return h.routerMap
}

func (h *kochaHandlerReadWrite) Get(w http.ResponseWriter, r *http.Request) {
	var name string
	for _, param := range h.params {
		if param.Name == driver.ParamNameReadWrite {
			name = param.Value
			break
		}
	}
	io.WriteString(w, name)
}

func kochaFactory(t driver.Type) kochaController {
	switch t {
	case driver.Static, driver.Parameterized:
		return &kochaHandler{routerMap: map[string]urlrouter.URLRouter{
			"GET":    urlrouter.NewURLRouter("doublearray"),
			"POST":   urlrouter.NewURLRouter("doublearray"),
			"PUT":    urlrouter.NewURLRouter("doublearray"),
			"PATCH":  urlrouter.NewURLRouter("doublearray"),
			"DELETE": urlrouter.NewURLRouter("doublearray"),
		}}
	case driver.ParameterReadWrite:
		return &kochaHandlerReadWrite{routerMap: map[string]urlrouter.URLRouter{
			"GET":    urlrouter.NewURLRouter("doublearray"),
			"POST":   urlrouter.NewURLRouter("doublearray"),
			"PUT":    urlrouter.NewURLRouter("doublearray"),
			"PATCH":  urlrouter.NewURLRouter("doublearray"),
			"DELETE": urlrouter.NewURLRouter("doublearray"),
		}}
	default:
		panic("Unknown benchmark type passed")
	}
}

func kochaRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := kochaFactory(t)
	h := make(map[string][]urlrouter.Record)

	for _, r := range f {
		var ch http.HandlerFunc

		switch r.Method {
		case "GET":
			ch = m.Get
		case "POST":
			ch = m.Post
		case "PUT":
			ch = m.Put
		case "PATCH":
			ch = m.Patch
		case "DELETE":
			ch = m.Delete
		}

		h[r.Method] = append(h[r.Method], urlrouter.NewRecord(r.Path, ch))
	}

	for method, records := range h {
		if err := m.RouterMap()[method].Build(records); err != nil {
			panic(err)
		}
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "Kocha",
		Router:   kochaRouter,
		Homepage: "http://github.com/naoina/kocha-urlrouter",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
