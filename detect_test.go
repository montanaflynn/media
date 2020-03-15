package media

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"
)

func make1KiloBytes() []byte {
	token := make([]byte, 1024)
	rand.Read(token)
	return token
}

func TestDetectContentType(t *testing.T) {
	var tests = []struct {
		name      string
		data      []byte
		mediaType ContentType
		err       error
	}{
		// Some nonsense.
		{"Empty", []byte{}, "", io.EOF},
		{"Binary", []byte{1, 2, 3}, "", ErrUnknownContentType},
		{"Long", make1KiloBytes(), "", ErrUnknownContentType},
		{"Whitespace Prefix", []byte("    \twhitespace"), "", ErrUnknownContentType},

		// Image types.
		{"BMP image", []byte("BM..."), BMPType, nil},
		{"GIF 87a", []byte(`GIF87a`), GIFType, nil},
		{"GIF 89a", []byte(`GIF89a...`), GIFType, nil},
		{"PNG image", []byte("\x89PNG\x0D\x0A\x1A\x0A"), PNGType, nil},
		{"JPEG image", []byte("\xFF\xD8\xFF"), JPEGType, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt, _, err := DetectContentType(bytes.NewReader(tt.data))
			if err != tt.err {
				t.Errorf("%v: DetectContentType = %q, want %q", tt.name, err, tt.err)
			}
			if mt != tt.mediaType {
				t.Errorf("%v: DetectContentType = %q, want %q", tt.name, mt, tt.mediaType)
			}
		})
	}
}
