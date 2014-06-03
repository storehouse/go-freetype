package freetype

import (
	. "gopkg.in/check.v1"
	"io/ioutil"
)

func (s *FreetypeSuite) TestFace(c *C) {
	face, err := NewFace(s.lib, "", 0)
	c.Assert(err, Equals, ErrCanNotOpenResource)
	c.Assert(face, IsNil)
}

func (s *FreetypeSuite) TestTutorial1(c *C) {
	face, err := NewFace(s.lib, s.fileName, 0)
	c.Assert(err, IsNil)
	c.Assert(face, Not(IsNil))

	err = face.SetCharSize(0, 16*64, 300, 300)
	c.Assert(err, IsNil)

	err = face.SetPixelSizes(0, 16)
	c.Assert(err, IsNil)

	index := face.GetCharIndex('A')
	c.Check(index, Equals, uint(36))

	err = face.LoadGlyph(index, LoadDefault)
	c.Assert(err, IsNil)

	err = face.Glyph().Render(RenderModeNormal)
	c.Assert(err, IsNil)

	// From memory
	data, err := ioutil.ReadFile(s.fileName)
	c.Assert(err, IsNil)
	memoryFace, err := NewMemoryFace(s.lib, data, 0)

	err = memoryFace.SetCharSize(0, 16*64, 300, 300)
	c.Assert(err, IsNil)

	err = memoryFace.SetPixelSizes(0, 16)
	c.Assert(err, IsNil)

	memoryIndex := memoryFace.GetCharIndex('A')
	c.Check(memoryIndex, Equals, index)

	err = memoryFace.LoadGlyph(memoryIndex, LoadDefault)
	c.Assert(err, IsNil)

	err = memoryFace.Glyph().Render(RenderModeNormal)
	c.Assert(err, IsNil)
}
