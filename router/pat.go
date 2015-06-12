// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !martini

package router

import (
	"io"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

const patRWP = ":" + driver.ParamNameReadWrite

func patHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func patHandlerReadWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get(patRWP))
}

func patFactory(t driver.Type) http.Handler {
	switch t {
	case driver.Static:
		return http.HandlerFunc(patHandler)
	case driver.Parameterized:
		return http.HandlerFunc(patHandler)
	case driver.ParameterReadWrite:
		return http.HandlerFunc(patHandlerReadWrite)
	default:
		panic("Unknown benchmark type passed")
	}
}

func patRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := pat.New()
	h := patFactory(t)

	for _, r := range f {
		switch r.Method {
		case "GET":
			m.Get(r.Path, h)
		case "POST":
			m.Post(r.Path, h)
		case "PUT":
			m.Put(r.Path, h)
		case "DELETE":
			m.Del(r.Path, h)
		default:
			panic("Unknow HTTP method: " + r.Method)
		}
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "Pat",
		Router:   patRouter,
		Homepage: "http://github.com/bmizerany/pat",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
