// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/julienschmidt/go-http-routing-benchmark/driver"
	"github.com/julienschmidt/go-http-routing-benchmark/driver/ui"

	_ "github.com/julienschmidt/go-http-routing-benchmark/router"
	_ "github.com/julienschmidt/go-http-routing-benchmark/suite"
)

const highlight = 3

var (
	std      = flag.Bool("std", false, "use default benchmark output")
	test     = flag.Bool("test", false, "only run tests and not benchmarks")
	list     = flag.Bool("list", false, "print the names of all registered Routers")
	markdown = flag.String("markdown", "", "write the benchmark output in markdown format to a file")
	color    = flag.String("color", "on", "colorize output, valid options are 'on', 'off' and 'ansi'")
	duration = flag.Duration("time", 5*time.Second, "approximate run time for each benchmark")
)

var (
	mdfile *os.File
	mdbuf  *bytes.Buffer
)

func init() {
	// beego sets it to runtime.NumCPU() and as of Go1.5 the default
	// value may change anyway. Currently none of the contesters does
	// concurrent routing.
	runtime.GOMAXPROCS(1)

	// Disable log output
	log.SetOutput(&driver.ResponseStub{})
}

func main() {
	flag.Parse()
	flag.Set("test.benchtime", duration.String())

	isvalid := false
	for _, op := range []string{"on", "off", "ansi"} {
		if *color == op {
			isvalid = true
		}
	}
	if !isvalid {
		fmt.Fprint(os.Stderr, "error: test: no packages registered\n")
		os.Exit(1)
	}

	if *list {
		fmt.Print(strings.Join(driver.PackageNames(), "\n") + "\n")
		os.Exit(0)
	}

	console := ui.Console(*color)
	console.Print("Running Tests...\n\n")
	tests := driver.RunTests()
	if tests != nil {
		RenderTest(console, tests)
		os.Exit(2)
	}

	if *test {
		os.Exit(0)
	}

	if *markdown != "" {
		f, err := os.Create(*markdown)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			os.Exit(1)
		}

		mdfile = f
		mdbuf = &bytes.Buffer{}
	}

	length := driver.NumPackages()
	renderer := &Renderer{
		span: &Span{
			Names:     make([]string, length),
			Homepages: make([]string, length),

			HeapMemory: &markUint64{marker: &marker{marked: make([]int, length)}},

			NsPerOp:     &markInt64{marker: &marker{marked: make([]int, length)}},
			AllocsPerOp: &markInt64{marker: &marker{marked: make([]int, length)}},
			BytesPerOp:  &markInt64{marker: &marker{marked: make([]int, length)}},
		},

		NsPerOp:     make([]int64, length),
		BytesPerOp:  make([]int64, length),
		AllocsPerOp: make([]int64, length),
	}

	br := &BenchmarkRenderer{
		mark:  console.Green,
		color: console.Yellow,
	}
	renderer.Add(br.Render)

	if *std {
		br.std = true
		br.mark = console.Print
	}

	console.Print("Running Benchmarks...\n\n")
	benchmarks := driver.RunBenchmarks()

	if *markdown != "" {
		mdbuf.WriteString("Results\n=======\n\n\n")
		RenderMemoryTable(benchmarks)
		renderer.Add(RenderMarkdown)
	}

	for _, result := range benchmarks {
		renderer.Render(result)
	}

	if mdfile != nil {
		// Create router name links
		urls := make(map[string]string)
		for _, result := range benchmarks {
			for _, pkg := range result.Packages() {
				urls[pkg.Name] = pkg.Homepage
			}
		}
		for name, url := range urls {
			if url != "" {
				fmt.Fprintf(mdbuf, "[%s]: %s\n", name, url)
			}
		}

		mdfile.Write(mdbuf.Bytes())
		mdfile.Close()
	}
}

