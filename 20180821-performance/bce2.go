package main

func Reverse4(k []int) []int {
	_ = k[3]
	a := k[0]
	b := k[1]
	c := k[2]
	d := k[3]
	return []int{d, c, b, a}
}

func main() {
	k := []int{1, 2, 3, 4}
	_ = Reverse4(k)
}
