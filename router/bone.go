// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !bone

package router

import (
	"io"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/julienschmidt/go-http-routing-benchmark/driver"
)

func boneHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func boneHandlerReadWrite(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, bone.GetValue(req, driver.ParamNameReadWrite))
}

func boneFactory(t driver.Type) http.Handler {
	switch t {
	case driver.Static:
		return http.HandlerFunc(boneHandler)
	case driver.Parameterized:
		return http.HandlerFunc(boneHandler)
	case driver.ParameterReadWrite:
		return http.HandlerFunc(boneHandlerReadWrite)
	default:
		panic("Unknown benchmark type passed")
	}
}

func boneRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := bone.New()
	h := boneFactory(t)

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
		Name:     "Bone",
		Router:   boneRouter,
		Homepage: "http://github.com/go-zoo/bone",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
