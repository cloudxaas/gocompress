package cxcompressbrotli

import (
	"io"
	"bytes"
	
	"github.com/andybalholm/brotli"
	"github.com/valyala/bytebufferpool"
)

// Compress compresses the input byte slice using Brotli compression.
func Compress(input []byte) ([]byte, error) {
	buf := bytebufferpool.Get()
	

	w := brotli.NewWriter(buf)
	if _, err := w.Write(input); err != nil {
		bytebufferpool.Put(buf)
		return nil, err
	}
	if err := w.Close(); err != nil {
		bytebufferpool.Put(buf)
		return nil, err
	}

	// Copy the data to a new slice before returning
	output := make([]byte, buf.Len())
	copy(output, buf.B)
	bytebufferpool.Put(buf)
	return output, nil
}

// Decompress decompresses the input byte slice using Brotli decompression.
func Decompress(input []byte) ([]byte, error) {
	buf := bytebufferpool.Get()
	

	r := brotli.NewReader(bytes.NewReader(input))
	if _, err := io.Copy(buf, r); err != nil {
		bytebufferpool.Put(buf)
		return nil, err
	}

	// Copy the data to a new slice before returning
	output := make([]byte, buf.Len())
	copy(output, buf.B)
	bytebufferpool.Put(buf)
	return output, nil
}
