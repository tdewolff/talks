package main

func Reverse4(k []int) []int {
    d := k[3]
    c := k[2]
    b := k[1]
    a := k[0]
    return []int{d, c, b, a}
}

func main() {
    k := []int{1, 2, 3, 4}
    _ = Reverse4(k)
}
