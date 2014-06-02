package freetype

import (
	. "gopkg.in/check.v1"
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
}
