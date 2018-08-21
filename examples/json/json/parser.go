/*
Sniperkit-Bot
- Status: analyzed
*/

package json

import (
	"github.com/sniperkit/snk.fork.parsley/ast/interpreter"
	"github.com/sniperkit/snk.fork.parsley/combinator"
	"github.com/sniperkit/snk.fork.parsley/parser"
	"github.com/sniperkit/snk.fork.parsley/text"
	"github.com/sniperkit/snk.fork.parsley/text/terminal"
)

// NewParser returns with a new JSON parser
func NewParser() *parser.NamedFunc {
	var value parser.NamedFunc

	array := combinator.Seq("ARRAY", "array",
		terminal.Rune('['),
		combinator.SepBy(
			text.LeftTrim(&value, text.WsSpacesNl),
			text.LeftTrim(terminal.Rune(','), text.WsSpaces),
		).Bind(interpreter.Array()),
		text.LeftTrim(terminal.Rune(']'), text.WsSpacesNl),
	).Bind(interpreter.Select(1))

	keyValue := combinator.Seq("OBJ_KV", "key-value pair",
		terminal.String(false),
		text.LeftTrim(terminal.Rune(':'), text.WsSpaces),
		text.LeftTrim(&value, text.WsSpaces),
	)

	object := combinator.Seq("OBJ", "object",
		terminal.Rune('{'),
		combinator.SepBy(
			text.LeftTrim(keyValue, text.WsSpacesNl),
			text.LeftTrim(terminal.Rune(','), text.WsSpaces),
		).Bind(interpreter.Object()),
		text.LeftTrim(terminal.Rune('}'), text.WsSpacesNl),
	).Bind(interpreter.Select(1))

	value = *combinator.Choice("value",
		terminal.String(false),
		terminal.Float(),
		terminal.Integer(),
		array,
		object,
		terminal.Word("false", false),
		terminal.Word("true", true),
		terminal.Word("null", nil),
	)

	return text.Trim(&value)
}
