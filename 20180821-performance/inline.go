package main

import "fmt"

func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

func minSlice(c ...int) int {
    lowest := c[0]
    for _, x := range c[1:] {
        lowest = min(lowest, x)
    }
    return lowest
}

func main() {
    lowest := minSlice(9,2,5,8,-3,135)
    fmt.Println("Lowest:", lowest)
}
