// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !httprouter

package router

import (
	"io"
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/julienschmidt/httprouter"
)

func httpRouterHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func httpRouterHandlerReadWrite(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, p.ByName(driver.ParamNameReadWrite))
}

func httpRouterFactory(t driver.Type) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	switch t {
	case driver.Static:
		return httpRouterHandler
	case driver.Parameterized:
		return httpRouterHandler
	case driver.ParameterReadWrite:
		return httpRouterHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func httpRouterRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := httprouter.New()
	h := httpRouterFactory(t)

	for _, r := range f {
		m.Handle(r.Method, r.Path, h)
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "HTTP Router",
		Router:   httpRouterRouter,
		Homepage: "http://github.com/julienschmidt/httprouter",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
