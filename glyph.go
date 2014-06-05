package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
#include FT_GLYPH_H
*/
import "C"

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
