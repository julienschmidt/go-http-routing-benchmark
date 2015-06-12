// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

// +build !revel

package router

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"path/filepath"

	"go/build"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/revel/revel"
	"github.com/robfig/pathtree"
)

// In the following code some Revel internals are modelled.
// The original revel code is copyrighted by Rob Figueiredo.
// See https://github.com/revel/revel/blob/master/LICENSE
type RevelController struct {
	*revel.Controller
	router *revel.Router
}

func (rc *RevelController) Handle() revel.Result {
	return revelResult{}
}

func (rc *RevelController) HandleReadWrite() revel.Result {
	return rc.RenderText(rc.Params.Get(driver.ParamNameReadWrite))
}

func (rc *RevelController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Dirty hacks, do NOT copy!
	revel.MainRouter = rc.router

	upgrade := r.Header.Get("Upgrade")
	if upgrade == "websocket" || upgrade == "Websocket" {
		panic("Not implemented")
	}

	var (
		req  = revel.NewRequest(r)
		resp = revel.NewResponse(w)
		c    = revel.NewController(req, resp)
	)
	req.Websocket = nil

	revel.Filters[0](c, revel.Filters[1:])

	if c.Result != nil {
		c.Result.Apply(req, resp)
	} else if c.Response.Status != 0 {
		c.Response.Out.WriteHeader(c.Response.Status)
	} else {
		panic("Result is empty")
	}
	// Close the Writer if we can
	if w, ok := resp.Out.(io.Closer); ok {
		w.Close()
	}
}

type revelResult struct{}

func (rr revelResult) Apply(req *revel.Request, resp *revel.Response) {
	resp.Out.WriteHeader(http.StatusOK)
}

func revelFactory(t driver.Type) string {
	switch t {
	case driver.Static:
		return "RevelController.Handle"
	case driver.Parameterized:
		return "RevelController.Handle"
	case driver.ParameterReadWrite:
		return "RevelController.HandleReadWrite"
	default:
		panic("Unknown benchmark type passed")
	}
}

func revelRouter(t driver.Type, f driver.Fixtures) http.Handler {
	h := revelFactory(t)
	m := revel.NewRouter("")

	rs := make([]*revel.Route, 0, len(f))
	for _, r := range f {
		rs = append(rs, revel.NewRoute(r.Method, r.Path, h, "", "", 0))
	}
	m.Routes = rs

	m.Tree = pathtree.New()
	for _, r := range m.Routes {
		err := m.Tree.Add(r.TreePath, r)

		// Allow GETs to respond to HEAD requests.
		if err == nil && r.Method == "GET" {
			err = m.Tree.Add("/HEAD"+r.Path, r)
		}

		// Error adding a route to the pathtree.
		if err != nil {
			panic(err)
		}
	}

	rc := &RevelController{}
	rc.router = m

	return rc
}

// findSrcPaths uses the "go/build" package to find the source root for Revel
// and the app.
//
// Directly copied from revel/revel.go.
// The original revel code is copyrighted by Rob Figueiredo.
// See https://github.com/revel/revel/blob/master/LICENSE
func revelFindSrcPaths(importPath string) (revelSourcePath, appSourcePath string) {
	var (
		gopaths = filepath.SplitList(build.Default.GOPATH)
		goroot  = build.Default.GOROOT
	)

	if len(gopaths) == 0 {
		panic("GOPATH environment variable is not set. " +
			"Please refer to http://golang.org/doc/code.html to configure your Go environment.")
	}

	if revelContainsString(gopaths, goroot) {
		panic(fmt.Sprintf("GOPATH (%s) must not include your GOROOT (%s). "+
			"Please refer to http://golang.org/doc/code.html to configure your Go environment.",
			gopaths, goroot))
	}

	appPkg, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		panic("Failed to import" + importPath + "with error:" + err.Error())
	}

	revelPkg, err := build.Import(revel.REVEL_IMPORT_PATH, "", build.FindOnly)
	if err != nil {
		panic("Failed to find Revel with error:" + err.Error())
	}

	return revelPkg.SrcRoot, appPkg.SrcRoot
}

// Directly copied from revel/utils.go.
// The original revel code is copyrighted by Rob Figueiredo.
// See https://github.com/revel/revel/blob/master/LICENSE
func revelContainsString(list []string, target string) bool {
	for _, el := range list {
		if el == target {
			return true
		}
	}
	return false
}

func revelInit() {
	// Only use the Revel filters required for this benchmark
	revel.Filters = []revel.Filter{
		revel.RouterFilter,
		revel.ParamsFilter,
		revel.ActionInvoker,
	}

	revel.RegisterController(
		(*RevelController)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Handle",
			},
			&revel.MethodType{
				Name: "HandleReadWrite",
			},
		},
	)

	// Load mime config from revel, otherwise revel panics (nil pointer dereference)
	srcPath, _ := revelFindSrcPaths(revel.REVEL_IMPORT_PATH)
	revel.ConfPaths = []string{path.Join(
		srcPath,
		filepath.FromSlash(revel.REVEL_IMPORT_PATH),
		"conf"),
	}
	revel.LoadMimeConfig()

	// Must be set before instantiationg the TemplateLoader object to avoid panics.
	revel.RevelPath = path.Join(srcPath, filepath.FromSlash(revel.REVEL_IMPORT_PATH))
	revel.TemplatePaths = []string{path.Join(revel.RevelPath, "templates")}

	// Otherwise revel panics (nil pointer dereference) in revel/results.go:47
	// ErrorResult.Apply: tmpl, err := MainTemplateLoader.Template(templatePath)
	revel.MainTemplateLoader = revel.NewTemplateLoader(revel.TemplatePaths)
	revel.MainTemplateLoader.Refresh()
}

func init() {
	revelInit()

	driver.RegisterPackage(&driver.Package{
		Name:     "Revel",
		Router:   revelRouter,
		Homepage: "http://github.com/revel/revel",
		Supports: driver.Static | driver.Parameterized | driver.ParameterReadWrite,
	})
}
