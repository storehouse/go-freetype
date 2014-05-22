package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

type Bitmap struct {
	handle C.FT_Bitmap
}

func (b *Bitmap) Rows() int {
	return int(b.handle.rows)
}

func (b *Bitmap) Width() int {
	return int(b.handle.width)
}

func (b *Bitmap) Pitch() int {
	return int(b.handle.pitch)
}

func (b *Bitmap) Buffer() []byte {
	// TODO
	return nil
}

func (b *Bitmap) NumGrays() int {
	return int(b.handle.num_grays)
}

func (b *Bitmap) PixelMode() int {
	return int(b.handle.pixel_mode)
}
