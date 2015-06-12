// Copyright 2014-2015 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package ui

import (
	"fmt"
	"runtime"
	"strings"
)

var defaultConsole ConsoleAdapter

// Console sets the default console and returns it.
func Console(style string) ConsoleAdapter {
	var stylizer Colorizer

	switch style {
	case "on":
		stylizer = color{}
	case "ansi":
		stylizer = ansiColor{}
	case "off":
		stylizer = nullColor{}
	default:
		panic("Unknown color option")
	}

	switch runtime.GOOS {
	case "windows":
		defaultConsole = windowsConsole{baseConsole{stylizer}}
	default:
		defaultConsole = posixConsole{baseConsole{stylizer}}
	}

	return defaultConsole
}

type ConsoleAdapter interface {
	Reset()
	Width() int
	Height() int
	Red(string)
	Green(string)
	Yellow(string)
	Print(string)
	Printf(string, ...interface{})
	Color() Colorizer
}

type baseConsole struct {
	style Colorizer
}

func (c baseConsole) Width() int                           { return 80 }
func (c baseConsole) Height() int                          { return 25 }
func (c baseConsole) Color() Colorizer                     { return c.style }
func (c baseConsole) Reset()                               { fmt.Print("\f") }
func (c baseConsole) Red(b string)                         { fmt.Print(c.style.Red(b)) }
func (c baseConsole) Green(b string)                       { fmt.Print(c.style.Green(b)) }
func (c baseConsole) Yellow(b string)                      { fmt.Print(c.style.Yellow(b)) }
func (c baseConsole) Print(p string)                       { fmt.Print(p) }
func (c baseConsole) Printf(f string, args ...interface{}) { fmt.Printf(f, args...) }

type posixConsole struct {
	baseConsole
}

type windowsConsole struct {
	baseConsole
}

func (c windowsConsole) Reset() { fmt.Print(strings.Repeat("\r\n", c.Height())) }

type Colorizer interface {
	Red(string) string
	Green(string) string
	Yellow(string) string
}

type nullColor struct{}

func (c nullColor) Red(b string) string    { return b }
func (c nullColor) Green(b string) string  { return b }
func (c nullColor) Yellow(b string) string { return b }

type color struct{}

func (c color) Red(b string) string {
	return "\033[1;31m" + b + "\033[0m"
}

func (c color) Green(b string) string {
	return "\033[1;32m" + b + "\033[0m"
}

func (c color) Yellow(b string) string {
	return "\033[1;33m" + b + "\033[0m"
}

type ansiColor struct{}

func (c ansiColor) Red(b string) string {
	return "\x1b[1;31m" + b + "\x1b[0m"
}

func (c ansiColor) Green(b string) string {
	return "\x1b[1;32m" + b + "\x1b[0m"
}

func (c ansiColor) Yellow(b string) string {
	return "\x1b[1;33m" + b + "\x1b[0m"
}
