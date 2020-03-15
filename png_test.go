package media

import (
	"os"
	"testing"
)

func TestPNG(t *testing.T) {
	tests := []struct {
		name      string
		filePath  string
		height    int
		width     int
		mediaType ContentType
		err       error
	}{
		{"valid png", "./test-images/test.png", 876, 1446, PNGType, nil},
		{"invalid gif", "./test-images/test-fail-10-bytes.png", 372, 480, PNGType, ErrPNGMissingIHDR},
	}
	for _, tt := range tests {
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
