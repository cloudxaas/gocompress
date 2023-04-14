package cxcompressbrotli

import (
	"io"

	"github.com/andybalholm/brotli"
	"github.com/valyala/bytebufferpool"
)

// Compress compresses the input byte slice using Brotli compression.
func Compress(input []byte) ([]byte, error) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	w := brotli.NewWriter(buf)
	if _, err := w.Write(input); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decompress decompresses the input byte slice using Brotli decompression.
func Decompress(input []byte) ([]byte, error) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	r := brotli.NewReader(bytes.NewReader(input))
	if _, err := io.Copy(buf, r); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
