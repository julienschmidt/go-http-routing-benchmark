// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package driver

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"
)

// Type defines a benchmark type
type Type uint

// isset determines wether b is set.
func (t Type) isset(b Type) bool {
	return t&b != 0
}

const (
	// Benchmark route with a static path. This benchmark must only contain
	// static routes.
	//
	// Route:   GET /test/example
	// Request: GET /test/example
	Static Type = 1 << 0

	// Benchmark routes with one or more dynamic path segments. This benchmark
	// may also contain static routes.
	//
	// Route:     GET /test/:example
	// Request:   GET /test/benchmark
	// Parameter: example=benchmark
	Parameterized Type = 1 << 1

	// Benchmark a route with multiple dynamic path segments where one parameter
	// (dynamicparam, see ParamNameReadWrite) is read and written to /dev/null.
	// This benchmark must not contain other types of benchmarks.
	//
	// Route:     GET /test/:one/:two/:three/:dynamicparam
	// Request:   GET /test/one/two/three/benchmark
	// Parameter: one=one, two=two, three=three, dynamicparam=benchmark
	// Response:  benchmark
	ParameterReadWrite Type = 1 << 2

	// ParameterReadWrite benchmark expects a parameter named dynamicparam.
	ParamNameReadWrite = "dynamicparam"
)

// Fixture is used to benchmark and test Router implementations.
type Fixture struct {
	// Method used for routing and requesting the endpoint.
	Method string

	// Path used for routing and requesting the endpoint.
	Path string
}

// Fixtures is a list of Fixture objects.
type Fixtures []*Fixture

// Router initializes and returns a http.Handler used to route requests.
type Router func(Type, Fixtures) http.Handler

// Normalizer normalizes parameterized routes before they get passed to the Router.
// Route normalization is importatnt for Routers that differ in the way they
// indicate a parameter (e.g. not prefixing it with a colon).
type Normalizer func(Fixtures) Fixtures

var nre = regexp.MustCompile(":([^/]*)")

// CurlyBracesNormalizer replaces colon (:param) parameters with curly braces
// (i.e. {param}).
func CurlyBracesNormalizer(f Fixtures) Fixtures {
	s := make(Fixtures, len(f))

	for i, r := range f {
		s[i] = &Fixture{r.Method, nre.ReplaceAllString(r.Path, "{$1}")}
	}

	return s
}

// XThanSignNormalizer replaces colon (:param) parameters with greater- and
// less-than sign (i.e. <param>).
func XThanSignNormalizer(f Fixtures) Fixtures {
	s := make(Fixtures, len(f))

	for i, r := range f {
		s[i] = &Fixture{r.Method, nre.ReplaceAllString(r.Path, "<$1>")}
	}

	return s
}

// RegExAllNormalizer replaces colon (:param) parameters with a regular expression
// which matches everything.except a slash
func RegExAllNormalizer(f Fixtures) Fixtures {
	s := make(Fixtures, len(f))

	for i, r := range f {
		s[i] = &Fixture{r.Method, nre.ReplaceAllString(r.Path, "(?P<$1>[^/]*)")}
	}

	return s
}

// Benchmark describes a benchmark. A benchmark may only cover one benchmark type.
type Benchmark struct {
	Type        Type
	Name        string
	Description string
	tests       []*Test
	routes      Fixtures
	requests    []*http.Request
}

var benchmarks []*Benchmark

// NewBenchmark registers and returns a new benchmark object. This function
// panics on error.
func NewBenchmark(t Type, name, desc string) *Benchmark {
	var errors []string

	if t == Type(0) {
		errors = append(errors, "Benchmark type cannot be zero")
	}
	if t&(t-Type(1)) != Type(0) {
		errors = append(errors, "Benchmark cannot cover multiple types")
	}
	if name == "" {
		errors = append(errors, "Benchmark name cannot be empty")
	}
	if desc == "" {
		errors = append(errors, "Benchmark description cannot be empty")
	}

	if errors != nil {
		panic(strings.Join(errors, ", "))
	}

	b := &Benchmark{
		Type:        t,
		Name:        name,
		Description: desc,
	}
	benchmarks = append(benchmarks, b)

	return b
}

