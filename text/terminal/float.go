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
	"strconv"

	"github.com/sniperkit/snk.fork.parsley/ast"
	"github.com/sniperkit/snk.fork.parsley/data"
	"github.com/sniperkit/snk.fork.parsley/parser"
	"github.com/sniperkit/snk.fork.parsley/parsley"
	"github.com/sniperkit/snk.fork.parsley/text"
)

// Float matches a float literal
func Float() *parser.NamedFunc {
	return parser.Func(func(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet) {
		tr := r.(*text.Reader)
		if readerPos, result := tr.ReadRegexp(pos, "[-+]?[0-9]*\\.[0-9]+(?:[eE][-+]?[0-9]+)?"); result != nil {
			val, err := strconv.ParseFloat(string(result), 64)
			if err != nil {
				return nil, parsley.NewErrorf(pos, "invalid float value encountered"), data.EmptyIntSet
			}
			return ast.NewTerminalNode("FLOAT", val, pos, readerPos), nil, data.EmptyIntSet
		}
		return nil, nil, data.EmptyIntSet
	}).WithName("float value")
}
