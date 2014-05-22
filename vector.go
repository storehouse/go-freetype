package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

type Vector struct {
	handle C.FT_Vector
}

func (v *Vector) X() int64 {
	return int64(v.handle.x)
}

func (v *Vector) Y() int64 {
	return int64(v.handle.y)
}
