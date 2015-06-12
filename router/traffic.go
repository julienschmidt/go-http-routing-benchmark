// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !traffic

package router

import (
	"io"
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/pilu/traffic"
)

func trafficHandler(w traffic.ResponseWriter, _ *traffic.Request) {
	w.WriteHeader(http.StatusOK)
}

func trafficHandlerReadWrite(w traffic.ResponseWriter, r *traffic.Request) {
	io.WriteString(w, r.URL.Query().Get(driver.ParamNameReadWrite))
}

func trafficFactory(t driver.Type) func(traffic.ResponseWriter, *traffic.Request) {
	switch t {
	case driver.Static:
		return trafficHandler
	case driver.Parameterized:
		return trafficHandler
	case driver.ParameterReadWrite:
		return trafficHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func trafficRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := traffic.New()
	h := trafficFactory(t)

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
	traffic.SetVar("env", "bench")

	driver.RegisterPackage(&driver.Package{
		Name:     "Traffic",
		Router:   trafficRouter,
		Homepage: "http://github.com/pilu/traffic",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
