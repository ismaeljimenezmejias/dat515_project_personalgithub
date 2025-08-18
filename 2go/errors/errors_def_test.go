package errors

import "testing"

func init() {
	scores.Add(TestGo_Errors, len(errTests), 5)
}

func TestGo_Errors(t *testing.T) {
	sc := scores.Max()
	defer sc.Print(t)
	testErrors(t, sc.Dec)
}
