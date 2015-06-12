// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !macaron

package router

import (
	"net/http"

	"github.com/Unknwon/macaron"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

func macaronHandler(c *macaron.Context) {
	c.Resp.WriteHeader(http.StatusOK)
}

func macaronHandlerReadWrite(c *macaron.Context) string {
	return c.Params(driver.ParamNameReadWrite)
}

func macaronFactory(t driver.Type) interface{} {
	switch t {
	case driver.Static:
		return macaronHandler
	case driver.Parameterized:
		return macaronHandler
	case driver.ParameterReadWrite:
		return macaronHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func macaronRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := macaron.New()
	h := macaronFactory(t)

	for _, r := range f {
		m.Handle(r.Method, r.Path, []macaron.Handler{h})
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "Macaron",
		Router:   macaronRouter,
		Homepage: "http://github.com/Unknwon/macaron",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
