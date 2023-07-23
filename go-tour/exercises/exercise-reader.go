// Exercise: Readers
//
// Implement an `io.Reader` type that emits an ASCII character 'A'
package exercises

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (reader MyReader) Read(dst []byte) (int, error) {
	n := len(dst)

	for i := range dst {
		dst[i] = 'A'
	}

	return n, nil
}

func ReaderValidate() {
	reader.Validate(MyReader{})
}
