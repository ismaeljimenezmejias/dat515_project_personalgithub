package stringer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testStudent(t *testing.T, scoreDec func()) {
	for _, test := range stringerStudentTests {
		if diff := cmp.Diff(test.want, test.in.String()); diff != "" {
			t.Errorf("String(%q): (-want +got):\n%s", test.in, diff)
			scoreDec()
		}
	}
}

var stringerStudentTests = []struct {
	in   Student
	want string
}{
	{Student{
		ID:        42,
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}, "Student ID: 42. Name: Doe, John. Age: 25."},
	{Student{
		ID:        1234,
		FirstName: "Tormod",
		LastName:  "Lea",
		Age:       30,
	}, "Student ID: 1234. Name: Lea, Tormod. Age: 30."},
	{Student{
		ID:        1814,
		FirstName: "Ola",
		LastName:  "Nordmann",
		Age:       50,
	}, "Student ID: 1814. Name: Nordmann, Ola. Age: 50."},
}
