// Package cipher provides an exercise for the ROT13 character substitution cipher.
package cipher

import (
	"io"
)

/*
Task: Rot 13

This task is taken from http://tour.golang.org.

A common pattern is an io.Reader that wraps another io.Reader, modifying the
stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of
compressed data) and returns a *gzip.Reader that also implements io.Reader (a
stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
modifying the stream by applying the rot13 substitution cipher to all
alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing
its Read method.
*/

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
    // Leer del Reader interno
    n, err = r.r.Read(p) //el primer r corresponder a rot12Reader, el segundo r al reader de esa funcion
    if n > 0 {
        // Transformar los bytes en ROT13
        for i := 0; i < n; i++ {
            c := p[i]
            switch {
            case 'A' <= c && c <= 'Z':
                p[i] = 'A' + (c-'A'+13)%26
            case 'a' <= c && c <= 'z':
                p[i] = 'a' + (c-'a'+13)%26
            }
        }
    }
    return n, err
}