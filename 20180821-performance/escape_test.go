package main

import "testing"

func f() *int {
    i := 6
    return &i
}

var I int = 10

func g() *int {
    return &I
}

func BenchmarkF(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = f()
    }
}

func BenchmarkG(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = g()
    }
}
