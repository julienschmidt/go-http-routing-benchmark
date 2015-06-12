// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !goji

package router

import (
	"io"
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/zenazn/goji/web"
)

func gojiHandler(_ web.C, w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func gojiHandlerReadWrite(c web.C, w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, c.URLParams[driver.ParamNameReadWrite])
}

func gojiFactory(t driver.Type) func(web.C, http.ResponseWriter, *http.Request) {
	switch t {
	case driver.Static:
		return gojiHandler
	case driver.Parameterized:
		return gojiHandler
	case driver.ParameterReadWrite:
		return gojiHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func gojiRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := web.New()
	h := gojiFactory(t)

	for _, r := range f {
		switch r.Method {
		case "GET":
			m.Get(r.Path, h)
		case "POST":
			m.Post(r.Path, h)
		case "PUT":
			m.Put(r.Path, h)
		case "PATCH":
			m.Patch(r.Path, h)
		case "DELETE":
			m.Delete(r.Path, h)
		default:
			panic("Unknow HTTP method: " + r.Method)
		}
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "Goji",
		Router:   gojiRouter,
		Homepage: "http://github.com/zenazn/goji/web",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