func RenderMemoryTable(b []*driver.Result) {
	tcw := []int{4} // We'll add 2 later anyway
	thd := []string{"Router"}
	num := driver.NumPackages()
	rts := driver.PackageNames()
	col := make([]*markUint64, num)

	for i, name := range rts {
		col[i] = &markUint64{
			zero:   true,
			data:   make([]uint64, len(b)),
			marker: &marker{marked: make([]int, len(b))},
		}

		if tcw[0] < len(name) {
			tcw[0] = len(name)
		}
	}

	for _, r := range b {
		bname := r.Name()
		if !strings.HasSuffix(bname, "All") {
			continue
		}

		x := len(bname) - 3
		thd = append(thd, bname[:x])
		if x < 9 {
			x = 9
		}
		tcw = append(tcw, x)

		for i, name := range rts {
			pkgs := r.Packages()
			mems := r.RouterMemory()

			for x, pkg := range pkgs {
				if pkg.Name == name {
					col[i].data[len(thd)-2] = mems[x]
				}
			}
		}
	}

	fmt.Fprint(mdbuf, "### Memory Consumption\n\n\n")

	for _, m := range col {
		m.Set(m.data[:])
	}

	var header string
	for i, cell := range thd {
		tcw[i] += 2
		header += fmt.Sprintf("| %-*s ", tcw[i], cell)
	}
	header += "|\n"
	for i, x := range tcw {
		if i == 0 {
			header += fmt.Sprintf("|:%s-", strings.Repeat("-", x))
		} else {
			header += fmt.Sprintf("|-%s:", strings.Repeat("-", x))
		}
	}
	fmt.Fprint(mdbuf, header+"|\n")

	for i, name := range rts {
		fmt.Fprintf(mdbuf, "| %-*s ", tcw[0], "["+name+"]")

		for x, width := range tcw[1:] {
			var memory string
			if col[i].data[x] == 0 {
				memory = "-  "
			} else {
				if col[i].IsMarked(x) {
					memory = fmt.Sprintf("__%d__", col[i].data[x])
				} else {
					memory = fmt.Sprintf("%d  ", col[i].data[x])
				}
			}

			fmt.Fprintf(mdbuf, "| %*s", width+1, memory)
		}

		fmt.Fprint(mdbuf, "|\n")
	}

	fmt.Fprint(mdbuf, "\n\n")
}

type Renderer struct {
	span  *Span
	funcs []func(*driver.Result, *Span)

	// Slice cache
	NsPerOp     []int64
	BytesPerOp  []int64
	AllocsPerOp []int64
}

func (r *Renderer) Add(fn func(*driver.Result, *Span)) {
	r.funcs = append(r.funcs, fn)
}

func (r *Renderer) Render(result *driver.Result) {
	span := r.span
	pkgs := result.Packages()
	slen := len(pkgs)

	for i, pkg := range pkgs {
		span.Names[i] = pkg.Name
		span.Homepages[i] = pkg.Homepage
	}

	benchmarks := result.Results()
	for i, b := range benchmarks {
		r.NsPerOp[i] = b.NsPerOp()
		r.AllocsPerOp[i] = b.AllocsPerOp()
		r.BytesPerOp[i] = b.AllocedBytesPerOp()
	}

	span.length = slen
	span.NsPerOp.Set(r.NsPerOp[:slen])
	span.BytesPerOp.Set(r.BytesPerOp[:slen])
	span.AllocsPerOp.Set(r.AllocsPerOp[:slen])
	span.HeapMemory.Set(result.RouterMemory())

	for _, fn := range r.funcs {
		fn(result, span)
	}
}

type BenchmarkRenderer struct {
	std   bool
	mark  func(string)
	color func(string)
}

func (b *BenchmarkRenderer) Render(r *driver.Result, s *Span) {
	if !b.std {
		b.color(fmt.Sprintf("%s:\n%s\n\n", r.Name(), r.Description()))
	}

	packages := r.Packages()
	benchmarks := r.Results()

	colum := 0
	for _, pkg := range packages {
		if colum < len(pkg.Name) {
			colum = len(pkg.Name)
		}
	}

	if b.std {
		colum += 9 + len(r.Name())
	}

	for i := 0; i < s.length; i++ {
		var mwd int

		name := packages[i].Name
		if b.std {
			name = "Benchmark" + name + r.Name() + "\t"
		}

		if b.std {
			fmt.Printf("%-*s  %8d  ", colum, name, benchmarks[i].N)
		} else {
			fmt.Printf("%-*s  ", colum, name)

			if s.BytesPerOp.IsMarked(i) {
				b.mark(fmt.Sprintf("%8d", s.HeapMemory.data[i]))
			} else {
				fmt.Printf("%8d", s.HeapMemory.data[i])
			}
			fmt.Print(" B  ")
		}

		var nsopstr string
		nsop := s.NsPerOp.data[i]
		if benchmarks[i].N > 0 && nsop < 100 {
			if nsop < 10 {
				mwd = 3
				nsopstr = fmt.Sprintf("%13.2f", float64(benchmarks[i].T.Nanoseconds())/float64(benchmarks[i].N))
			} else {
				mwd = 2
				nsopstr = fmt.Sprintf("%12.1f", float64(benchmarks[i].T.Nanoseconds())/float64(benchmarks[i].N))
			}
		} else {
			nsopstr = fmt.Sprintf("%10d", nsop)
		}

		if s.NsPerOp.IsMarked(i) {
			b.mark(nsopstr)
		} else {
			fmt.Print(nsopstr)
		}
		fmt.Print(" ns/op  ")

		if s.BytesPerOp.IsMarked(i) {
			b.mark(fmt.Sprintf("%*d", 10-mwd, s.BytesPerOp.data[i]))
		} else {
			fmt.Printf("%*d", 10-mwd, s.BytesPerOp.data[i])
		}
		fmt.Print(" B/op  ")

		if s.AllocsPerOp.IsMarked(i) {
			b.mark(fmt.Sprintf("%8d", s.AllocsPerOp.data[i]))
		} else {
			fmt.Printf("%8d", s.AllocsPerOp.data[i])
		}
		fmt.Print(" allocs/op\n")
	}

	if !b.std {
		fmt.Print("\n\n")
	}
}

