package cxcompressgzip

import (
	"io"

	"github.com/klauspost/compress/gzip"
	"github.com/valyala/bytebufferpool"
)

// Compress compresses the input byte slice using gzip compression.
func Compress(input []byte) ([]byte, error) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	gz := gzip.NewWriter(buf)
	if _, err := gz.Write(input); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decompress decompresses the input byte slice using gzip decompression.
func Decompress(input []byte) ([]byte, error) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	gz, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	if _, err := io.Copy(buf, gz); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
