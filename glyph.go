package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
#include FT_GLYPH_H
*/
import "C"

const (
	GlyphBBoxUnscaled  int = C.FT_GLYPH_BBOX_UNSCALED
	GlyphBBoxSubpixels int = C.FT_GLYPH_BBOX_SUBPIXELS
	GlyphBBoxGridfit   int = C.FT_GLYPH_BBOX_GRIDFIT
	GlyphBBoxTruncate  int = C.FT_GLYPH_BBOX_TRUNCATE
	GlyphBBoxPixels    int = C.FT_GLYPH_BBOX_PIXELS
)

type Glyph struct {
	handle C.FT_Glyph
}

func (g *Glyph) Library() *Library {
	return &Library{g.handle.library}
}

func (g *Glyph) Format() int {
	return int(g.handle.format)
}

func (g *Glyph) Advance() *Vector {
	return &Vector{g.handle.advance}
}

// functions

func (g *Glyph) Done() {
	C.FT_Done_Glyph(g.handle)
}

func (g *Glyph) Copy() (*Glyph, error) {
	var glyph2 C.FT_Glyph
	errno := C.FT_Glyph_Copy(g.handle, &glyph2)
	if errno != 0 {
		return nil, GetError(errno)
	}
	return &Glyph{glyph2}, nil
}

func (g *Glyph) GetCBox(bboxMode uint) *BBox {
	var bbox C.FT_BBox
	C.FT_Glyph_Get_CBox(g.handle, C.FT_UInt(bboxMode), &bbox)
	return &BBox{bbox}
}

func (g *Glyph) ToBitmap(renderMode int, origin *Vector, destroy bool) error {
	var d C.FT_Bool
	if destroy {
		d = 1
	}
	errno := C.FT_Glyph_To_Bitmap(&g.handle, C.FT_Render_Mode(renderMode), &origin.handle, d)
	if errno != 0 {
		return GetError(errno)
	}
	return nil
}

func (g *Glyph) Transform(matrix *Matrix, delta *Vector) error {
	errno := C.FT_Glyph_Transform(g.handle, &matrix.handle, &delta.handle)
	if errno != 0 {
		return GetError(errno)
	}
	return nil
}
