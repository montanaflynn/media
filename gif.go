package media

import (
	"encoding/binary"
)

// GIF holds the GIF
type GIF struct {
	ContentType ContentType
	Dimensions  Dimensions
}

// Size returns the GIF size in width and height
func (b GIF) Size() Dimensions {
	return b.Dimensions
}

// Type returns the GIF media type
func (b GIF) Type() ContentType {
	return b.ContentType
}

func gifDimensions(buf []byte) (Media, error) {
	offset := 6
	width := int(binary.LittleEndian.Uint16(buf[offset : offset+2]))
	height := int(binary.LittleEndian.Uint16(buf[(offset + 2) : offset+4]))
	dimensions := Dimensions{width, height}
	return GIF{GIFType, dimensions}, nil
}
