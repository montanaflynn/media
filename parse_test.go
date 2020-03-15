package media

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
		mediaType ContentType
		err       error
	}{
		{"valid png", "./test-images/test.png", 876, 1446, PNGType, nil},
		{"valid gif", "./test-images/test.gif", 186, 240, GIFType, nil},
		{"valid windows bmp", "./test-images/test.bmp", 512, 512, BMPType, nil},
		{"valid os/2 bmp", "./test-images/test-os2.bmp", 512, 512, BMPType, nil},
		{"valid jpeg", "./test-images/test.jpg", 0, 0, "", ErrUnsupportedSize},
		{"invalid gif", "./test-images/test-fail-1-byte.gif", 0, 0, "", ErrUnknownContentType},
		{"invalid gif", "./test-images/test-fail-5-bytes.gif", 0, 0, "", ErrUnknownContentType},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			png, err := os.Open(tt.filePath)
			if err != nil {
				t.Error(err)
				return
			}
			m, err := Parse(png)
			if err != nil {
				if err != tt.err {
					t.Errorf("Error got %q, want %q", err, tt.err)
				}
				return
			}
			mediaType := m.Type()
			mediaSize := m.Size()
			if mediaType != tt.mediaType {
				t.Errorf("ContentType got %q, want %q", mediaType, tt.mediaType)
			}
			if mediaSize.Width != tt.width {
				t.Errorf("Width got %d, want %d", mediaSize.Width, tt.width)
			}
			if mediaSize.Height != tt.height {
				t.Errorf("Height got %d, want %d", mediaSize.Height, tt.height)
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
		{"second read error", bytes.NewReader([]byte("GIF")), ErrUnknownContentType},
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
