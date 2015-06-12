// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package driver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/julienschmidt/go-http-routing-benchmark/driver/ui"
)

// RunTests runs all registerd tests against all registerd packages.
func RunTests() []*Assertion {
	pn, bn := len(packages), len(benchmarks)

	if pn == 0 {
		fmt.Fprint(os.Stderr, "error: test: no packages registered\n")
		os.Exit(1)
	}
	if bn == 0 {
		fmt.Fprint(os.Stderr, "error: test: no benchmarks registered\n")
		os.Exit(1)
	}

	var runs int
	var errs []*Assertion

	// Calculate the number of tests we have to run
	for _, b := range benchmarks {
		var rs, qs, ps int

		for _, t := range b.tests {
			if t.Route != nil {
				rs += len(b.routes)
			}
			if t.Request != nil {
				qs += len(b.requests)
			}
		}
		for _, p := range packages {
			if p.Supports.isset(b.Type) {
				ps++
			}
		}

		runs += ((rs + qs) * ps) + (len(packages) - ps)
	}

	rw := &Response{}
	ui := ui.NewText(runs)

	defer func() {
		if rerr := recover(); rerr != nil {
			ui.Error()
			stack := make([]byte, 10000)
			runtime.Stack(stack, false)
			fmt.Fprintf(os.Stderr, "\n\n%s\n\n%s", rerr, stack)
			os.Exit(2)
		}
	}()

	for _, b := range benchmarks {
		for _, p := range packages {
			if !p.Supports.isset(b.Type) {
				ui.Skip()
				continue
			}

			h := p.Router(b.Type, normalize(b.routes, p.Normalizer))

			for _, t := range b.tests {
				rt := t.Route
				if rt != nil {
					for _, r := range b.routes {
						x := verify(rt(r))

						if x != nil {
							ui.Error()
							x.pkg = p
							x.test = t
							x.bench = b
							errs = append(errs, x)
						} else {
							ui.Success()
						}
					}
				}

				qt := t.Request
				if qt != nil {
					for _, r := range b.requests {
						rw.Reset()
						h.ServeHTTP(rw, r)
						x := verify(qt(r, rw))

						if x != nil {
							if x.Error == "" && x.Expect == x.Actual {
								ui.Success()
								continue
							}

							ui.Error()
							x.pkg = p
							x.test = t
							x.bench = b
							errs = append(errs, x)
						} else {
							ui.Success()
						}
					}
				}
			}
		}
	}

	return errs
}

// RunBenchmarks runs all registerd benchmarks against all registerd packages.
func RunBenchmarks() []*Result {
	pn, bn := len(packages), len(benchmarks)

	if pn == 0 {
		fmt.Fprint(os.Stderr, "error: benchmark: no packages registered\n")
		os.Exit(1)
	}
	if bn == 0 {
		fmt.Fprint(os.Stderr, "error: benchmark: no benchmarks registered\n")
		os.Exit(1)
	}

	var (
		results []*Result
		memory  runtime.MemStats
	)

	ui := ui.NewText(pn * bn)

	for _, b := range benchmarks {
		r := &Result{
			benchmark: b,
		}

		for _, p := range packages {
			if !p.Supports.isset(b.Type) {
				ui.Skip()
				continue
			}

			f := normalize(b.routes, p.Normalizer)

			runtime.GC()
			runtime.ReadMemStats(&memory)
			m := memory.HeapAlloc

			h := p.Router(b.Type, f)

			runtime.GC()
			runtime.ReadMemStats(&memory)
			m = memory.HeapAlloc - m

			x := testing.Benchmark(benchmarkify(h, b.requests))

			r.pkgs = append(r.pkgs, p)
			r.alloc = append(r.alloc, m)
			r.results = append(r.results, &x)

			ui.Success()
		}

		if r.pkgs != nil && r.alloc != nil && r.results != nil {
			results = append(results, r)
		}

	}

	return results
}

// benchmarkify wraps a http.Handler in a benchmarkable closure.
func benchmarkify(h http.Handler, rs []*http.Request) func(*testing.B) {
	rw := &ResponseStub{}

	if len(rs) > 1 {
		r := rs[0]

		return func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				h.ServeHTTP(rw, r)
			}
		}
	}

	return func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			for _, r := range rs {
				h.ServeHTTP(rw, r)
			}
		}
	}
}

var ncache = make(map[nkey]Fixtures)

type nkey struct {
	f *Fixtures
	n *Normalizer
}

// normalize's and caches normalized Fixtures.
func normalize(f Fixtures, n Normalizer) Fixtures {
	if n == nil {
		return f
	}

	k := nkey{&f, &n}
	x, ok := ncache[k]

	if ok {
		return x
	}

	o := n(f)
	ncache[k] = o

	return o
}

// verify checks that the Assertion is valid or nil. A Assertion is only valid
// if Error is set and Expect and Actual are empty or vice versa. See Assertion.
func verify(a *Assertion) *Assertion {
	if a == nil {
		return a
	}

	var errors []string

	if a.Error != "" {
		if a.Expect != "" {
			panic("Expect must be empty if Error is set")
		}
		if a.Actual != "" {
			errors = append(errors, "Actual must be empty if Error is set")
		}
	} else {
		if a.Expect == "" {
			errors = append(errors, "Expect must be set")
		}
		if a.Actual == "" {
			errors = append(errors, "Actual must be set")
		}
	}

	if errors != nil {
		panic(strings.Join(errors, ", "))
	}

	return a
}

type ResponseStub struct{}

func (r *ResponseStub) Header() http.Header {
	return http.Header{}
}

func (r *ResponseStub) Write(p []byte) (int, error) {
	return len(p), nil
}

func (r *ResponseStub) WriteString(p string) (int, error) {
	return len(p), nil
}

func (r *ResponseStub) WriteHeader(int) {}

var LoggerStub *log.Logger

func init() {
	LoggerStub = log.New(&ResponseStub{}, "", 0)
}
