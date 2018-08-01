package benc_test

import (
	"strings"
	"fmt"
	"bytes"
	"testing"
)

const x = "hello"

func withPlus() string {
	return x + x + x + x + x + x + x + x + x + x
}

func withSprintf() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", x, x, x, x, x, x, x, x, x, x)
}

func withBuffer() string {
	bb := &bytes.Buffer{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}

func withStringBuilder() string {
	bb := &strings.Builder{}
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	bb.WriteString(x)
	return bb.String()
}

var result string

func BenchmarkWithPlus(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = withPlus()
	}
	result = r
}

func BenchmarkWithSprintf(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = withSprintf()
	}
	result = r
}

func BenchmarkWithBuffer(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = withBuffer()
	}
	result = r
}

func BenchmarkWithStringBuilder(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = withStringBuilder()
	}
	result = r
}
