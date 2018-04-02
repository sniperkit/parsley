package parsley_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opsidian/parsley/parsley"
	"github.com/opsidian/parsley/parsley/parsleyfakes"
)

var _ = Describe("File set", func() {
	var (
		files []parsley.File
		fs    *parsley.FileSet
	)

	BeforeEach(func() {
		files = []parsley.File{}
	})

	JustBeforeEach(func() {
		fs = parsley.NewFileSet(files...)
	})

	Describe("NewFile()", func() {
		It("returns with a non-nil file set", func() {
			Expect(fs).ToNot(BeNil())
		})

		Context("when created with files", func() {
			var f *parsleyfakes.FakeFile

			BeforeEach(func() {
				f = &parsleyfakes.FakeFile{}
				f.LenReturns(10)
				files = []parsley.File{f}
			})

			It("sets the offset for the file", func() {
				Expect(f.SetOffsetCallCount()).To(Equal(1))
				offset := f.SetOffsetArgsForCall(0)
				Expect(offset).To(Equal(1))
			})
		})
	})

	Describe("AddFile()", func() {
		It("sets the offset for the file", func() {
			f := &parsleyfakes.FakeFile{}
			fs.AddFile(f)
			Expect(f.SetOffsetCallCount()).To(Equal(1))
			offset := f.SetOffsetArgsForCall(0)
			Expect(offset).To(Equal(1))
		})

		Context("when you add an additional file", func() {
			BeforeEach(func() {
				f := &parsleyfakes.FakeFile{}
				f.LenReturns(10)
				files = []parsley.File{f}
			})

			It("sets the right offset for the new file", func() {
				f := &parsleyfakes.FakeFile{}
				fs.AddFile(f)
				Expect(f.SetOffsetCallCount()).To(Equal(1))
				offset := f.SetOffsetArgsForCall(0)
				Expect(offset).To(Equal(11))
			})
		})

		Context("when called with nil", func() {
			It("panics", func() {
				Expect(func() { fs.AddFile(nil) }).To(Panic())
			})
		})
	})

	Describe("Position()", func() {
		Context("when called with 0 offset", func() {
			It("returns with a nil position", func() {
				Expect(fs.Position(parsley.Pos(0))).To(Equal(parsley.NilPosition))
			})
		})

		Context("when the file set is empty", func() {
			It("returns with a nil position", func() {
				Expect(fs.Position(parsley.Pos(1))).To(Equal(parsley.NilPosition))
			})
		})

		Context("it has files", func() {
			var (
				f1, f2 *parsleyfakes.FakeFile
				p1, p2 *parsleyfakes.FakePosition
			)

			BeforeEach(func() {
				p1 = &parsleyfakes.FakePosition{}
				p2 = &parsleyfakes.FakePosition{}
				f1 = &parsleyfakes.FakeFile{}
				f2 = &parsleyfakes.FakeFile{}
				f1.LenReturns(10)
				f1.PositionReturns(p1)
				f2.LenReturns(20)
				f2.PositionReturns(p2)
				files = []parsley.File{f1, f2}
			})

			It("returns with a position from the first file (beginning)", func() {
				Expect(fs.Position(parsley.Pos(1))).To(BeIdenticalTo(p1))
				Expect(f1.PositionArgsForCall(0)).To(Equal(0))

				Expect(fs.Position(parsley.Pos(5))).To(BeIdenticalTo(p1))
				Expect(f1.PositionArgsForCall(1)).To(Equal(4))

				Expect(fs.Position(parsley.Pos(10))).To(BeIdenticalTo(p1))
				Expect(f1.PositionArgsForCall(2)).To(Equal(9))
			})

			It("returns with a position from the second file", func() {
				Expect(fs.Position(parsley.Pos(11))).To(BeIdenticalTo(p2))
				Expect(f2.PositionArgsForCall(0)).To(Equal(0))

				Expect(fs.Position(parsley.Pos(20))).To(BeIdenticalTo(p2))
				Expect(f2.PositionArgsForCall(1)).To(Equal(9))

				Expect(fs.Position(parsley.Pos(30))).To(BeIdenticalTo(p2))
				Expect(f2.PositionArgsForCall(2)).To(Equal(19))
			})

			It("returns with a nil position if position is outside of the file set", func() {
				Expect(fs.Position(parsley.Pos(31))).To(Equal(parsley.NilPosition))
			})
		})
	})

})