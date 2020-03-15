package media

import (
	"bytes"
	"io"
)

const bytesToRead = 32

type detector interface {
	match(data []byte) (ContentType, error)
}

var mediaSignatures = []detector{
	&mediaSignature{[]byte("BM"), BMPType},
	&mediaSignature{[]byte("GIF87a"), GIFType},
	&mediaSignature{[]byte("GIF89a"), GIFType},
	&mediaSignature{[]byte("\x89PNG\x0D\x0A\x1A\x0A"), PNGType},
	&mediaSignature{[]byte("\xFF\xD8\xFF"), JPEGType},
}

type mediaSignature struct {
	sig       []byte
	mediaType ContentType
}

func (e *mediaSignature) match(data []byte) (ContentType, error) {
	if bytes.HasPrefix(data, e.sig) {
		return e.mediaType, nil
	}
	return "", ErrUnknownContentType
}

// DetectContentType returns the ContentType from the first 32 bytes
func DetectContentType(r io.Reader) (ContentType, []byte, error) {
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

	return "", readBytes, ErrUnknownContentType
}
