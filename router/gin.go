// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !gin

package router

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

func ginHandler(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}

func ginHandlerReadWrite(c *gin.Context) {
	io.WriteString(c.Writer, c.Params.ByName(driver.ParamNameReadWrite))
}

func ginFactory(t driver.Type) func(*gin.Context) {
	switch t {
	case driver.Static:
		return ginHandler
	case driver.Parameterized:
		return ginHandler
	case driver.ParameterReadWrite:
		return ginHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func ginRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := gin.New()
	h := ginFactory(t)

	for _, r := range f {
		m.Handle(r.Method, r.Path, h)
	}

	return m
}

func init() {
	gin.SetMode(gin.ReleaseMode)

	driver.RegisterPackage(&driver.Package{
		Name:     "Gin",
		Router:   ginRouter,
		Homepage: "http://github.com/gin-gonic/gin",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
