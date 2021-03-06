/*
Sniperkit-Bot
- Status: analyzed
*/

package json_test

import (
	encoding_json "encoding/json"
	"io/ioutil"
	"testing"

	"github.com/sniperkit/snk.fork.parsley/combinator"
	"github.com/sniperkit/snk.fork.parsley/examples/json/json"
	"github.com/sniperkit/snk.fork.parsley/parser"
	"github.com/sniperkit/snk.fork.parsley/parsley"
	"github.com/sniperkit/snk.fork.parsley/text"
)

func benchmarkParsleyJSON(b *testing.B, jsonFilePath string) {
	f, err := text.ReadFile(jsonFilePath)
	if err != nil {
		b.Fatal(err)
	}

	s := combinator.Sentence(json.NewParser())
	r := text.NewReader(f)
	h := parser.NewHistory()
	if _, err = parsley.Evaluate(h, r, s, nil); err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		h := parser.NewHistory()
		_, _ = parsley.Evaluate(h, r, s, nil)
	}
}

func BenchmarkParsleyJSON1k(b *testing.B)   { benchmarkParsleyJSON(b, "../example_1k.json") }
func BenchmarkParsleyJSON10k(b *testing.B)  { benchmarkParsleyJSON(b, "../example_10k.json") }
func BenchmarkParsleyJSON100k(b *testing.B) { benchmarkParsleyJSON(b, "../example_100k.json") }

func benchmarkEncodingJSON(b *testing.B, jsonFilePath string) {
	input, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		b.Fatal(err)
	}

	var val interface{}
	if err := encoding_json.Unmarshal(input, &val); err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		var val interface{}
		_ = encoding_json.Unmarshal(input, &val)
	}
}

func BenchmarkEncodingJSON1k(b *testing.B)   { benchmarkEncodingJSON(b, "../example_1k.json") }
func BenchmarkEncodingJSON10k(b *testing.B)  { benchmarkEncodingJSON(b, "../example_10k.json") }
func BenchmarkEncodingJSON100k(b *testing.B) { benchmarkEncodingJSON(b, "../example_100k.json") }
