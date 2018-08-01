package src_test

import (
	"strings"
	"testing"
)

// Write a test for strings.HasPrefix
// https://golang.org/pkg/strings/#HasPrefix
// Given the value "main.go", test that it has the prefix "main"
// Remember that your test has to start with the name `Test` and be in an `_test.go` file

// Write a table drive test for strings.Index
// https://golang.org/pkg/strings/#Index
// Use the following test conditions
// |------------------------------------------------|
// | Value                     | Substring | Answer |
// |===========================|===========|========|
// | "Gophers are amazing!"    | "are"     | 8      |
// | "Testing in Go is fun."   | "fun"     | 17     |
// | "The answer is 42."       | "is"      | 11     |
// |------------------------------------------------|

// Rewrite the above test for strings.Index using subtests

// Here is a simple test that tests `strings.HasSuffix`.i
// https://golang.org/pkg/strings/#HasSuffix
func Test_HasSuffix(t *testing.T) {
	value := "main.go"
	if !strings.HasSuffix(value, ".go") {
		t.Fatalf("expected %s to have suffix %s", value, ".go")
	}
}

func Test_Index(t *testing.T) {
	tt := []struct {
		Value     string
		Substring string
		Answer    int
	}{
		{Value: "Gophers are amazing!", Substring: "are", Answer: 8},
		{Value: "Testing in Go is fun.", Substring: "fun", Answer: 17},
		{Value: "The answer is 42.", Substring: "is", Answer: 11},
	}

	for _, x := range tt {
		if got := strings.Index(x.Value, x.Substring); got != x.Answer{
			t.Errorf("expected %d, got %d", x.Answer, got)
		}
	}
}
