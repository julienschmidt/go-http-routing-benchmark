// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !possum

package router

import (
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/mikespook/possum"
	"github.com/mikespook/possum/router"
	"github.com/mikespook/possum/view"
)

func possumHandler(c *possum.Context) error {
	c.Response.WriteHeader(http.StatusOK)
	return nil
}

func possumHandlerReadWrite(c *possum.Context) error {
	io.WriteString(c.Response, c.Request.URL.Query().Get(driver.ParamNameReadWrite))
	return nil
}

func possumFactory(t driver.Type) func(*possum.Context) error {
	switch t {
	case driver.Static:
		return possumHandler
	case driver.Parameterized:
		return possumHandler
	case driver.ParameterReadWrite:
		return possumHandlerReadWrite
	default:
		panic("Unknown benchmark type passed")
	}
}

func possumRouter(t driver.Type, f driver.Fixtures) http.Handler {
	h := possumFactory(t)
	m := possum.NewServerMux()

	for _, r := range f {
		var routable router.Router = router.Simple(r.Path)

		if strings.Contains(r.Path, "[^/]*)") {
			routable = router.RegEx(r.Path)
		}

		m.HandleFunc(routable, h, view.Simple("text/html", "utf-8"))
	}

	return m
}

func init() {
	driver.RegisterPackage(&driver.Package{
		Name:       "Possum",
		Router:     possumRouter,
		Normalizer: driver.RegExAllNormalizer,
		Homepage:   "http://github.com/mikespook/possum",
		Supports:   driver.Static | driver.Parameterized, /* | driver.ParameterReadWrite */
	})
}
