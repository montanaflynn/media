package size

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
		mediaType MediaType
		err       error
	}{
		{"valid png", "./test-images/test.png", 876, 1446, PNG, nil},
		{"valid gif", "./test-images/test.gif", 186, 240, GIF, nil},
		{"valid jpeg", "./test-images/test.jpg", 0, 0, "", ErrUnsupportedSize},
		{"invalid gif", "./test-images/test-fail-1-byte.gif", 0, 0, "", ErrUnknownMediaType},
		{"invalid gif", "./test-images/test-fail-5-bytes.gif", 0, 0, "", ErrUnknownMediaType},
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

func TestParseReaderErrors(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		err    error
	}{
		{"first read error", bytes.NewReader([]byte("")), io.EOF},
		{"second read error", bytes.NewReader([]byte("GIF")), ErrUnknownMediaType},
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
