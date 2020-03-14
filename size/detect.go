package size

import (
	"bytes"
	"io"
)

const bytesToRead = 32

type detector interface {
	match(data []byte) (MediaType, error)
}

var mediaSignatures = []detector{
	&mediaSignature{[]byte("BM"), BMP},
	&mediaSignature{[]byte("GIF87a"), GIF},
	&mediaSignature{[]byte("GIF89a"), GIF},
	&mediaSignature{[]byte("\x89PNG\x0D\x0A\x1A\x0A"), PNG},
	&mediaSignature{[]byte("\xFF\xD8\xFF"), JPEG},
}

type mediaSignature struct {
	sig       []byte
	mediaType MediaType
}

func (e *mediaSignature) match(data []byte) (MediaType, error) {
	if bytes.HasPrefix(data, e.sig) {
		return e.mediaType, nil
	}
	return "", ErrUnknownMediaType
}

// DetectMediaType returns the MediaType from the first 32 bytes
func DetectMediaType(r io.Reader) (MediaType, []byte, error) {
	readBytes := make([]byte, bytesToRead)
	_, err := r.Read(readBytes)
	if err != nil {
		return "", readBytes, err
	}

	for _, sig := range mediaSignatures {
		mt, err := sig.match(readBytes)
		if err == nil {
			return mt, readBytes, nil
		}
	}

	return "", readBytes, ErrUnknownMediaType
}
