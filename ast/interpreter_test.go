package ast_test

import (
	"errors"
	"testing"

	"github.com/opsidian/parsley/ast"
	"github.com/opsidian/parsley/test"
	"github.com/stretchr/testify/assert"
)

func TesInterpreterFuncShouldCallFunction(t *testing.T) {
	ctx := "textCtx"
	nodes := []ast.Node{
		ast.NewTerminalNode("INT", test.NewPosition(0), 1),
	}
	expectedValue := 1
	expectedErr := errors.New("e")
	var actualCtx interface{}
	var actualNodes []ast.Node
	interpreterFunc := ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, error) {
		actualCtx = ctx
		actualNodes = nodes
		return expectedValue, expectedErr
	})

	actualValue, actualErr := interpreterFunc.Eval(ctx, nodes)

	assert.Equal(t, ctx, actualCtx)
	assert.Equal(t, nodes, actualNodes)
	assert.Equal(t, expectedErr, actualErr)
	assert.Equal(t, expectedValue, actualValue)
}