// AddTest registers a new test case which will be run against a Router
// implementation. This function panics on error.
func (b *Benchmark) AddTest(t *Test) {
	var errors []string

	if t.Description == "" {
		errors = append(errors, "Test.Description cannot be empty")
	}
	if t.Route == nil && t.Request == nil {
		errors = append(errors, "Test.Route or Test.Request must be set")
	}

	if errors != nil {
		panic(strings.Join(errors, ", "))
	}

	b.tests = append(b.tests, t)
}

// AddRoute registers a new route which will be registered with a Router.
// This function panics on error.
func (b *Benchmark) AddRoute(method, path string) {
	var errors []string

	if method == "" {
		errors = append(errors, "Method cannot be empty")
	}
	if path == "" {
		errors = append(errors, "Path cannot be empty")
	}

	if errors != nil {
		panic(strings.Join(errors, ", "))
	}

	b.routes = append(b.routes, &Fixture{method, path})
}

// AddRoutes registers routes based on Fixtures which will be registered with a Router.
// This function panics on error.
func (b *Benchmark) AddRoutes(fs Fixtures) {
	var errors []string

	for i, f := range fs {
		if f.Method == "" {
			errors = append(errors, fmt.Sprintf("Fixture #%d: Method cannot be empty", i))
		}
		if f.Path == "" {
			errors = append(errors, fmt.Sprintf("Fixture #%d: Path cannot be empty", i))
		}

		b.routes = append(b.routes, f)
	}

	if errors != nil {
		panic(strings.Join(errors, ", "))
	}
}

var request = strings.NewReplacer("/:", "/")

// AddRoute registers a new route which will be registered with a Router.
// This function panics on error.
func (b *Benchmark) AddRequest(method, path string) {
	var errors []string
	r, err := http.NewRequest(method, request.Replace(path), nil)

	if err != nil {
		errors = append(errors, err.Error())
	}
	if method == "" {
		errors = append(errors, "Method cannot be empty")
	}
	if path == "" {
		errors = append(errors, "Path cannot be empty")
	}

	if errors != nil {
		panic(strings.Join(errors, ", "))
	}

	b.requests = append(b.requests, r)
}

// AddRequests generates requests based on Fixtures. The paths will be striped
// of parameters (/:). // This function panics on error.
func (b *Benchmark) AddRequests(fs Fixtures) {
	var errors []string

	for i, f := range fs {
		r, err := http.NewRequest(f.Method, request.Replace(f.Path), nil)

		if err != nil {
			errors = append(errors, err.Error())
		}
		if f.Method == "" {
			errors = append(errors, fmt.Sprintf("Fixture #%d: Method cannot be empty", i))
		}
		if f.Path == "" {
			errors = append(errors, fmt.Sprintf("Fixture #%d: Path cannot be empty", i))
		}

		b.requests = append(b.requests, r)
	}

	if errors != nil {
		panic(strings.Join(errors, ", "))
	}
}

// Package describes a router or framework. Every field except for Normalizer
// is required.
type Package struct {
	Name       string     // Name of the Router
	Homepage   string     // Homepage or repository URL for the router/framework
	Router     Router     // Router to initialize a http.Handler
	Supports   Type       // Bitmask of supported benchmark Type's
	Normalizer Normalizer // Normalizer to run before routes get passed to the Router
}

var packages []*Package

// Returns the number of registered packages.
func NumPackages() int {
	return len(packages)
}

// Returns the names of registered packages.
func PackageNames() []string {
	r := make([]string, len(packages))
	for i, pkg := range packages {
		r[i] = pkg.Name
	}

	return r
}

// RegisterPackage registers router/framework package.
func RegisterPackage(p *Package) {
	var errors []string

	if p == nil {
		errors = append(errors, "Package cannot be nil")
	} else {
		if p.Name == "" {
			errors = append(errors, "Package.Name cannot be empty")
		}
		if p.Homepage == "" {
			errors = append(errors, "Package.Homepage cannot be empty")
		} else if _, err := url.Parse(p.Homepage); p != nil && err != nil {
			errors = append(errors, err.Error())
		}
		if p.Router == nil {
			errors = append(errors, "Package.Router cannot be nil")
		}
		if p.Supports == Type(0) {
			errors = append(errors, "Package.Supports cannot be zero")
		}
	}

	if errors != nil {
		panic(strings.Join(errors, ", "))
	}

	packages = append(packages, p)
}

