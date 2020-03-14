package size

import (
	"encoding/binary"
)

func bmpDimensions(buf []byte) (Size, error) {
	headerSize := int(binary.LittleEndian.Uint16(buf[14:18]))
	if headerSize == 12 {
		offset := 18
		width := int(binary.LittleEndian.Uint16(buf[offset : offset+2]))
		height := int(binary.LittleEndian.Uint16(buf[(offset + 2) : offset+4]))
		return Size{width, height, BMP}, nil
	} else if headerSize == 40 {
		offset := 18
		width := int(binary.LittleEndian.Uint16(buf[offset : offset+4]))
		height := int(binary.LittleEndian.Uint16(buf[(offset + 4) : offset+8]))
		return Size{width, height, BMP}, nil
	}

	return Size{}, ErrBMPInvalidHeaderLength
}
