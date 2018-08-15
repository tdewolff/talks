package main

import (
    "testing"
    "bytes"
    "strings"
)

// START OMIT
func BenchmarkStrings(b *testing.B) {
    s := ""
    for i := 0; i < b.N; i++ {
        s += "x"
    }
}

func BenchmarkStringsBuilder(b *testing.B) {
    s := strings.Builder{}
    for i := 0; i < b.N; i++ {
        s.WriteString("x")
    }
}

func BenchmarkBytesBuffer(b *testing.B) {
    s := bytes.Buffer{}
    for i := 0; i < b.N; i++ {
        s.WriteString("x")
    }
}
// END OMIT
