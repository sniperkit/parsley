// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package parser contains the main structs for parsing
package parser

import (
	"github.com/opsidian/parsley/data"
	"github.com/opsidian/parsley/reader"
)

// Parser defines a parser interface
type Parser interface {
	Parse(leftRecCtx data.IntMap, r reader.Reader) (data.IntSet, ResultSet, Error)
}

// Func defines a helper to implement the Parser interface with functions
type Func func(leftRecCtx data.IntMap, r reader.Reader) (data.IntSet, ResultSet, Error)

// Parse parses the next token and returns with an AST node and the updated reader
func (f Func) Parse(leftRecCtx data.IntMap, r reader.Reader) (data.IntSet, ResultSet, Error) {
	return f(leftRecCtx, r)
}

// EmptyLeftRecCtx creates an empty left recursion context
func EmptyLeftRecCtx() data.IntMap {
	return data.EmptyIntMap()
}

// NoCurtailingParsers returns with an empty int set
func NoCurtailingParsers() data.IntSet {
	return data.EmptyIntSet()
}
