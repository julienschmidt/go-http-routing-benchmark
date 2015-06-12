// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !echo

package router

import (
	"io"
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/labstack/echo"
)

func echoHandler(c *echo.Context) error {
	c.Response().WriteHeader(http.StatusOK)
	return nil
}

func echoHandlerReadWrite(c *echo.Context) error {
	io.WriteString(c.Response(), c.Param(driver.ParamNameReadWrite))
	return nil
}

func echoFactory(t driver.Type) func(*echo.Context) error {
	switch t {
	case driver.Static:
		return echoHandler
	case driver.Parameterized:
		return echoHandler
	case driver.ParameterReadWrite:
		return echoHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func echoRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := echo.New()
	h := echoFactory(t)

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

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:     "Echo",
		Router:   echoRouter,
		Homepage: "http://github.com/labstack/echo",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
