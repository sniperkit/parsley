/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package terminal

import (
	"github.com/sniperkit/snk.fork.parsley/ast"
	"github.com/sniperkit/snk.fork.parsley/data"
	"github.com/sniperkit/snk.fork.parsley/parser"
	"github.com/sniperkit/snk.fork.parsley/parsley"
	"github.com/sniperkit/snk.fork.parsley/text"
)

// Whitespaces matches one or more spaces or tabs. If newLine is true it also matches \n and \f characters.
func Whitespaces(wsMode text.WsMode) parsley.Parser {
	if wsMode == text.WsNone {
		return parser.Nil()
	}
	var name string
	if wsMode == text.WsSpaces {
		name = "spaces or tabs"
	} else {
		name = "spaces, tabs or newline"
	}
	return parser.Func(func(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet) {
		tr := r.(*text.Reader)
		if readerPos := tr.SkipWhitespaces(pos, wsMode); readerPos > pos {
			return ast.NilNode(readerPos), nil, data.EmptyIntSet
		}

		return nil, nil, data.EmptyIntSet
	}).WithName(name)
}
