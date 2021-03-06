/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package terminal_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/sniperkit/snk.fork.parsley/ast"
	"github.com/sniperkit/snk.fork.parsley/data"
	"github.com/sniperkit/snk.fork.parsley/parsley"
	"github.com/sniperkit/snk.fork.parsley/text"
	"github.com/sniperkit/snk.fork.parsley/text/terminal"
)

var _ = Describe("Whitespaces", func() {

	Context("when not allowing any whitespaces", func() {

		var p = terminal.Whitespaces(text.WsNone)

		It("should always return with an empty node", func() {
			r := text.NewReader(text.NewFile("textfile", []byte("abc")))
			res, err, curtailingParsers := p.Parse(nil, data.EmptyIntMap, r, 0)
			Expect(curtailingParsers).To(Equal(data.EmptyIntSet))
			Expect(res).To(Equal(ast.NilNode(0)))
			Expect(err).ToNot(HaveOccurred())
		})

		It("should not match whitespaces", func() {
			r := text.NewReader(text.NewFile("textfile", []byte(" abc")))
			res, err, curtailingParsers := p.Parse(nil, data.EmptyIntMap, r, 0)
			Expect(curtailingParsers).To(Equal(data.EmptyIntSet))
			Expect(res).To(Equal(ast.NilNode(0)))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when not allowing new lines", func() {

		var p = terminal.Whitespaces(text.WsSpaces)

		DescribeTable("should match",
			func(input string, startPos int, nodePos parsley.Pos, endPos int) {
				f := text.NewFile("textfile", []byte(input))
				r := text.NewReader(f)
				res, err, curtailingParsers := p.Parse(nil, data.EmptyIntMap, r, f.Pos(startPos))
				Expect(curtailingParsers).To(Equal(data.EmptyIntSet))
				Expect(err).ToNot(HaveOccurred())
				node := res.(ast.NilNode)
				Expect(node.Token()).To(Equal("NIL"))
				Expect(node.Value(nil)).To(BeNil())
				Expect(node.Pos()).To(Equal(f.Pos(endPos)))
				Expect(node.ReaderPos()).To(Equal(f.Pos(endPos)))
			},
			Entry("ws beginning", " \t---", 0, parsley.Pos(1), 2),
			Entry("ws middle", "--- \t---", 3, parsley.Pos(4), 5),
			Entry("ws end", "--- \t", 3, parsley.Pos(4), 5),
			Entry("should not match new line", " \t\n\f", 0, parsley.Pos(1), 2),
		)

		DescribeTable("should not match",
			func(input string, startPos int) {
				f := text.NewFile("textfile", []byte(input))
				r := text.NewReader(f)
				res, err, curtailingParsers := p.Parse(nil, data.EmptyIntMap, r, f.Pos(startPos))
				Expect(curtailingParsers).To(Equal(data.EmptyIntSet))
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(BeNil())
			},
			Entry("empty", "", 0),
			Entry("not whitespace", `a`, 0),
			Entry("new line", `\n\t`, 0),
		)
	})

	Context("when allowing new lines", func() {

		var p = terminal.Whitespaces(text.WsSpacesNl)

		DescribeTable("should match (with new lines)",
			func(input string, startPos int, nodePos parsley.Pos, endPos int) {
				f := text.NewFile("textfile", []byte(input))
				r := text.NewReader(f)
				res, err, curtailingParsers := p.Parse(nil, data.EmptyIntMap, r, f.Pos(startPos))
				Expect(curtailingParsers).To(Equal(data.EmptyIntSet))
				Expect(err).ToNot(HaveOccurred())
				node := res.(ast.NilNode)
				Expect(node.Token()).To(Equal("NIL"))
				Expect(node.Value(nil)).To(BeNil())
				Expect(node.Pos()).To(Equal(f.Pos(endPos)))
				Expect(node.ReaderPos()).To(Equal(f.Pos(endPos)))
			},
			Entry("ws beginning", " \t\n\f---", 0, parsley.Pos(1), 4),
			Entry("ws middle", "--- \t\n\f---", 3, parsley.Pos(4), 7),
			Entry("ws end", "--- \t\n\f", 3, parsley.Pos(4), 7),
		)

		DescribeTable("should not match (with new lines)",
			func(input string, startPos int) {
				f := text.NewFile("textfile", []byte(input))
				r := text.NewReader(f)
				res, err, curtailingParsers := p.Parse(nil, data.EmptyIntMap, r, f.Pos(startPos))
				Expect(curtailingParsers).To(Equal(data.EmptyIntSet))
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(BeNil())
			},
			Entry("empty", "", 0),
			Entry("not whitespace", `a`, 0),
		)
	})

})
