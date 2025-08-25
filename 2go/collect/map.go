// Package collect provides exercises for working with Go's collection types, including maps and slices.
package collect

// Task:
//
// Run the TestKeys and TestValues tests:
//   go test -run TestKeys
//   go test -run TestValues
//
// You can also run all tests in the package, like this:
//   go test
//
// The above mentioned tests will sometimes pass and sometimes fail.
// Your task is to fix the functions so that they pass the tests.
// You are expected to use the [maps] and [slices] packages from the
// Go standard library to solve the exercises. Requires Go 1.23 or later.
//
// However, to avoid having to run the tests manually many times,
// you can instead use the -count flag like this:
//   go test -run TestKeys -count 100
//   go test -run TestValues -count 100
//
// Once you have fixed the functions, run the commands again to confirm
// that the tests pass when run multiple times.
import (
	"maps"
	"slices"
	"sort"
)

// keys returns the keys of a map x as a slice of strings.
func keys(x map[string]int) []string {
	keys := slices.Collect(maps.Keys(x)) // iter -> []string
	slices.Sort(keys)                     // ordena in place
	return keys
}


// values returns the values of a map x as a slice of ints.
func values(x map[string]int) []int {
    res := make([]int, 0, len(x))
    for _, v := range x {
        res = append(res, v)
    }
    sort.Ints(res) // stdlib
    return res
}

