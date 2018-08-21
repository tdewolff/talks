package main

import (
	"bytes"
	"fmt"
)

// START OMIT
func GetWord(b, word []byte) []byte {
	for start := 0; start < len(b)-len(word); start++ {
		end := start + len(word)
		if bytes.Equal(b[start:end], word) {
			return b[start:end:end] // <<< Set capacity!
		}
	}
	return nil
}

func main() {
	b := []byte("There was no collusion. Everybody knows there\nwas no collusion.")
	word := GetWord(b, []byte("collusion"))
	word = append(word, []byte(", FAKE NEWS!!")...)

	fmt.Println(string(b))
}

// END OMIT
