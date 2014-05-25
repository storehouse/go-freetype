package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

type GlyphMetrics struct {
	handle C.FT_Glyph_Metrics
}

func (m *GlyphMetrics) Width() int64 {
	return int64(m.handle.width)
}

func (m *GlyphMetrics) Height() int64 {
	return int64(m.handle.height)
}

func (m *GlyphMetrics) HoriBearingX() int64 {
	return int64(m.handle.horiBearingX)
}

func (m *GlyphMetrics) HoriBearingY() int64 {
	return int64(m.handle.horiBearingY)
}

func (m *GlyphMetrics) HoriAdvance() int64 {
	return int64(m.handle.horiAdvance)
}

func (m *GlyphMetrics) VertBearingX() int64 {
	return int64(m.handle.vertBearingX)
}

func (m *GlyphMetrics) VertBearingY() int64 {
	return int64(m.handle.vertBearingY)
}

func (m *GlyphMetrics) VertAdvance() int64 {
	return int64(m.handle.vertAdvance)
}
