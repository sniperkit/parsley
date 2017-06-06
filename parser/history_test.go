package parser_test

import (
	"testing"

	"github.com/opsidian/parsley/ast"
	"github.com/opsidian/parsley/data"
	"github.com/opsidian/parsley/parser"
	"github.com/opsidian/parsley/test"
	"github.com/stretchr/testify/assert"
)

func TestRegisterResultShouldSaveResultForPosition(t *testing.T) {
	h := parser.NewHistory()
	node := ast.NewTerminalNode("t", test.NewPosition(0), nil)
	cp := parser.NoCurtailingParsers()
	rs := parser.NewResult(node, nil).AsSet()
	h.RegisterResults(h.GetParserIndex("p1"), 2, cp, rs, parser.EmptyLeftRecCtx())

	actualCP, actualRS, ok := h.GetResults(h.GetParserIndex("p1"), 2, parser.EmptyLeftRecCtx())
	assert.Equal(t, cp, actualCP)
	assert.Equal(t, rs, actualRS)
	assert.True(t, ok)
}

func TestRegisterResultShouldReturnNilResult(t *testing.T) {
	h := parser.NewHistory()
	h.RegisterResults(h.GetParserIndex("p1"), 2, parser.NoCurtailingParsers(), nil, parser.EmptyLeftRecCtx())
	cp, rs, ok := h.GetResults(h.GetParserIndex("p1"), 2, parser.EmptyLeftRecCtx())
	assert.Equal(t, parser.NoCurtailingParsers(), cp)
	assert.Nil(t, rs)
	assert.True(t, ok)
}

func TestRegisterResultShouldReturnFalseWhenNoResultWasRegistered(t *testing.T) {
	h := parser.NewHistory()
	cp, rs, ok := h.GetResults(h.GetParserIndex("p1"), 2, parser.EmptyLeftRecCtx())
	assert.Equal(t, parser.NoCurtailingParsers(), cp)
	assert.Nil(t, rs)
	assert.False(t, ok)
}

func TestRegisterResultShouldHandleMultipleParsers(t *testing.T) {
	h := parser.NewHistory()
	node := ast.NewTerminalNode("t", test.NewPosition(0), nil)
	cp1 := parser.NoCurtailingParsers()
	cp2 := data.NewIntSet(1)
	rs1 := parser.NewResult(node, nil).AsSet()
	var rs2 parser.ResultSet
	h.RegisterResults(h.GetParserIndex("p1"), 1, cp1, rs1, parser.EmptyLeftRecCtx())
	h.RegisterResults(h.GetParserIndex("p2"), 2, cp2, rs2, parser.EmptyLeftRecCtx())

	actualCP, actualRS, ok := h.GetResults(h.GetParserIndex("p1"), 1, parser.EmptyLeftRecCtx())
	assert.Equal(t, cp1, actualCP)
	assert.Equal(t, rs1, actualRS)
	assert.True(t, ok)

	actualCP, actualRS, ok = h.GetResults(h.GetParserIndex("p2"), 2, parser.EmptyLeftRecCtx())
	assert.Equal(t, cp2, actualCP)
	assert.Equal(t, rs2, actualRS)
	assert.True(t, ok)
}

func TestGetResultsShouldNotReturnCurtailedResult(t *testing.T) {
	h := parser.NewHistory()
	ctx := data.NewIntMap(map[int]int{
		h.GetParserIndex("p1"): 2,
		h.GetParserIndex("p2"): 1,
	})
	cp := data.NewIntSet(h.GetParserIndex("p1"))
	h.RegisterResults(h.GetParserIndex("p1"), 1, cp, nil, ctx)

	ctx = data.NewIntMap(map[int]int{
		h.GetParserIndex("p1"): 1,
		h.GetParserIndex("p2"): 1,
	})
	cp, rs, found := h.GetResults(h.GetParserIndex("p1"), 1, ctx)
	assert.Equal(t, parser.NoCurtailingParsers(), cp)
	assert.Nil(t, rs)
	assert.False(t, found)
}

func TestGetResultsShouldReturnCurtailedResult(t *testing.T) {
	h := parser.NewHistory()
	ctx := data.NewIntMap(map[int]int{
		h.GetParserIndex("p1"): 2,
		h.GetParserIndex("p2"): 1,
	})
	cp := data.NewIntSet(h.GetParserIndex("p1"))
	rs := parser.NewResult(nil, nil).AsSet()
	h.RegisterResults(h.GetParserIndex("p1"), 1, cp, rs, ctx)

	ctx = data.NewIntMap(map[int]int{
		h.GetParserIndex("p1"): 1,
		h.GetParserIndex("p2"): 1,
	})

	ctx = ctx.Inc(h.GetParserIndex("p1"))
	actualCP, actualRS, found := h.GetResults(h.GetParserIndex("p1"), 1, ctx)
	assert.Equal(t, cp, actualCP)
	assert.Equal(t, rs, actualRS)
	assert.True(t, found)
}