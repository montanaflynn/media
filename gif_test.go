package media

import (
	"os"
	"testing"
)

func TestGIF(t *testing.T) {
	tests := []struct {
		name      string
		filePath  string
		height    int
		width     int
		mediaType ContentType
		err       error
	}{
		{"valid gif", "./test-images/test.gif", 186, 240, GIFType, nil},
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
