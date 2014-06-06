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

// functions

func (g *Glyph) Copy() (*Glyph, error) {
	var glyph2 C.FT_Glyph
	errno := C.FT_Glyph_Copy(g.handle, &glyph2)
	if errno != 0 {
		return nil, GetError(errno)
	}
	return &Glyph{glyph2}, nil
}
