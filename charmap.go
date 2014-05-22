package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

// A handle to a given character map.
// A charmap is used to translate character codes in a given encoding into glyph indexes for its parent's face.
// Some font formats may provide several charmaps per font.
type CharMap struct {
	handle C.FT_CharMap
}

func (c *CharMap) Face() Face {
	return Face{handle: c.handle.face}
}

func (c *CharMap) Encoding() uint64 {
	return uint64(c.handle.encoding)
}

func (c *CharMap) PlatformID() uint16 {
	return uint16(c.handle.platform_id)
}

func (c *CharMap) EncodingID() uint16 {
	return uint16(c.handle.encoding_id)
}
