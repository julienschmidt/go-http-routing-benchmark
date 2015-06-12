// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !gocraftweb

package router

import (
	"io"
	"net/http"

	"github.com/gocraft/web"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

type gocraftWebContext struct{}

func gocraftWebHandler(w web.ResponseWriter, _ *web.Request) {
	w.WriteHeader(http.StatusOK)
}

func gocraftWebHandlerReadWrite(w web.ResponseWriter, r *web.Request) {
	io.WriteString(w, r.PathParams[driver.ParamNameReadWrite])
}

func gocraftWebFactory(t driver.Type) func(web.ResponseWriter, *web.Request) {
	switch t {
	case driver.Static:
		return gocraftWebHandler
	case driver.Parameterized:
		return gocraftWebHandler
	case driver.ParameterReadWrite:
		return gocraftWebHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func gocraftWebRouter(t driver.Type, f driver.Fixtures) http.Handler {
	h := gocraftWebFactory(t)
	m := web.New(gocraftWebContext{})

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
		Name:     "GoCraft Web",
		Router:   gocraftWebRouter,
		Homepage: "http://github.com/gocraft/web",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
