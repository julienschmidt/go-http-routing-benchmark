// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !gojsonrest

package router

import (
	"io"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

func goJsonRestHandler(w rest.ResponseWriter, _ *rest.Request) {
	w.WriteHeader(http.StatusOK)
}

func goJsonRestHandlerReadWrite(w rest.ResponseWriter, r *rest.Request) {
	io.WriteString(w.(io.Writer), r.PathParam(driver.ParamNameReadWrite))
}

func goJsonRestFactory(t driver.Type) func(rest.ResponseWriter, *rest.Request) {
	switch t {
	case driver.Static:
		return goJsonRestHandler
	case driver.Parameterized:
		return goJsonRestHandler
	case driver.ParameterReadWrite:
		return goJsonRestHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func goJsonRestRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := rest.NewApi()
	h := goJsonRestFactory(t)

	restRoutes := make([]*rest.Route, 0, len(f))

	for _, r := range f {
		restRoutes = append(restRoutes,
			&rest.Route{r.Method, r.Path, h},
		)
	}

	router, err := rest.MakeRouter(restRoutes...)
	if err != nil {
		log.Fatal(err)
	}

	m.SetApp(router)
	return m.MakeHandler()
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "Go Json Rest",
		Router:   goJsonRestRouter,
		Homepage: "http://github.com/ant0ine/go-json-rest/rest",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