// Test describes a test case used to test that benchmarks are configured correctly.
// A Test function may return a pointer to a Assertion object or nil, depending whether
// the test was successful or not.
type Test struct {
	// Description for the test case.
	Description string

	// Tests the raw, not normalized routes.
	Route func(*Fixture) *Assertion

	// Tests the http.Request or Response object.
	Request func(*http.Request, *Response) *Assertion
}

// Assertion is used in test cases to indicate that an error occured or to compare
// the expected output is equal to the actual output. If the Error property is
// set, Actual and Expect must be empty and vice versa. The comparison of the expected
// and actual output is not part of the test and will be handled externaly.
type Assertion struct {
	test   *Test
	pkg    *Package
	bench  *Benchmark
	Error  string
	Actual string
	Expect string
}

// Returns the description of the test that failed.
func (a *Assertion) Description() string {
	if a.test == nil {
		return ""
	}

	return a.test.Description
}

// Returns the router name of the test that failed.
func (a *Assertion) Router() string {
	if a.pkg == nil {
		return ""
	}

	return a.pkg.Name
}

// Returns the benchmark description of the test that failed.
func (a *Assertion) Benchmark() string {
	if a.bench == nil {
		return ""
	}

	return a.bench.Name + ": " + a.bench.Description
}

// Result is a benchmark result.
type Result struct {
	benchmark *Benchmark
	alloc     []uint64
	pkgs      []*Package
	results   []*testing.BenchmarkResult
}

// Name returns the name of the benchmark.
func (r *Result) Name() string {
	if r.benchmark == nil {
		panic("benchmark is nil")
	}

	return r.benchmark.Name
}

// Description returns the description of the benchmark.
func (r *Result) Description() string {
	if r.benchmark == nil {
		panic("benchmark is nil")
	}

	return r.benchmark.Description
}

// Package returns all packages which were benchmarked.
func (r *Result) Packages() []*Package {
	if len(r.pkgs) != len(r.results) && len(r.pkgs) != len(r.alloc) {
		panic("Package, BenchmarkResult and HeapMemory mismatch")
	}

	return r.pkgs
}

// RouterMemory returns the heap memory used for the router object.
func (r *Result) RouterMemory() []uint64 {
	if len(r.pkgs) != len(r.results) && len(r.pkgs) != len(r.alloc) {
		panic("Package, BenchmarkResult and HeapMemory mismatch")
	}

	return r.alloc
}

// Result returns the benchmarks result of all benchmarked packages.
func (r *Result) Results() []*testing.BenchmarkResult {
	if len(r.pkgs) != len(r.results) && len(r.pkgs) != len(r.alloc) {
		panic("Package, BenchmarkResult and HeapMemory mismatch")
	}

	return r.results
}

// Response is used in during tests and captures the written response.
type Response struct {
	flushed bool
	status  int
	output  bytes.Buffer
	header  http.Header
}

// Reset's the captured status code and output.
func (r *Response) Reset() {
	r.status = 0
	r.header = nil
	r.flushed = false
	r.output.Reset()
}

// StatusCode returns the written status code.
func (r *Response) Status() int {
	return r.status
}

// Output returns the written output.
func (r *Response) Output() []byte {
	return r.output.Bytes()
}

// Header returns the header map that will be sent by WriteHeader.
func (r *Response) Header() http.Header {
	if r.header == nil {
		r.header = http.Header{}
	}

	return r.header
}

// Write's the data to the connection as part of an HTTP reply.
func (r *Response) Write(b []byte) (int, error) {
	if !r.flushed {
		r.WriteHeader(http.StatusOK)
	}

	return r.output.Write(b)
}

// WriteHeader sends an HTTP response header with status code.
func (r *Response) WriteHeader(c int) {
	if !r.flushed {
		r.status = c
		r.flushed = true
	}
}
