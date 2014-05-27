package freetype

import (
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

var (
	fonts = []string{
	  "/usr/share/fonts/truetype/ttf-droid/DroidSans.ttf",
	  "/usr/share/fonts/truetype/DejaVuSans.ttf",
	}
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

	var fontFilePathName string

	for i := range fonts {
		_, err := os.Open(fonts[i])
		if err == nil {
			fontFilePathName = fonts[i]
			break
		}
	}

	if len(fontFilePathName) == 0 {
		c.Skip("no font file found")
		return
	}

	s.face, err = NewFace(s.lib, fontFilePathName, 0)
	c.Assert(err, IsNil)
	c.Assert(s.face, Not(IsNil))
}

func (s *FreetypeSuite) TearDownSuite(c *C) {
	var err error
	if s.face != nil {
		err = s.face.Done()
		c.Assert(err, IsNil)
	}

	err = s.lib.Done()
	c.Assert(err, IsNil)
}
