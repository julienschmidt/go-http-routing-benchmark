// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !denco

package router

import (
	"io"
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/naoina/denco"
)

func dencoHandler(w http.ResponseWriter, _ *http.Request, params denco.Params) {
	w.WriteHeader(http.StatusOK)
}

func dencoHandlerReadWrite(w http.ResponseWriter, req *http.Request, params denco.Params) {
	io.WriteString(w, params.Get(driver.ParamNameReadWrite))
}

func dencoFactory(t driver.Type) func(http.ResponseWriter, *http.Request, denco.Params) {
	switch t {
	case driver.Static:
		return dencoHandler
	case driver.Parameterized:
		return dencoHandler
	case driver.ParameterReadWrite:
		return dencoHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func dencoRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := denco.NewMux()
	h := dencoFactory(t)

	handlers := make([]denco.Handler, 0, len(f))

	for _, r := range f {
		handler := m.Handler(r.Method, r.Path, h)
		handlers = append(handlers, handler)
	}

	handler, err := m.Build(handlers)

	if err != nil {
		panic(err)
	}

	return handler
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "Denco",
		Router:   dencoRouter,
		Homepage: "http://github.com/naoina/denco",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
