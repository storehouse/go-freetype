package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

const (
	GLYPH_FORMAT_NONE      int = C.FT_GLYPH_FORMAT_NONE
	GLYPH_FORMAT_COMPOSITE int = C.FT_GLYPH_FORMAT_COMPOSITE
	GLYPH_FORMAT_OUTLINE   int = C.FT_GLYPH_FORMAT_OUTLINE
	GLYPH_FORMAT_PLOTTER   int = C.FT_GLYPH_FORMAT_PLOTTER
)
