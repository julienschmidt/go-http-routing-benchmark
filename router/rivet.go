// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !rivet

package router

import (
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/typepress/rivet"
)

func rivetHandler(c rivet.Context) {
	c.Response().WriteHeader(http.StatusOK)
}

func rivetHandlerReadWrite(c rivet.Context) {
	c.WriteString(c.GetParams().Get(driver.ParamNameReadWrite))
}

func rivetFactory(t driver.Type) func(rivet.Context) {
	switch t {
	case driver.Static:
		return rivetHandler
	case driver.Parameterized:
		return rivetHandler
	case driver.ParameterReadWrite:
		return rivetHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func rivetRouter(t driver.Type, f driver.Fixtures) http.Handler {
	h := rivetFactory(t)
	m := rivet.NewRouter(nil)

	for _, r := range f {
		m.Handle(r.Method, r.Path, h)
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "Rivet",
		Router:   rivetRouter,
		Homepage: "http://github.com/typepress/rivet",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
