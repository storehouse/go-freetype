package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

type GlyphSlot struct {
	handle C.FT_GlyphSlot
}

func (g *GlyphSlot) Advance() *Vector {
	return &Vector{g.handle.advance}
}

func (g *GlyphSlot) Bitmap() *Bitmap {
	return &Bitmap{handle: g.handle.bitmap}
}

func (g *GlyphSlot) BitmapLeft() int {
	return int(g.handle.bitmap_left)
}

func (g *GlyphSlot) BitmapTop() int {
	return int(g.handle.bitmap_top)
}
