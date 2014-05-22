package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

type SizeMetrices struct {
	handle C.FT_Size_Metrics
}

// horizontal pixels per EM
func (size *SizeMetrices) XPpem() int64 {
	return int64(size.handle.x_ppem)
}

// vertical pixels per EM
func (size *SizeMetrices) YPpem() int64 {
	return int64(size.handle.y_ppem)
}

// scaling values used to convert font
func (size *SizeMetrices) XScale() int64 {
	return int64(size.handle.x_scale)
}

// units to 26.6 fractional pixels
func (size *SizeMetrices) YScale() int64 {
	return int64(size.handle.y_scale)
}

// ascender in 26.6 frac. pixels
func (size *SizeMetrices) Ascender() int64 {
	return int64(size.handle.ascender)
}

// descender in 26.6 frac. pixels
func (size *SizeMetrices) Descender() int64 {
	return int64(size.handle.descender)
}

// text height in 26.6 frac. pixels
func (size *SizeMetrices) Height() int64 {
	return int64(size.handle.height)
}

// max horizontal advance, in 26.6 pixels
func (size *SizeMetrices) MaxAdvance() int64 {
	return int64(size.handle.max_advance)
}
