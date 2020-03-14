package size

import (
	"fmt"
	"io"
)

var (
	// ErrUnknownMediaType is when media is an unknown type
	ErrUnknownMediaType = fmt.Errorf("Unknown media type")

	// ErrUnsupportedSize is when media type doesn't implement size
	ErrUnsupportedSize = fmt.Errorf("Unsupported size")

	// ErrPNGMissingIHDR is when a png is missing the HDR header
	ErrPNGMissingIHDR = fmt.Errorf("Invalid png missing IHDR")
)

// MediaType is the type of the media
type MediaType string

const (
	// PNG media type
	PNG MediaType = "PNG"

	// GIF media type
	GIF = "GIF"

	// BMP media type
	BMP = "BMP"

	// JPEG media type
	JPEG = "JPEG"
)

// Size holds the media dimensions
type Size struct {
	Width     int
	Height    int
	MediaType MediaType
}

// Parse returns the media information including file type and dimensions
func Parse(r io.Reader) (Size, error) {
	mediaType, readBytes, err := DetectMediaType(r)
	if err != nil {
		return Size{}, err
	}

	switch mediaType {
	case GIF:
		return gifDimensions(readBytes)
	case PNG:
		return pngDimensions(readBytes)
	}

	return Size{MediaType: mediaType}, ErrUnsupportedSize
}
