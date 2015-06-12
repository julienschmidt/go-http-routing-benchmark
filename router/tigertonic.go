// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !tigertonic

package router

import (
	"io"
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/rcrowley/go-tigertonic"
)

const tigerTonicRWP = ":" + driver.ParamNameReadWrite

func tigerTonicHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func tigerTonicHandlerReadWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get(tigerTonicRWP))
}

func tigerTonicFactory(t driver.Type) func(http.ResponseWriter, *http.Request) {
	switch t {
	case driver.Static:
		return tigerTonicHandler
	case driver.Parameterized:
		return tigerTonicHandler
	case driver.ParameterReadWrite:
		return tigerTonicHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func tigerTonicRouter(t driver.Type, f driver.Fixtures) http.Handler {
	h := tigerTonicFactory(t)
	m := tigertonic.NewTrieServeMux()

	for _, r := range f {
		m.HandleFunc(r.Method, r.Path, h)
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:       "Tiger Tonic",
		Router:     tigerTonicRouter,
		Normalizer: driver.CurlyBracesNormalizer,
		Homepage:   "http://github.com/rcrowley/go-tigertonic",
		Supports:   driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