var widths = []int{9 /*Memory*/, 10 /*Iterations*/, 10 /*Ns*/, 8 /*B*/, 10 /*Allocs*/}

func RenderMarkdown(r *driver.Result, s *Span) {
	fmt.Fprintf(mdbuf, "#### %s\n\n%s\n\n", r.Name(), r.Description())

	packages := r.Packages()
	benchmarks := r.Results()

	c := 4 // We'll add 2 later anyway
	w := widths
	for _, pkg := range packages {
		if c < len(pkg.Name) {
			c = len(pkg.Name)
		}
	}
	c += 2

	fmt.Fprintf(mdbuf, "| %-*s | %-*s  | %-*s | %-*s  | %-*s  | %-*s  |\n",
		c, "Router", w[0], "Memory", w[1], "Iterations", w[2], "Ns/Op", w[3], "Bytes/Op", w[4], "Allocs/Op",
	)
	fmt.Fprintf(mdbuf, "|:%s-|-%s-:|-%s:|-%s-:|-%s-:|-%s-:|\n",
		strings.Repeat("-", c), strings.Repeat("-", w[0]), strings.Repeat("-", w[1]),
		strings.Repeat("-", w[2]), strings.Repeat("-", w[3]), strings.Repeat("-", w[4]),
	)

	for i := 0; i < s.length; i++ {

		var (
			nw     int
			name   string
			memory string
			nsop   string
			bop    string
			allocs string
		)

		if packages[i].Homepage != "" {
			nw = c
			name = "[" + packages[i].Name + "]"
		} else {
			nw = c + 2
			name = packages[i].Name
		}

		if s.BytesPerOp.IsMarked(i) {
			memory = fmt.Sprintf("__%d__", s.HeapMemory.data[i])
		} else {
			memory = fmt.Sprintf("%d  ", s.HeapMemory.data[i])
		}

		ns := s.NsPerOp.data[i]
		if benchmarks[i].N > 0 && ns < 100 {
			if ns < 10 {
				nsop = fmt.Sprintf("%.2f", float64(benchmarks[i].T.Nanoseconds())/float64(benchmarks[i].N))
			} else {
				nsop = fmt.Sprintf("%.1f", float64(benchmarks[i].T.Nanoseconds())/float64(benchmarks[i].N))
			}
		} else {
			nsop = fmt.Sprintf("%d", ns)
		}
		if s.NsPerOp.IsMarked(i) {
			nsop = fmt.Sprintf("__%s__", nsop)
		} else {
			nsop = fmt.Sprintf("%s  ", nsop)
		}

		if s.BytesPerOp.IsMarked(i) {
			bop = fmt.Sprintf("__%d__", s.BytesPerOp.data[i])
		} else {
			bop = fmt.Sprintf("%d  ", s.BytesPerOp.data[i])
		}

		if s.AllocsPerOp.IsMarked(i) {
			allocs = fmt.Sprintf("__%d__", s.AllocsPerOp.data[i])
		} else {
			allocs = fmt.Sprintf("%d  ", s.AllocsPerOp.data[i])
		}

		fmt.Fprintf(mdbuf, "| %-*s | %*s| %*d | %*s| %*s| %*s|\n",
			nw, name, w[0]+2, memory, w[1], benchmarks[i].N, w[2]+2, nsop, w[3]+2, bop, w[4]+2, allocs,
		)
	}

	mdbuf.WriteString("\n\n\n")
}

