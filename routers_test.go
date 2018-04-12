package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// load functions of all routers
	routers = []struct {
		name string
		load func(routes []route) http.Handler
	}{
		{"Bon", loadBon},
		{"Chi", loadChi},
		{"Denco", loadDenco},
		{"Gin", loadGin},
		{"Echo", loadEcho},
	}

	// all APIs
	apis = []struct {
		name   string
		routes []route
	}{
		{"GitHub", githubAPIColon},
		{"GPlus", gplusAPIColon},
		{"Parse", parseAPIColon},
		{"Static", staticRoutes},
	}
)

func TestRouters(t *testing.T) {
	loadTestHandler = true

	for _, router := range routers {
		req, _ := http.NewRequest("GET", "/", nil)
		u := req.URL
		rq := u.RawQuery

		for _, api := range apis {
			r := router.load(api.routes)

			for _, route := range api.routes {
				w := httptest.NewRecorder()
				req.Method = route.method
				req.RequestURI = route.path
				u.Path = route.path
				u.RawQuery = rq
				r.ServeHTTP(w, req)
				if w.Code != 200 || w.Body.String() != route.path {
					t.Errorf(
						"%s in API %s: %d - %s; expected %s %s\n",
						router.name, api.name, w.Code, w.Body.String(), route.method, route.path,
					)
				}
			}
		}
	}

	loadTestHandler = false
}
