// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !bear

package router

import (
	"io"
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/ursiform/bear"
)

func bearHandler(w http.ResponseWriter, _ *http.Request, _ *bear.Context) {
	w.WriteHeader(http.StatusOK)
}

func bearHandlerReadWrite(w http.ResponseWriter, _ *http.Request, ctx *bear.Context) {
	io.WriteString(w, ctx.Params[driver.ParamNameReadWrite])
}

func bearFactory(t driver.Type) func(http.ResponseWriter, *http.Request, *bear.Context) {
	switch t {
	case driver.Static:
		return bearHandler
	case driver.Parameterized:
		return bearHandler
	case driver.ParameterReadWrite:
		return bearHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func bearRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := bear.New()
	h := bearFactory(t)

	for _, r := range f {
		switch r.Method {
		case "GET", "POST", "PUT", "PATCH", "DELETE":
			m.On(r.Method, r.Path, h)
		default:
			panic("Unknown HTTP method: " + r.Method)
		}
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:       "Bear",
		Router:     bearRouter,
		Normalizer: driver.CurlyBracesNormalizer,
		Homepage:   "http://github.com/ursiform/bear",
		Supports:   driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
