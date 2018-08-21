package main

func f() *int {
	i := 6
	return &i
}

var j int = 10

func g() *int {
	return &j
}

func main() {
	_ = f()
	_ = g()
}
