// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// Parse
// https://parse.com/docs/rest#summary
var parseAPIColon = []route{
	// Objects
	{"POST", "/1/classes/:className"},
	{"GET", "/1/classes/:className/:objectId"},
	{"PUT", "/1/classes/:className/:objectId"},
	{"GET", "/1/classes/:className"},
	{"DELETE", "/1/classes/:className/:objectId"},

	// Users
	{"POST", "/1/users"},
	{"GET", "/1/login"},
	{"GET", "/1/users/:objectId"},
	{"PUT", "/1/users/:objectId"},
	{"GET", "/1/users"},
	{"DELETE", "/1/users/:objectId"},
	{"POST", "/1/requestPasswordReset"},

	// Roles
	{"POST", "/1/roles"},
	{"GET", "/1/roles/:objectId"},
	{"PUT", "/1/roles/:objectId"},
	{"GET", "/1/roles"},
	{"DELETE", "/1/roles/:objectId"},

	// Files
	{"POST", "/1/files/:fileName"},

	// Analytics
	{"POST", "/1/events/:eventName"},

	// Push Notifications
	{"POST", "/1/push"},

	// Installations
	{"POST", "/1/installations"},
	{"GET", "/1/installations/:objectId"},
	{"PUT", "/1/installations/:objectId"},
	{"GET", "/1/installations"},
	{"DELETE", "/1/installations/:objectId"},

	// Cloud Functions
	{"POST", "/1/functions"},
}

var parseAPIBrace = []route{
	// Objects
	{"POST", "/1/classes/{className}"},
	{"GET", "/1/classes/{className}/{objectId}"},
	{"PUT", "/1/classes/{className}/{objectId}"},
	{"GET", "/1/classes/{className}"},
	{"DELETE", "/1/classes/{className}/{objectId}"},

	// Users
	{"POST", "/1/users"},
	{"GET", "/1/login"},
	{"GET", "/1/users/{objectId}"},
	{"PUT", "/1/users/{objectId}"},
	{"GET", "/1/users"},
	{"DELETE", "/1/users/{objectId}"},
	{"POST", "/1/requestPasswordReset"},

	// Roles
	{"POST", "/1/roles"},
	{"GET", "/1/roles/{objectId}"},
	{"PUT", "/1/roles/{objectId}"},
	{"GET", "/1/roles"},
	{"DELETE", "/1/roles/{objectId}"},

	// Files
	{"POST", "/1/files/{fileName}"},

	// Analytics
	{"POST", "/1/events/{eventName}"},

	// Push Notifications
	{"POST", "/1/push"},

	// Installations
	{"POST", "/1/installations"},
	{"GET", "/1/installations/{objectId}"},
	{"PUT", "/1/installations/{objectId}"},
	{"GET", "/1/installations"},
	{"DELETE", "/1/installations/{objectId}"},

	// Cloud Functions
	{"POST", "/1/functions"},
}

var (
	parseBon http.Handler
	parseChi http.Handler
)

func init() {
	println("#ParseAPI Routes:", len(parseAPIColon))

	calcMem("Bon", func() {
		parseBon = loadBon(parseAPIColon)
	})
	calcMem("Chi", func() {
		parseChi = loadChi(parseAPIBrace)
	})

	println()
}

// Static
func BenchmarkBon_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseBon, req)
}
func BenchmarkChi_ParseStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/users", nil)
	benchRequest(b, parseChi, req)
}

// One Param
func BenchmarkBon_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseBon, req)
}
func BenchmarkChi_ParseParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go", nil)
	benchRequest(b, parseChi, req)
}

// Two Params
func BenchmarkBon_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseBon, req)
}
func BenchmarkChi_Parse2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/1/classes/go/123456789", nil)
	benchRequest(b, parseChi, req)
}

// All Routes
func BenchmarkBon_ParseAll(b *testing.B) {
	benchRoutes(b, parseBon, parseAPIColon)
}
func BenchmarkChi_ParseAll(b *testing.B) {
	benchRoutes(b, parseChi, parseAPIBrace)
}
