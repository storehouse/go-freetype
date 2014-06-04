package freetype

import (
	. "gopkg.in/check.v1"
	"image"
	"io/ioutil"
	//"image/png"
	//"os"
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

	// free memory
	err = face.Done()
	c.Assert(err, IsNil)

	err = memoryFace.Done()
	c.Assert(err, IsNil)
}

func (s *FreetypeSuite) TestTutorial1Refined(c *C) {
	face, err := NewFace(s.lib, s.fileName, 0)
	c.Assert(err, IsNil)
	c.Assert(face, Not(IsNil))

	err = face.SetCharSize(0, 16*64, 72, 72)
	c.Assert(err, IsNil)

	slot := face.Glyph()
	var (
		penX int = 16
		penY int = 16
		n    int
		text string = "Hello, World" // "Hello, 世界" doesn't render correctly
		// because the last the chars are under East Asian Scripts, CJK Unified Ideographs
		// the charmap picked by default is the first one which is European Alphabets, Basic Latin

		img *image.Gray = image.NewGray(image.Rect(0, 0, 256, 256))
	)

	for n = 0; n < len(text); n++ {
		err = face.LoadChar(uint64(text[n]), LoadRender)
		c.Check(err, IsNil)

		drawBitmap(img, slot.Bitmap(), penX, penY, slot.BitmapLeft(), slot.BitmapTop())

		penX += int(slot.Advance().X()) >> 6
	}
	/*
		file, err := os.Create("test.png")
		c.Assert(err, IsNil)
		err = png.Encode(file, img)
		c.Assert(err, IsNil)
	*/

	err = face.Done()
	c.Assert(err, IsNil)
}

func drawBitmap(img *image.Gray, bitmap *Bitmap, x, y, left, top int) error {
	b, err := bitmap.GrayImage()
	if err != nil {
		return err
	}

	rec := b.Bounds()
	for i := 0; i < rec.Dx(); i++ {
		for j := 0; j < rec.Dy(); j++ {
			img.Set(x+i, y+j-top, b.At(i, j))
		}
	}
	return nil
}
