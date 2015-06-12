// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !ace

package router

import (
	"io"
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/plimble/ace"
)

func aceHandler(c *ace.C) {
	c.Writer.WriteHeader(http.StatusOK)
}

func aceHandlerReadWrite(c *ace.C) {
	io.WriteString(c.Writer, c.Param(driver.ParamNameReadWrite))
}

func aceFactory(t driver.Type) []ace.HandlerFunc {
	switch t {
	case driver.Static:
		return []ace.HandlerFunc{aceHandler}
	case driver.Parameterized:
		return []ace.HandlerFunc{aceHandler}
	case driver.ParameterReadWrite:
		return []ace.HandlerFunc{aceHandlerReadWrite}
	default:
		panic("Unknown benchmark type passed")
	}
}

func aceRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := ace.New()
	h := aceFactory(t)

	for _, r := range f {
		m.Handle(r.Method, r.Path, h)
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "Ace",
		Router:   aceRouter,
		Homepage: "http://github.com/plimble/ace",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
