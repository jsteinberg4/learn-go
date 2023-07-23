// Exercise: rot13Reader
//
// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
// For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns
// a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).
//
// TODO:
// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by
// applying the rot13 substitution cipher to all alphabetical characters. The rot13Reader type is provided
// for you. Make it an io.Reader by implementing its Read method.
// rot13 cipher: https://en.wikipedia.org/wiki/ROT13
package exercises

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func isUpperAscii(b byte) bool {
	return (b >= 65 && b <= 90)
}

func isLowerAscii(b byte) bool {
	return b >= 97 && b <= 122
}

func rotate(b byte, upper byte) byte {
	b = upper - (b - 13)
	b %= 26
	return upper - b
}

func (rot13 rot13Reader) Read(dst []byte) (int, error) {
	n, err := rot13.r.Read(dst)

	// Apply cipher if the underlying reader didn't error
	if err == nil {
		for i := 0; i < n; i++ {
			if isUpperAscii(dst[i]) {
				dst[i] = rotate(dst[i], 'Z')
			} else if isLowerAscii(dst[i]) {
				dst[i] = rotate(dst[i], 'z')
			}
		}

	}

	return n, err
}

func Rot13ReaderTest() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
