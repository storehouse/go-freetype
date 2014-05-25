package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

import (
	"unsafe"
)

type GlyphSlot struct {
	handle C.FT_GlyphSlot
}

func (g *GlyphSlot) Library() *Library {
	return &Library{g.handle.library}
}

func (g *GlyphSlot) Face() *Face {
	return &Face{g.handle.face}
}

func (g *GlyphSlot) Next() *GlyphSlot {
	return &GlyphSlot{g.handle.next}
}

func (g *GlyphSlot) Metrices() *GlyphMetrics {
	return &GlyphMetrics{g.handle.metrics}
}

func (g *GlyphSlot) LinearHoriAdvance() int64 {
	return int64(g.handle.linearHoriAdvance)
}

func (g *GlyphSlot) LinearVertAdvance() int64 {
	return int64(g.handle.linearVertAdvance)
}

func (g *GlyphSlot) Advance() *Vector {
	return &Vector{g.handle.advance}
}

func (g *GlyphSlot) Format() uint64 {
	return uint64(g.handle.format)
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

func (g *GlyphSlot) NumSubGlyphs() uint {
	return uint(g.handle.num_subglyphs)
}

func (g *GlyphSlot) ControlData() []byte {
	return C.GoBytes(unsafe.Pointer(g.handle.control_data), C.int(g.handle.control_len))
}

func (g *GlyphSlot) ControlLen() int {
	return int(g.handle.control_len)
}

func (g *GlyphSlot) LsbDelta() int64 {
	return int64(g.handle.lsb_delta)
}

func (g *GlyphSlot) RsbDelta() int64 {
	return int64(g.handle.rsb_delta)
}
