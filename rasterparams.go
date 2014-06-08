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

const (
	RasterFlagDefault int = C.FT_RASTER_FLAG_DEFAULT
	RasterFlagAA      int = C.FT_RASTER_FLAG_AA
	RasterFlagDirect  int = C.FT_RASTER_FLAG_DIRECT
	RasterFlagClip    int = C.FT_RASTER_FLAG_CLIP
)

type RasterParams struct {
	handle C.FT_Raster_Params
}

func NewRasterParams() *RasterParams {
	return &RasterParams{C.FT_Raster_Params{}}
}

// The target bitmap.
func (params *RasterParams) Target() *Bitmap {
	return &Bitmap{*params.handle.target}
}

func (params *RasterParams) SetTarget(target *Bitmap) {
	params.handle.target = &target.handle
}

// TODO const void* source

// The rendering flags.
func (params *RasterParams) Flags() int {
	return int(params.handle.flags)
}

func (params *RasterParams) SetFlags(flags int) {
	params.handle.flags = C.int(flags)
}

// TODO FT_SpanFunc gray_spans

func (params *RasterParams) User() unsafe.Pointer {
	return params.handle.user
}

func (params *RasterParams) SetUser(user *interface{}) {
	params.handle.user = unsafe.Pointer(user)
}

// An optional clipping box. It is only used in direct rendering mode.
// Note that coordinates here should be expressed in integer pixels
// (and not in 26.6 fixed-point units)
func (params *RasterParams) ClipBox() *BBox {
	return &BBox{params.handle.clip_box}
}

func (params *RasterParams) SetClipBox(bbox *BBox) {
	params.handle.clip_box = bbox.handle
}
