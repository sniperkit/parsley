/*
Sniperkit-Bot
- Status: analyzed
*/

package parsley_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sniperkit/snk.fork.parsley/parsley"
)

var _ = Describe("NilPosition", func() {
	It("implements the Positon interface", func() {
		var _ parsley.Position = parsley.NilPosition
	})

	It("returns with a non-empty string representation", func() {
		Expect(parsley.NilPosition.String()).ToNot(BeEmpty())
	})
})
