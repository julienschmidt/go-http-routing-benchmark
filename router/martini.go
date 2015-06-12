// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !martini

package router

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

func martiniHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func martiniHandlerReadWrite(p martini.Params) string {
	return p[driver.ParamNameReadWrite]
}

func martiniFactory(t driver.Type) interface{} {
	switch t {
	case driver.Static:
		return martiniHandler
	case driver.Parameterized:
		return martiniHandler
	case driver.ParameterReadWrite:
		return martiniHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func martiniRouter(t driver.Type, f driver.Fixtures) http.Handler {
	h := martiniFactory(t)
	m := martini.NewRouter()

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

	martini := martini.New()
	martini.Action(m.Handle)

	return martini
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "Martini",
		Router:   martiniRouter,
		Homepage: "http://github.com/go-martini/martini",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
