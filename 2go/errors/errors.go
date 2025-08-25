// Package errors provides custom error handling implementations.
package errors

/*
Task: Errors needed for multiwriter

You may find this blog post useful:
http://blog.golang.org/error-handling-and-go

Similar to a the Stringer interface, the error interface also defines a
method that returns a string.

	type error interface {
	    Error() string
	}

Thus also the error type can describe itself as a string. The fmt package (and
many others) use this Error() method to print errors.

Implement the Error() method for the Errors type defined above.

The following conditions should be covered:

1. When there are no errors in the slice, it should return:

"(0 errors)"

2. When there is one error in the slice, it should return:

The error string return by the corresponding Error() method.

3. When there are two errors in the slice, it should return:

The first error + " (and 1 other error)"

4. When there are X>1 errors in the slice, it should return:

The first error + " (and X other errors)"
*/

// Error returns a string representation of the Errors type.
// Asumo que Errors está definido así:

import "fmt"

// Error implements the error interface for Errors.
func (m Errors) Error() string {
	n := len(m)
	if n == 0 {
		return "(0 errors)"
	}

	// Primer error no nil
	var first string
	for _, e := range m {
		if e != nil {
			first = e.Error()
			break
		}
	}

	if first == "" { // todos los errores son nil
		return "(0 errors)"
	}

	// Contamos cuántos errores no nil hay aparte del primero
	count := 0
	for _, e := range m {
		if e != nil {
			count++
		}
	}

	switch count {
	case 1:
		return first
	case 2:
		return fmt.Sprintf("%s (and 1 other error)", first)
	default:
		return fmt.Sprintf("%s (and %d other errors)", first, count-1)
	}
}