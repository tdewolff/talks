package main

func f() *int {
    i := 6
    return &i
}

var I int = 10
func g() *int {
    return &I
}

func main() {
    _ = f()
    _ = g()
}
