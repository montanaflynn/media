package media

import (
	"encoding/binary"
)

// PNG holds the PNG
type PNG struct {
	ContentType ContentType
	Dimensions  Dimensions
}

// Size returns the PNG size in width and height
func (b PNG) Size() Dimensions {
	return b.Dimensions
}

// Type returns the PNG content type
func (b PNG) Type() ContentType {
	return b.ContentType
}

func pngDimensions(buf []byte) (Media, error) {
	IHDR := "IHDR"
	if string(buf[12:16]) != IHDR {
		return nil, ErrPNGMissingIHDR
	}

	offset := 16
	width := int(binary.BigEndian.Uint32(buf[offset : offset+4]))
	height := int(binary.BigEndian.Uint32(buf[(offset + 4) : offset+12]))
	dimensions := Dimensions{width, height}
	return PNG{PNGType, dimensions}, nil
}
