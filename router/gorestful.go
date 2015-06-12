// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !gorestful

package router

import (
	"io"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

func goRestfulHandler(_ *restful.Request, w *restful.Response) {
	// Cannot use w.WriteHeader for some reason
	w.ResponseWriter.WriteHeader(http.StatusOK)
}

func goRestfulHandlerReadWrite(r *restful.Request, w *restful.Response) {
	io.WriteString(w, r.PathParameter(driver.ParamNameReadWrite))
}

func goRestfulFactory(t driver.Type) func(*restful.Request, *restful.Response) {
	switch t {
	case driver.Static:
		return goRestfulHandler
	case driver.Parameterized:
		return goRestfulHandler
	case driver.ParameterReadWrite:
		return goRestfulHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func goRestfulRouter(t driver.Type, f driver.Fixtures) http.Handler {
	h := goRestfulFactory(t)
	s := &restful.WebService{}
	m := restful.NewContainer()

	for _, r := range f {
		switch r.Method {
		case "GET":
			s.Route(s.GET(r.Path).To(h))
		case "POST":
			s.Route(s.POST(r.Path).To(h))
		case "PUT":
			s.Route(s.PUT(r.Path).To(h))
		case "PATCH":
			s.Route(s.PATCH(r.Path).To(h))
		case "DELETE":
			s.Route(s.DELETE(r.Path).To(h))
		default:
			panic("Unknow HTTP method: " + r.Method)
		}
	}

	m.Add(s)

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:       "Go Restful",
		Router:     goRestfulRouter,
		Normalizer: driver.CurlyBracesNormalizer,
		Homepage:   "http://github.com/emicklei/go-restful",
		Supports:   driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
