// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !http_servemux

package router

import (
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

type httpServeMuxHandler struct{}

func (h *httpServeMuxHandler) ServeHTTP(r http.ResponseWriter, _ *http.Request) {
	r.WriteHeader(http.StatusOK)
}

func httpServeMuxFactory(t driver.Type) http.Handler {
	switch t {
	case driver.Static:
		return &httpServeMuxHandler{}
	default:
		panic("Unknown benchmark type passed")
	}
}

func httpServeMuxRouter(t driver.Type, f driver.Fixtures) http.Handler {
	h := httpServeMuxFactory(t)
	m := http.NewServeMux()

	for _, r := range f {
		m.Handle(r.Path, h)
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "http.ServeMux",
		Router:   httpServeMuxRouter,
		Homepage: "http://golang.org/pkg/net/http/#ServeMux",
		Supports: driver.Static,
	})
}
