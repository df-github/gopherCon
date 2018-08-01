package profile_test

import (
	"testing"
	"strings"
)

const Message = "Education is what remains after one has forgotten what one has learned in school."

var result string

func BenchmarkStringToUppper(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = strings.ToUpper(Message)
	}
	result = r
}
