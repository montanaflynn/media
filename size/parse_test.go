package mediasize

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		filePath  string
		height    int
		width     int
		imageType ImageType
		err       error
	}{
		{"valid png", "./test-images/test.png", 876, 1446, PNG, nil},
		{"valid gif", "./test-images/test.gif", 186, 240, GIF, nil},
		{"valid jpeg", "./test-images/test.jpg", 186, 240, GIF, ErrUnknownImageType},
		{"invalid gif", "./test-images/test-fail-1-byte.gif", 186, 240, GIF, io.EOF},
		{"invalid gif", "./test-images/test-fail-5-bytes.gif", 186, 240, GIF, ErrMissingGIFHeaders},
	}
	for _, tt := range tests {
		tt := tt
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
			if info.ImageType != tt.imageType {
				t.Errorf("ImageType got %q, want %q", info.ImageType, tt.imageType)
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

func TestParseReaderErrors(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		err    error
	}{
		{"first read error", bytes.NewReader([]byte("")), io.EOF},
		{"second read error", bytes.NewReader([]byte("GIF")), io.EOF},
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
