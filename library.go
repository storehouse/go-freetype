package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

import (
	"fmt"
)

// A handle to a FreeType library instance.
// Each ‘library’ is completely independent from the others;
// it is the ‘root’ of a set of objects like fonts, faces, sizes, etc.
// In multi-threaded applications, make sure that the same library object or any of its children doesn't get accessed in parallel.
type Library struct {
	handle C.FT_Library
}

func InitFreeType() (*Library, error) {
	lib := &Library{}
	errno := C.FT_Init_FreeType(&lib.handle)
	if errno != 0 {
		return nil, fmt.Errorf("Could not init freetype library. Error code: %d", errno)
	}
	return lib, nil
}

func (l *Library) Done() error {
	errno := C.FT_Done_FreeType(l.handle)
	if errno != 0 {
		return fmt.Errorf("Library.Done: Could not free Library. Error code: %d", errno)
	}
	return nil
}
