package main

// START OMIT
func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

func minThree(a, b, c int) int {
    tmp := min(a, b)
    return min(tmp, c)
}

func minSlice(c0 int, c ...int) int {
    lowest := c0
    for _, x := range c {
        lowest = min(lowest, x)
    }
    return lowest
}
// END OMIT

func main() {
}
