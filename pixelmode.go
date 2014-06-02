package freetype

/*
#cgo pkg-config: freetype2
#include <ft2build.h>
#include FT_FREETYPE_H
*/
import "C"

const (
	PixelModeNone  = C.FT_PIXEL_MODE_NONE
	PixelModeMono  = C.FT_PIXEL_MODE_MONO
	PixelModeGray  = C.FT_PIXEL_MODE_GRAY
	PixelModeGray2 = C.FT_PIXEL_MODE_GRAY2
	PixelModeGray4 = C.FT_PIXEL_MODE_GRAY4
	PixelModeLCD   = C.FT_PIXEL_MODE_LCD
	PixelModeLCDV  = C.FT_PIXEL_MODE_LCD_V
	PixelModeBGRA  = C.FT_PIXEL_MODE_BGRA
	PixelModeMax   = C.FT_PIXEL_MODE_MAX
)
