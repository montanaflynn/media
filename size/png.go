package size

import (
	"encoding/binary"
)

func pngDimensions(buf []byte) (Size, error) {
	offset := 16
	width := int(binary.BigEndian.Uint32(buf[offset : offset+4]))
	height := int(binary.BigEndian.Uint32(buf[(offset + 4) : offset+12]))
	return Size{width, height, PNG}, nil
}
