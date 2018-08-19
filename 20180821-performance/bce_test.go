package main

import "testing"

func Reverse4a(k []int) []int {
    a := k[0]
    b := k[1]
    c := k[2]
    d := k[3]
    return []int{d, c, b, a}
}

func Reverse4b(k []int) []int {
    d := k[3]
    c := k[2]
    b := k[1]
    a := k[0]
    return []int{d, c, b, a}
}

var K []int = []int{1, 2, 3, 4}

func BenchmarkBoundsA(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = Reverse4a(K)
    }
}

func BenchmarkBoundsB(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = Reverse4b(K)
    }
}
