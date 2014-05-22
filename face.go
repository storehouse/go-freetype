package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Face struct {
	handle C.FT_Face
}

// NewFace opens a font by its filePathName.
func NewFace(library *Library, filePathName string, faceIndex int64) (*Face, error) {
	f := &Face{}
	cfilePathName := C.CString(filePathName)
	defer C.free(unsafe.Pointer(cfilePathName))
	errno := C.FT_New_Face(library.handle, cfilePathName, C.FT_Long(faceIndex), &f.handle)
	if errno != 0 {
		return nil, fmt.Errorf("Library.NewFace error code: %d", errno)
	}
	return f, nil
}

// TODO OpenFace

func (f *Face) Done() {
	C.FT_Done_Face(f.handle)
}

func (f *Face) NumFaces() int64 {
	return int64(f.handle.num_faces)
}

func (f *Face) FaceIndex() int64 {
	return int64(f.handle.face_index)
}

func (f *Face) FaceFlags() int64 {
	return int64(f.handle.face_flags)
}

func (f *Face) StyleFlags() int64 {
	return int64(f.handle.style_flags)
}

func (f *Face) FamilyName() string {
	return C.GoString((*C.char)(unsafe.Pointer(f.handle.family_name)))
}

func (f *Face) StyleName() string {
	return C.GoString((*C.char)(unsafe.Pointer(f.handle.style_name)))
}

func (f *Face) NumFixedSizes() int {
	return int(f.handle.num_fixed_sizes)
}

func (f *Face) AvailableSizes() []BitmapSize {
	l := f.NumFixedSizes()
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(f.handle.available_sizes)),
		Len:  l,
		Cap:  l,
	}
	sizes := make([]BitmapSize, l)
	goSlice := *(*[]C.FT_Bitmap_Size)(unsafe.Pointer(&hdr))
	for i := range goSlice {
		sizes[i] = BitmapSize{goSlice[i]}
	}
	return sizes
}

func (f *Face) NumCharmaps() int {
	return int(f.handle.num_charmaps)
}

func (f *Face) CharMaps() []CharMap {
	l := f.NumCharmaps()
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(f.handle.charmaps)),
		Len:  l,
		Cap:  l,
	}
	charmaps := make([]CharMap, l)
	goSlice := *(*[]C.FT_CharMap)(unsafe.Pointer(&hdr))
	for i := range goSlice {
		charmaps[i] = CharMap{goSlice[i]}
	}
	return charmaps
}

// The following member variables (down to `underline_thickness')
// are only relevant to scalable outlines; cf. @FT_Bitmap_Size
// for bitmap fonts.

func (f *Face) BBox() BBox {
	return BBox{f.handle.bbox}
}

func (f *Face) Ascender() int16 {
	return int16(f.handle.ascender)
}

func (f *Face) Descender() int16 {
	return int16(f.handle.descender)
}

func (f *Face) Height() int16 {
	return int16(f.handle.height)
}

func (f *Face) MaxAdvanceHeight() int16 {
	return int16(f.handle.max_advance_height)
}

func (f *Face) MaxAdvanceWidth() int16 {
	return int16(f.handle.max_advance_width)
}

func (f *Face) UnderlinePosition() int16 {
	return int16(f.handle.underline_position)
}

func (f *Face) UnderlineThickness() int16 {
	return int16(f.handle.underline_thickness)
}

func (f *Face) Glyph() *GlyphSlot {
	return &GlyphSlot{handle: f.handle.glyph}
}

func (f *Face) Size() Size {
	return Size{f.handle.size}
}

func (f *Face) CharMap() CharMap {
	return CharMap{f.handle.charmap}
}

// Methods

func (f *Face) SetCharSize(charWidth, charHeight int64, horzResolution, vertResolution uint) error {
	errno := C.FT_Set_Char_Size(f.handle, C.FT_F26Dot6(charWidth), C.FT_F26Dot6(charHeight), C.FT_UInt(horzResolution), C.FT_UInt(vertResolution))
	if errno != 0 {
		return fmt.Errorf("Face.SetCharSize error code: %d", errno)
	}
	return nil
}

func (f *Face) SetPixelSizes(width, height uint) error {
	errno := C.FT_Set_Pixel_Sizes(f.handle, C.FT_UInt(width), C.FT_UInt(height))
	if errno != 0 {
		return fmt.Errorf("Face.SetPixelSizes error code: %d", errno)
	}
	return nil
}

func (f *Face) LoadChar(char rune, flags int32) error {
	errno := C.FT_Load_Char(f.handle, C.FT_ULong(char), C.FT_Int32(flags))
	if errno != 0 {
		return fmt.Errorf("Face.LoadChar error code: %d", errno)
	}
	return nil
}
