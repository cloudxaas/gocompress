package cxcompressgzip

import (
	"io"
	"bytes"
	
	"github.com/klauspost/compress/gzip"
	"github.com/valyala/bytebufferpool"
)

// Compress compresses the input byte slice using gzip compression.
func Compress(input []byte) ([]byte, error) {
	buf := bytebufferpool.Get()
	

	gz := gzip.NewWriter(buf)
	if _, err := gz.Write(input); err != nil {
		bytebufferpool.Put(buf)
		return nil, err
	}
	if err := gz.Close(); err != nil {
		bytebufferpool.Put(buf)
		return nil, err
	}

	// Copy the data to a new slice before returning
	output := make([]byte, buf.Len())
	copy(output, buf.B)
	bytebufferpool.Put(buf)
	return output, nil
}

// Decompress decompresses the input byte slice using gzip decompression.
func Decompress(input []byte) ([]byte, error) {
	buf := bytebufferpool.Get()


	gz, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		bytebufferpool.Put(buf)
		return nil, err
	}
	defer gz.Close()

	if _, err := io.Copy(buf, gz); err != nil {
		bytebufferpool.Put(buf)
		return nil, err
	}

	// Copy the data to a new slice before returning
	output := make([]byte, buf.Len())
	copy(output, buf.B)
	bytebufferpool.Put(buf)
	return output, nil
}
