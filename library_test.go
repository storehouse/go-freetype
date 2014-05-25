package freetype

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type FreetypeSuite struct{}

var _ = Suite(&FreetypeSuite{})

func (s *FreetypeSuite) TestLibrary(c *C) {
	lib, err := InitFreeType()
	c.Assert(err, IsNil)
	c.Assert(lib, Not(IsNil))

	err = lib.Done()
	c.Assert(err, IsNil)
}
