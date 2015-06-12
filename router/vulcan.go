// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !vulcan

package router

import (
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/mailgun/route"
)

func vulcanHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func vulcanHandlerReadWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get(driver.ParamNameReadWrite))
}

func vulcanFactory(t driver.Type) func(http.ResponseWriter, *http.Request) {
	switch t {
	case driver.Static:
		return vulcanHandler
	case driver.Parameterized:
		return vulcanHandler
	case driver.ParameterReadWrite:
		return vulcanHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func vulcanRouter(t driver.Type, f driver.Fixtures) http.Handler {
	m := route.NewMux()
	h := vulcanFactory(t)

	for _, r := range f {
		expr := fmt.Sprintf(`Method("%s") && Path("%s")`, r.Method, r.Path)

		if err := m.HandleFunc(expr, h); err != nil {
			panic(err)
		}
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:       "Mailgun Vulcan",
		Router:     vulcanRouter,
		Normalizer: driver.XThanSignNormalizer,
		Homepage:   "http://github.com/mailgun/route",
		Supports:   driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
