## 0.7.0

BACKWARDS INCOMPATIBILITIES:
* extend parsley.Error to have a `Cause() error` method
* rename parsley.NewError() to parsley.NewErrorf()
* add parsley.NewError with error type as input
* parsley.WrapError will store the original cause but update the error message

## 0.6.0

BACKWARDS INCOMPATIBILITIES:
* major refactor of most of the API
* most of the interfaces were moved to the parsley package
* the position handling was rewritten similar to go's token.Pos
* the builders were completely removed
* the reader doesn't handle whitespaces anymore
* the parsers' return value were simplified
* nil node values are not allowed anymore
* most of the combinators API's were simplified to avoid repetition (like name + token)

IMPROVEMENTS:
* the reader became stateless
* most of the tests were rewritten using Ginkgo/Gomega
* common interpreters were added (array, object)
* whitespaces can be handled precisely with new parsers (text.LeftTrim/RightTrim/Trim)
* new error type with position (parsley.Pos)
* new empty node type
* new file and fileset types were introduced to support parsing multiple files better

BUGFIXES:
* the History was using the wrong key when checking left-recursion and wasn't curtailing properly

TODO:
* some of the old tests in the combinator package needs to be rewritten using Gingko (this means we miss a lot of test coverage)

## 0.5.0

BACKWARDS INCOMPATIBILITIES:
* text.NewReader now expects a filename parameter

IMPROVEMENTS:
* Windows-style line endings (\r\n) are automatically replaced to Unix-style line endings (\n) in the text reader.

## 0.4.0

BACKWARDS INCOMPATIBILITIES:
* Typo fix in some methods: peak -> peek

IMPROVEMENTS:
* Move precompiled whitespace regexp to a separate variable in text reader

OTHER:
* Fix example JSON parser + add comparison benchmarks against encoding/json

## 0.3.3

IMPROVEMENTS:

* allow nil pos in reader errors, replace {{err}} placeholder in WrapError
* add filename to position, add new file reader constructor

## 0.3.2

IMPROVEMENTS:

* reader.WrapError() falls back to the cause's error message if no error message was given and the cause is not a reader error

## 0.3.1

IMPROVEMENTS:
* Add SepByOrValue and SepByOrValue1 combinators which will return the value node if only the value parser is matched

## 0.3.0

BACKWARDS INCOMPATIBILITIES:
* Add history object to parser.Parse
* Move Memoize back to the combinators
* Move parsley package to a directory
* Change top-level Parse/Evaluate methods
* Remove parser.Stat, collect call statistics in history

IMPROVEMENTS:
* Add Sentence root parser

## 0.2.3

IMPROVEMENTS:
* reader.WrapError keeps the original error message if empty string is given as message

## 0.2.2

IMPROVEMENTS:
* Add parser.FuncFactory interface

## 0.2.1

BUGFIXES:
* the ast.Node mock wasn't regenerated

IMPROVEMENTS:
* Generate mock for parser.Parser

## 0.2.0

BACKWARDS INCOMPATIBILITIES:
* combinator.Memoize was removed. Use the Memoize method on the history object instead.
* parser.Error, parser.NewError and parser.WrapError was moved to reader
* Interpreter.Eval now returns a reader.Error instead of a general error
* Node.Value now returns a reader.Error instead of a general error
* Terminal.Value now returns a reader.Error instead of a general error
* NonTerminal.Value now returns a reader.Error instead of a general error

## 0.1.5

IMPROVEMENTS:

* Add cause to reader.Error, add parser.WrapError constructor

## 0.1.4

IMPROVEMENTS:

* All combinators will return with a parser.Func type (not the parser.Parser interface)
* Change internal sepBy to a factory
* Define dependencies with Go Dep
* Generate mocks for testing

## 0.1.3

DEPRECATED:

* combinator.Memoize will be removed in version 0.2. Use the Memoize method on the history object instead.

IMPROVEMENTS:

* You don't need to use meaningless names for memoization anymore, but you have to be careful to call Memoize
  for your parsers only once.

CHANGES:

* History.GetParserIndex was removed as only the original combinator.Memoize needed it

## 0.1.2

IMPROVEMENTS:

* Add copyright and license headers to all .go files
* Improve code coverage, add notes for (hopefully) impossible panics
* Add codecov.io integration

## 0.1.1

BUG FIXES:

* IsEOF() in text.Reader was not ignoring whitespaces when ignoreWhitespaces was set to true.

IMPROVEMENTS:

* Add Reset() method to History to allow to reuse the parsers for multiple inputs

## 0.1.0

First release
