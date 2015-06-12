// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !gorilla_mux

package router

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

func gorillaHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func gorillaHandlerReadWrite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	io.WriteString(w, params[driver.ParamNameReadWrite])
}

func gorillaFactory(t driver.Type) func(http.ResponseWriter, *http.Request) {
	switch t {
	case driver.Static:
		return gorillaHandler
	case driver.Parameterized:
		return gorillaHandler
	case driver.ParameterReadWrite:
		return gorillaHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func gorillaRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := mux.NewRouter()
	h := gorillaFactory(t)

	for _, r := range f {
		m.HandleFunc(r.Path, h).Methods(r.Method)
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:       "Gorilla Mux",
		Router:     gorillaRouter,
		Normalizer: driver.CurlyBracesNormalizer,
		Homepage:   "http://github.com/gorilla/mux",
		Supports:   driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
