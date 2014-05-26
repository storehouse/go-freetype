package freetype

import (
	. "gopkg.in/check.v1"
)

func (s *FreetypeSuite) TestFace(c *C) {
	face, err := NewFace(s.lib, "", 0)
	c.Assert(err, Equals, ErrCanNotOpenResource)
	c.Assert(face, IsNil)
}
