package combinator

import (
	"github.com/opsidian/parsley/ast"
	"github.com/opsidian/parsley/ast/builder"
	"github.com/opsidian/parsley/parser"
)

// ManySepBy matches the given value parser zero or more times separated by the separator parser
func ManySepBy(name string, token string, h *parser.History, valueP parser.Parser, sepP parser.Parser, min int, interpreter ast.Interpreter) parser.Func {
	sepValue := Memoize(name+"_MSB", h, And(builder.All("SEP_VALUE", interpreter), sepP, valueP))
	sepValueMany := Memoize(name+"_MSB*", h, Many(builder.Flatten(token, interpreter), sepValue, 0, -1))
	return Try(mergeChildren(token, interpreter), min, valueP, sepValueMany)
}

func mergeChildren(token string, interpreter ast.Interpreter) ast.NodeBuilder {
	return ast.NodeBuilderFunc(func(nodes []ast.Node) ast.Node {
		children := []ast.Node{nodes[0]}
		if len(nodes) > 1 {
			node1 := nodes[1].(ast.NonTerminalNode)
			children = append(children, node1.Children()...)
		}
		return ast.NewNonTerminalNode(token, children, interpreter)
	})
}
