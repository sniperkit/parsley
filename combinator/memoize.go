package combinator

import (
	"github.com/opsidian/parsley/data"
	"github.com/opsidian/parsley/parser"
	"github.com/opsidian/parsley/reader"
)

// Memoize handles result cache and curtailing left recursion
func Memoize(name string, h *parser.History, p parser.Parser) parser.Func {
	parserIndex := h.GetParserIndex(name)
	return parser.Func(func(leftRecCtx data.IntMap, r reader.Reader) (data.IntSet, parser.ResultSet, parser.Error) {
		cp, rs, err, found := h.GetResults(parserIndex, r.Cursor().Pos(), leftRecCtx)
		if found {
			return cp, rs, err
		}

		if leftRecCtx.Get(parserIndex) > r.Remaining()+1 {
			return data.NewIntSet(parserIndex), nil, nil
		}

		cp, rs, err = p.Parse(leftRecCtx.Inc(parserIndex), r)
		leftRecCtx = leftRecCtx.Filter(cp)

		h.RegisterResults(parserIndex, r.Cursor().Pos(), cp, rs, err, leftRecCtx)

		return cp, rs, err
	})
}
