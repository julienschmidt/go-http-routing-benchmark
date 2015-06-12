// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !beego

package router

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

const beegoRWP = ":" + driver.ParamNameReadWrite

func beegoHandler(ctx *context.Context) {
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
}

func beegoHandlerReadWrite(ctx *context.Context) {
	ctx.WriteString(ctx.Input.Param(beegoRWP))
}

func beegoFactory(t driver.Type) func(ctx *context.Context) {
	switch t {
	case driver.Static:
		return beegoHandler
	case driver.Parameterized:
		return beegoHandler
	case driver.ParameterReadWrite:
		return beegoHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func beegoRouter(t driver.Type, f driver.Fixtures) http.Handler {
	h := beegoFactory(t)
	m := beego.NewControllerRegister()

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
	beego.RunMode = "prod"
	beego.BeeLogger.Close()

	driver.RegisterPackage(&driver.Package{
		Name:     "Beego",
		Router:   beegoRouter,
		Homepage: "http://beego.me/",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