func RenderTest(ui ui.ConsoleAdapter, tests []*driver.Assertion) {
	// TODO: Improve dedup suppression implementation.
	var l *driver.Assertion

	for i, t := range tests {
		n := i + 1

		if l == nil || l.Benchmark() != t.Benchmark() {
			l = nil
			print("\n\n", t.Benchmark(), "\n\n", ui.Print, ui.Yellow)
		}

		if l == nil {
			print("Test:   ", t.Description(), "\n", ui.Print, ui.Print)

			if t.Error != "" {
				if len(tests) > n {
					if tests[n].Benchmark() != t.Benchmark() || tests[n].Error != t.Error {
						print("Error:  ", t.Error, "\n", ui.Print, ui.Red)
					}
				} else {
					print("Error:  ", t.Error, "\n", ui.Print, ui.Red)
				}

				print("Router: ", t.Router(), "\n", ui.Print, ui.Yellow)
			} else {
				print("Expect: ", t.Expect, "\n", ui.Print, ui.Green)
				print("Router: ", t.Router(), "\n", ui.Print, ui.Yellow)

				if len(tests) > n {
					if tests[n].Benchmark() != t.Benchmark() || tests[n].Actual != t.Actual {
						print("Actual: ", t.Actual, "\n", ui.Print, ui.Red)
					}
				} else {
					print("Actual: ", t.Actual, "\n", ui.Print, ui.Red)
				}
			}
		} else if len(tests) == n {
			if t.Error != "" {
				print("Error:  ", t.Error, "\n", ui.Print, ui.Red)
			} else {
				print("Actual: ", t.Actual, "\n", ui.Print, ui.Red)
			}
		} else {
			if l.Description() != t.Description() {
				print("Test:   ", t.Description(), "\n", ui.Print, ui.Print)
			}
			if l.Router() != t.Router() {
				print("Router: ", t.Router(), "\n", ui.Print, ui.Yellow)
			}
			if t.Error != "" && l.Error != t.Error {
				print("Error:  ", t.Error, "\n", ui.Print, ui.Red)
			}
			if l.Expect != t.Expect {
				print("Expect: ", t.Expect, "\n", ui.Print, ui.Green)
			}
			if l.Actual != t.Actual {
				print("Actual: ", t.Actual, "\n", ui.Print, ui.Red)
			}
		}

		l = t
	}
}

type Span struct {
	length int

	Names     []string // Package names
	Homepages []string // Package homepage URLs

	HeapMemory *markUint64 // Heap memory allocated for the router object

	NsPerOp     *markInt64 // testing.B NsPerOp
	AllocsPerOp *markInt64 // testing.B AllocsPerOp
	BytesPerOp  *markInt64 // testing.B AllocBytesPerOp
}

type markInt64 struct {
	*marker
	data []int64
}

func (m *markInt64) Less(i, j int) bool  { return m.data[i] < m.data[j] }
func (m *markInt64) cmp(i, j int) bool   { return m.data[i] != m.data[j] }
func (m *markInt64) IsMarked(i int) bool { return m.marker.isMarked(m.cmp, i) }

func (m *markInt64) Set(data []int64) {
	m.data = data
	m.marker.length = len(data)
	m.marker.set()
	sort.Sort(m)
}

type markUint64 struct {
	*marker
	zero bool // Ignore zero values
	data []uint64
}

func (m *markUint64) cmp(i, j int) bool   { return m.data[i] != m.data[j] }
func (m *markUint64) IsMarked(i int) bool { return m.marker.isMarked(m.cmp, i) }

func (m *markUint64) Less(i, j int) bool {
	if m.zero {
		return true
	}

	return m.data[i] < m.data[j]
}

func (m *markUint64) Set(data []uint64) {
	m.data = data
	m.marker.length = len(data)
	m.marker.set()
	sort.Sort(m)
}

type marker struct {
	number int
	length int
	marked []int
}

func (m *marker) Len() int      { return m.length }
func (m *marker) Swap(i, j int) { m.marked[i], m.marked[j] = m.marked[j], m.marked[i] }

func (m *marker) set() {
	m.number = -1

	if m.length > len(m.marked) {
		m.marked = make([]int, m.length)
	}

	for i := 0; i < m.length; i++ {
		m.marked[i] = i
	}
}

func (m *marker) isMarked(cmp func(int, int) bool, i int) bool {
	if m.number == -1 {
		m.markable(cmp, i)
	}

	for x := 0; x < m.number; x++ {
		if m.marked[x] == i {
			return true
		}
	}

	return false
}

func (m *marker) markable(cmp func(int, int) bool, i int) {
	m.number = highlight

	if m.length < highlight {
		m.number = int(m.length / 2)
	}

	if m.number <= 1 {
		if m.length > 1 && !cmp(0, 1) {
			m.number = 0
		}
		return
	}

	n := m.number
	c := m.number - 1

	for {
		ni := m.marked[n]
		ci := m.marked[c]

		if cmp(ci, ni) {
			break
		}

		n++
		c++
		m.number++
	}
}

func print(p, m, s string, w, c func(string)) {
	w(p)
	c(m)
	w(s)
}
