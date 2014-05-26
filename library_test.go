package freetype

import (
	. "gopkg.in/check.v1"
	"testing"
)

const (
	fontFilePathName = "/usr/share/fonts/truetype/DejaVuSans.ttf"
)

func Test(t *testing.T) { TestingT(t) }

type FreetypeSuite struct {
	lib  *Library
	face *Face
}

var _ = Suite(&FreetypeSuite{})

func (s *FreetypeSuite) SetUpSuite(c *C) {
	var err error
	s.lib, err = InitFreeType()
	c.Assert(err, IsNil)
	c.Assert(s.lib, Not(IsNil))

	s.face, err = NewFace(s.lib, fontFilePathName, 0)
	c.Assert(err, IsNil)
	c.Assert(s.face, Not(IsNil))
}

func (s *FreetypeSuite) TearDownSuite(c *C) {
	var err error
	err = s.face.Done()
	c.Assert(err, IsNil)

	err = s.lib.Done()
	c.Assert(err, IsNil)
}
