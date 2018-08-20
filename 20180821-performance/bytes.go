package main

import "fmt"

func main() {
    // START OMIT
    b1 := []byte{}
    b2 := []byte("")
    b3 := append([]byte{}, 'c')

    fmt.Println(len(b1), cap(b1))
    fmt.Println(len(b2), cap(b2))
    fmt.Println(len(b3), cap(b3))
    // END OMIT
}
