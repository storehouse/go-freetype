package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

import (
	"image"
	"reflect"
	"unsafe"
)

type Bitmap struct {
	handle C.FT_Bitmap
}

// The number of bitmap rows.
func (b *Bitmap) Rows() int {
	return int(b.handle.rows)
}

// The number of pixels in bitmap row.
func (b *Bitmap) Width() int {
	return int(b.handle.width)
}

// The pitch's absolute value is the number of bytes taken by one bitmap row,
// including padding. However, the pitch is positive when the bitmap has a ‘down’ flow,
// and negative when it has an ‘up’ flow. In all cases,
// the pitch is an offset to add to a bitmap pointer in order to go down one row.
func (b *Bitmap) Pitch() int {
	return int(b.handle.pitch)
}

// Buffer returns the bitmap buffer.
func (b *Bitmap) Buffer() []byte {
	l := b.handle.rows * b.handle.pitch
	return C.GoBytes(unsafe.Pointer(b.handle.buffer), l)
}

// Image returns an Go image.Image.
func (b *Bitmap) Image() (image.Image, error) {
	// TODO Support the other pixel modes
	switch b.PixelMode() {
	case PixelModeNone, PixelModeMono, PixelModeGray2, PixelModeGray4, PixelModeLCD, PixelModeLCDV, PixelModeBGRA:
		return nil, ErrUnsupportedPixelMode
	case PixelModeGray:
		if b.handle.num_grays == 256 {
			size := int(b.handle.rows * b.handle.width)
			header := reflect.SliceHeader{
				Data: uintptr(unsafe.Pointer(b.handle.buffer)),
				Len:  size,
				Cap:  size,
			}
			return &image.Gray{
				Pix:    *(*[]byte)(unsafe.Pointer(&header)),
				Stride: int(b.handle.width),
				Rect:   image.Rect(0, 0, int(b.handle.width), int(b.handle.rows)),
			}, nil
		}
	}
	return nil, ErrUnsupportedPixelMode
}

// This field is only used with PixelModeGray;
// it gives the number of gray levels used in the bitmap.
func (b *Bitmap) NumGrays() int {
	return int(b.handle.num_grays)
}

// The pixel mode, i.e., how pixel bits are stored.
// See FT_Pixel_Mode for possible values.
func (b *Bitmap) PixelMode() int {
	return int(b.handle.pixel_mode)
}
