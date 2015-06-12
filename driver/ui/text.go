// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package ui

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

// Text describes a text runner user interface.
type Text struct {
	errors  int
	skipped int

	max    int // Maximum steps to take
	step   int // Steps already taken
	colum  int // Right colum width
	width  int // Current width without the colum
	digits int // Digits of max steps

	start   time.Time
	line    bytes.Buffer
	console ConsoleAdapter
}

// NewText returns a new text user interface which can be called steps times.
// After steps calls it prints stats. This function captures the default colorizer
// at instantiation.
func NewText(steps int) *Text {
	d := len(strconv.Itoa(steps))

	return &Text{
		digits:  d,
		max:     steps,
		start:   time.Now(),
		console: defaultConsole,
		colum:   6 + d + 1 + d, // " 000% " d "/" d,
	}
}

// Success marks the current step as successful.
func (t *Text) Success() {
	t.write(".")
}

// Error marks the current step as failed.
func (t *Text) Error() {
	t.errors++
	t.write(t.console.Color().Red("E"))
}

// Skip marks the current step as skipped.
func (t *Text) Skip() {
	t.skipped++
	t.write(t.console.Color().Yellow("S"))
}

func (t *Text) write(b string) {
	t.step++
	t.width++
	t.line.WriteString(b)

	pd := strings.Repeat(" ", t.console.Width()-(t.width+t.colum))

	if t.max < t.step {
		panic("Maximum steps already reached")
	}

	pc := int(100 * float64(t.step) / float64(t.max))
	t.console.Printf("\r%s%s %3d%% %*d/%*d", t.line.String(), pd, pc, t.digits, t.step, t.digits, t.max)

	if (t.width + t.colum) >= t.console.Width() {
		t.width = 0
		t.line.Reset()
		t.console.Print("\n")
	}

	if t.max == t.step {
		t.console.Print("\n\n")

		t.console.Printf("Total: %d", t.max)
		if t.errors > 0 {
			t.console.Printf(", %d failed", t.errors)
		}
		if t.skipped > 0 {
			t.console.Printf(", %d skipped", t.skipped)
		}

		d := time.Since(t.start)
		h := int(d.Hours())
		m := int(d.Minutes()) % 60
		s := int(d.Seconds()) % 60

		t.console.Print("; Time: ")
		if h > 0 {
			t.console.Printf("%d:", h)
		}
		t.console.Printf("%02d:%02d", m, s)

		t.console.Print("\n\n\n")
	}
}
