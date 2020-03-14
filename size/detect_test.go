package size

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
		mediaType MediaType
		err       error
	}{
		// Some nonsense.
		{"Empty", []byte{}, "", io.EOF},
		{"Binary", []byte{1, 2, 3}, "", ErrUnknownMediaType},
		{"Long", make1KiloBytes(), "", ErrUnknownMediaType},
		{"Whitespace Prefix", []byte("    \twhitespace"), "", ErrUnknownMediaType},

		// Image types.
		{"BMP image", []byte("BM..."), BMP, nil},
		{"GIF 87a", []byte(`GIF87a`), GIF, nil},
		{"GIF 89a", []byte(`GIF89a...`), GIF, nil},
		{"PNG image", []byte("\x89PNG\x0D\x0A\x1A\x0A"), PNG, nil},
		{"JPEG image", []byte("\xFF\xD8\xFF"), JPEG, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt, _, err := DetectMediaType(bytes.NewReader(tt.data))
			if err != tt.err {
				t.Errorf("%v: DetectContentType = %q, want %q", tt.name, err, tt.err)
			}
			if mt != tt.mediaType {
				t.Errorf("%v: DetectContentType = %q, want %q", tt.name, mt, tt.mediaType)
			}
		})
	}
}
