// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/opsidian/parsley/ast"
	"github.com/opsidian/parsley/data"
	"github.com/opsidian/parsley/parsley"
)

// Empty always matches and returns with an empty node result
// When using Empty you should not forget to handle for nil nodes in your node builders and/or interpreters.
func Empty() Func {
	return Func(func(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet) {
		return ast.EmptyNode(pos), nil, data.EmptyIntSet
	})
}

// End matches the end of the input
func End() *NamedFunc {
	return Func(func(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet) {
		if r.IsEOF(pos) {
			return ast.NewTerminalNode(ast.EOF, nil, pos, pos), nil, data.EmptyIntSet
		}
		return nil, nil, data.EmptyIntSet
	}).WithName("the end of input")
}
