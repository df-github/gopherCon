package profile_test

import (
	"testing"
	"strings"
)

var result string
var message = "Education is what remains after one has forgotten what one has learned in school."
func BenchmarkStringToUppper(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = strings.ToUpper(message)
	}
	result = r
}
