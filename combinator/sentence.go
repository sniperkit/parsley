/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package combinator

import (
	"github.com/sniperkit/snk.fork.parsley/ast/interpreter"
	"github.com/sniperkit/snk.fork.parsley/parser"
	"github.com/sniperkit/snk.fork.parsley/parsley"
)

// Sentence matches the given parser until the end of input
func Sentence(p parsley.Parser) *Recursive {
	return Seq("S", "", p, parser.End()).Bind(interpreter.Select(0))
}
