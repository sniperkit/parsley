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
	"fmt"

	"github.com/sniperkit/snk.fork.parsley/ast"
	"github.com/sniperkit/snk.fork.parsley/data"
	"github.com/sniperkit/snk.fork.parsley/parser"
	"github.com/sniperkit/snk.fork.parsley/parsley"
	"github.com/sniperkit/snk.fork.parsley/text"
)

// Word matches the given word
func Word(word string, value interface{}) *parser.NamedFunc {
	if word == "" {
		panic("Word() should not be called with empty word")
	}

	return parser.Func(func(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet) {
		tr := r.(*text.Reader)
		if readerPos, found := tr.MatchWord(pos, word); found {
			return ast.NewTerminalNode("WORD", value, pos, readerPos), nil, data.EmptyIntSet
		}
		return nil, nil, data.EmptyIntSet
	}).WithName(fmt.Sprintf("%q", word))
}
