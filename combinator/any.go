// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package combinator

import (
	"github.com/opsidian/parsley/ast"
	"github.com/opsidian/parsley/data"
	"github.com/opsidian/parsley/parser"
	"github.com/opsidian/parsley/parsley"
)

// Any tries all the given parsers independently and merges the results
func Any(name string, parsers ...parsley.Parser) *parser.NamedFunc {
	if parsers == nil {
		panic("no parsers were given")
	}

	return parser.Func(func(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (data.IntSet, parsley.Node, parsley.Error) {
		cp := data.EmptyIntSet
		var res parsley.Node
		var err parsley.Error
		for _, p := range parsers {
			h.RegisterCall()
			cp2, res2, err2 := p.Parse(h, leftRecCtx, r, pos)
			cp = cp.Union(cp2)
			res = ast.AppendNode(res, res2)
			if err2 != nil && (err == nil || err2.Pos() >= err.Pos()) {
				err = err2
			}
		}
		return cp, res, err
	}).WithName(name)
}
