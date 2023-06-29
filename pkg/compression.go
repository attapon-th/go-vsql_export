package pkg

import (
	"compress/gzip"
	"io"

	"github.com/klauspost/compress/zip"
)

// CreateGzip creates a gzip writer
func CreateGzip(fs io.Writer, compressionLevel int) (io.WriteCloser, error) {
	w, err := gzip.NewWriterLevel(fs, compressionLevel)
	return w, err
}

// CreateZip creates a zip writer
func CreateZip(fs io.Writer, anotherName string) (io.Writer, io.Closer, error) {
	z := zip.NewWriter(fs)
	w, err := z.Create(anotherName)
	return w, z, err
}
