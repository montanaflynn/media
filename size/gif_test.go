package size

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
		mediaType MediaType
		err       error
	}{
		{"valid gif", "./test-images/test.gif", 186, 240, GIF, nil},
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
