package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

const (
	GLYPH_FORMAT_NONE      uint64 = C.FT_GLYPH_FORMAT_NONE
	GLYPH_FORMAT_COMPOSITE uint64 = C.FT_GLYPH_FORMAT_COMPOSITE
	GLYPH_FORMAT_OUTLINE   uint64 = C.FT_GLYPH_FORMAT_OUTLINE
	GLYPH_FORMAT_PLOTTER   uint64 = C.FT_GLYPH_FORMAT_PLOTTER
)
