package media

import (
	"encoding/binary"
)

// BMPFormat is the type of the media
type BMPFormat string

const (
	// OS2Format bmp format
	OS2Format BMPFormat = "OS/2"

	// WindowsFormat bmp format
	WindowsFormat = "Windows"
)

// BMP holds the BMP
type BMP struct {
	ContentType ContentType
	Dimensions  Dimensions
	Format      BMPFormat
}

// Size returns the BMP size in width and height
func (b BMP) Size() Dimensions {
	return b.Dimensions
}

// Type returns the BMP content type
func (b BMP) Type() ContentType {
	return b.ContentType
}

func bmpDimensions(buf []byte) (Media, error) {
	headerSize := int(binary.LittleEndian.Uint16(buf[14:18]))
	offset := 18
	if headerSize == 12 {
		width := int(binary.LittleEndian.Uint16(buf[offset : offset+2]))
		height := int(binary.LittleEndian.Uint16(buf[(offset + 2) : offset+4]))
		dimensions := Dimensions{width, height}
		return BMP{BMPType, dimensions, OS2Format}, nil
	} else if headerSize == 40 {
		offset := 18
		width := int(binary.LittleEndian.Uint16(buf[offset : offset+4]))
		height := int(binary.LittleEndian.Uint16(buf[(offset + 4) : offset+8]))
		dimensions := Dimensions{width, height}
		return BMP{BMPType, dimensions, WindowsFormat}, nil
	}

	return nil, ErrBMPInvalidHeaderLength
}
