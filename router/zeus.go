// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build ignore,!zeus

package router

import (
	"io"
	"net/http"

	"github.com/daryl/zeus"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

func zeusHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func zeusHandlerReadWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, zeus.Var(r, driver.ParamNameReadWrite))
}

func zeusFactory(t driver.Type) func(http.ResponseWriter, *http.Request) {
	switch t {
	case driver.Static:
		return zeusHandler
	case driver.Parameterized:
		return zeusHandler
	case driver.ParameterReadWrite:
		return zeusHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func zeusRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := zeus.New()
	h := zeusFactory(t)

	for _, r := range f {
		switch r.Method {
		case "GET":
			m.GET(r.Path, h)
		case "POST":
			m.POST(r.Path, h)
		case "PUT":
			m.PUT(r.Path, h)
		case "DELETE":
			m.DELETE(r.Path, h)
		default:
			panic("Unknow HTTP method: " + r.Method)
		}
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "Zeus",
		Router:   zeusRouter,
		Homepage: "http://github.com/daryl/zeus",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
