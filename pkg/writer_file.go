package pkg

import (
	"io"
	"os"
)

// FileWriter returns a writer for the given output file.
func FileWriter(outputfile string) (w io.WriteCloser, err error) {
	switch outputfile {
	case "stdout", "":
		w = os.Stdout
	default:
		w, err = os.OpenFile(outputfile, os.O_CREATE|os.O_WRONLY, 0644)
	}
	return
}
