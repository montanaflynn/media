package size

import (
	"encoding/binary"
)

func gifDimensions(buf []byte) (Size, error) {
	offset := 6
	width := int(binary.LittleEndian.Uint16(buf[offset : offset+2]))
	height := int(binary.LittleEndian.Uint16(buf[(offset + 2) : offset+4]))
	return Size{width, height, GIF}, nil
}
