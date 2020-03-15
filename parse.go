package media

import (
	"errors"
	"io"
)

var (
	// ErrUnknownContentType is when media is an unknown type
	ErrUnknownContentType = errors.New("media: unknown media type")

	// ErrUnsupportedSize is when media type doesn't implement size
	ErrUnsupportedSize = errors.New("media: unsupported size")

	// ErrPNGMissingIHDR is when a png is missing the HDR header
	ErrPNGMissingIHDR = errors.New("media: invalid png missing IHDR")

	// ErrBMPInvalidHeaderLength is when a bmp has invalid header length
	ErrBMPInvalidHeaderLength = errors.New("media: invalid bmp header length")
)

// ContentType is the type of the media
type ContentType string

const (
	// PNGType media type
	PNGType ContentType = "image/png"

	// GIFType media type
	GIFType = "image/gif"

	// BMPType media type
	BMPType = "image/bmp"

	// JPEGType media type
	JPEGType = "image/jpeg"
)

// Dimensions holds the Dimensions
type Dimensions struct {
	Width  int
	Height int
}

// Media is an interface for getting information from media
type Media interface {
	Size() Dimensions
	Type() ContentType
}

// Parse returns the media information including type and dimensions
func Parse(r io.Reader) (Media, error) {
	contentType, readBytes, err := DetectContentType(r)
	if err != nil {
		return nil, err
	}

	switch contentType {
	case GIFType:
		return gifDimensions(readBytes)
	case PNGType:
		return pngDimensions(readBytes)
	case BMPType:
		return bmpDimensions(readBytes)
	}

	return nil, ErrUnsupportedSize
}
