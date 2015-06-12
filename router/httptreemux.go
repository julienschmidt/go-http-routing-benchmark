// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !httptreemux

package router

import (
	"io"
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

func httpTreeMuxHandler(w http.ResponseWriter, _ *http.Request, _ map[string]string) {
	w.WriteHeader(http.StatusOK)
}

func httpTreeMuxHandlerReadWrite(w http.ResponseWriter, r *http.Request, p map[string]string) {
	io.WriteString(w, p[driver.ParamNameReadWrite])
}

func httpTreeMuxFactory(t driver.Type) func(http.ResponseWriter, *http.Request, map[string]string) {
	switch t {
	case driver.Static:
		return httpTreeMuxHandler
	case driver.Parameterized:
		return httpTreeMuxHandler
	case driver.ParameterReadWrite:
		return httpTreeMuxHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func httpTreeMuxRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := httptreemux.New()
	h := httpTreeMuxFactory(t)

	for _, r := range f {
		m.Handle(r.Method, r.Path, h)
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "HTTP Tree Mux",
		Router:   httpTreeMuxRouter,
		Homepage: "http://github.com/dimfeld/httptreemux",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
