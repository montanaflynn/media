package size

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestBMP(t *testing.T) {
	tests := []struct {
		name      string
		filePath  string
		height    int
		width     int
		mediaType MediaType
		err       error
	}{
		{"valid bmp", "./test-images/test.bmp", 512, 512, BMP, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			png, err := os.Open(tt.filePath)
			if err != nil {
				t.Error(err)
				return
			}
			info, err := Parse(png)
			if err != nil {
				if err != tt.err {
					t.Errorf("Error got %q, want %q", err, tt.err)
				}
				return
			}
			if info.MediaType != tt.mediaType {
				t.Errorf("MediaType got %q, want %q", info.MediaType, tt.mediaType)
			}
			if info.Width != tt.width {
				t.Errorf("Width got %d, want %d", info.Width, tt.width)
			}
			if info.Height != tt.height {
				t.Errorf("Height got %d, want %d", info.Height, tt.height)
			}
		})
	}
}

func TestBMPErrors(t *testing.T) {
	invalidBMP := []byte{66, 77, 56, 4, 4, 0, 0, 0, 0, 0, 54, 4, 0, 0, 8}
	tests := []struct {
		name   string
		reader io.Reader
		err    error
	}{
		{"invalid header length error", bytes.NewReader(invalidBMP), ErrBMPInvalidHeaderLength},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.reader)
			if err != tt.err {
				t.Errorf("Error got %q, want %q", err, tt.err)
			}
		})
	}
}
