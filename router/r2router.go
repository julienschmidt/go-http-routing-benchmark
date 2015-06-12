// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !r2router

package router

import (
	"io"
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/vanng822/r2router"
)

func r2RouterHandler(w http.ResponseWriter, _ *http.Request, _ r2router.Params) {
	w.WriteHeader(http.StatusOK)
}

func r2RouterHandlerReadWrite(w http.ResponseWriter, _ *http.Request, p r2router.Params) {
	io.WriteString(w, p.Get(driver.ParamNameReadWrite))
}

func r2RouterFactory(t driver.Type) func(http.ResponseWriter, *http.Request, r2router.Params) {
	switch t {
	case driver.Static:
		return r2RouterHandler
	case driver.Parameterized:
		return r2RouterHandler
	case driver.ParameterReadWrite:
		return r2RouterHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func r2RouterRouter(t driver.Type, f driver.Fixtures) http.Handler {
	h := r2RouterFactory(t)
	m := r2router.NewRouter()

	for _, r := range f {
		m.AddHandler(r.Method, r.Path, h)
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "R2 Router",
		Router:   r2RouterRouter,
		Homepage: "http://github.com/vanng822/r2router",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
