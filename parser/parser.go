/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package parser contains the main structs for parsing
package parser

import (
	"github.com/sniperkit/snk.fork.parsley/data"
	"github.com/sniperkit/snk.fork.parsley/parsley"
)

// Func defines a helper to implement the Parser interface with functions
type Func func(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet)

// Parse parses the input using the function
func (f Func) Parse(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet) {
	return f(h, leftRecCtx, r, pos)
}

// Name returns with an empty name
func (f Func) Name() string {
	return ""
}

// WithName returns with the same parser function but with the given name
// If a function is passed then it will be called when Name() is called
func (f Func) WithName(name interface{}) *NamedFunc {
	nf := &NamedFunc{
		f: f,
	}
	switch n := name.(type) {
	case string:
		nf.name = func() string { return n }
	case func() string:
		nf.name = n
	default:
		panic("name should be a string or a function returning string")
	}
	return nf
}

// NamedFunc is a parser function with a custom name
type NamedFunc struct {
	name func() string
	f    Func
}

// Parse parses the input using the function
func (nf *NamedFunc) Parse(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet) {
	return nf.f(h, leftRecCtx, r, pos)
}

// Name returns with the parser name
func (nf *NamedFunc) Name() string {
	return nf.name()
}
