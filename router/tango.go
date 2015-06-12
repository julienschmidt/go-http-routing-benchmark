// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !tango

package router

import (
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	llog "github.com/lunny/log"
	"github.com/lunny/tango"
)

const tangoRWP = ":" + driver.ParamNameReadWrite

func tangoHandler(c *tango.Context) {
	c.ResponseWriter.WriteHeader(http.StatusOK)
}

func tangoHandlerReadWrite(c *tango.Context) {
	c.Write([]byte(c.Params().Get(tangoRWP)))
}

func tangoFactory(t driver.Type) func(*tango.Context) {
	switch t {
	case driver.Static:
		return tangoHandler
	case driver.Parameterized:
		return tangoHandler
	case driver.ParameterReadWrite:
		return tangoHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func tangoRouter(t driver.Type, f driver.Fixtures) http.Handler {
	h := tangoFactory(t)
	m := tango.NewWithLog(llog.Std)

	for _, r := range f {
		m.Route(r.Method, r.Path, h)
	}

	return m
}

func init() {
	llog.SetOutputLevel(llog.Lnone)
	llog.SetOutput(&driver.ResponseStub{})

	driver.RegisterPackage(&driver.Package{
		Name:     "Tango",
		Router:   tangoRouter,
		Homepage: "http://github.com/lunny/tango",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
