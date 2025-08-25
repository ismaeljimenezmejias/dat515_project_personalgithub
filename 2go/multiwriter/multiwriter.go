package multiwriter

import (
	"dat515/2go/errors"
	"io"
)

// WriteTo writes b to all provided writers and returns the bytes written
// and errors corresponding to each writer.
func WriteTo(b []byte, writers ...io.Writer) (n []int, errs errors.Errors) {
	if len(writers) == 0 {
		return []int{}, nil
	}

	n = make([]int, len(writers))
	errs = make(errors.Errors, len(writers))

	for i, w := range writers {
		written, err := w.Write(b)
		n[i] = written

		// Si hubo error de escritura o no se escribió todo
		if err != nil {
			errs[i] = err
		} else if written != len(b) {
			errs[i] = io.ErrShortWrite
		}
	}

	// Comprobamos si todos los errores son nil
	allNil := true
	for _, e := range errs {
		if e != nil {
			allNil = false
			break
		}
	}

	if allNil {
		return n, nil
	}

	return n, errs
}
