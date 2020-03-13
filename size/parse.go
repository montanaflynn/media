package mediasize

import (
	"bytes"
	"fmt"
	"io"
)

var (
	// ErrMissingGIFHeaders is when a gif is missing the headers
	ErrMissingGIFHeaders = fmt.Errorf("Invalid gif missing headers")

	// ErrUnknownImageType is when an image is an unknown type
	ErrUnknownImageType = fmt.Errorf("Unknown image type")

	// ErrPNGMissingIHDR is when a png is missing the HDR header
	ErrPNGMissingIHDR = fmt.Errorf("Invalid png missing IHDR")
)

// ImageType is the type of the image
type ImageType string

const (
	// PNG image type
	PNG ImageType = "PNG"

	// GIF image type
	GIF = "GIF"
)

// Size holds the image dimensions
type Size struct {
	Width     int
	Height    int
	ImageType ImageType
}

// Parse returns the image information including file type and dimensions
func Parse(r io.Reader) (Size, error) {

	initialBytes := make([]byte, 3)
	_, err := r.Read(initialBytes)
	if err != nil {
		return Size{}, err
	}

	if string(initialBytes[:3]) == "GIF" {
		moreBytes := make([]byte, 7)
		nBytes, err := r.Read(moreBytes)
		if err != nil {
			return Size{}, err
		}

		if nBytes != 7 {
			return Size{}, ErrMissingGIFHeaders
		}
		gifHeaderBytes := append(initialBytes, moreBytes...)
		return gifDimensions(gifHeaderBytes)
	}

	moreBytes := make([]byte, 25)
	_, err = r.Read(moreBytes)
	if err != nil {
		return Size{}, err
	}
	pngHeaderBytes := append(initialBytes, moreBytes...)
	pngSignature := []byte{137, 80, 78, 71, 13, 10, 26, 10}
	res := bytes.Compare(pngHeaderBytes[:8], pngSignature)
	if res != 0 {
		return Size{}, ErrUnknownImageType
	}

	IHDR := "IHDR"
	if string(pngHeaderBytes[12:16]) != IHDR {
		return Size{}, ErrPNGMissingIHDR
	}

	return pngDimensions(pngHeaderBytes)
}
