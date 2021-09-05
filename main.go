package main

import (
	"fmt"
	"strings"
)

func main() {

	// golang mapのゼロ値の活用
	counts := make(map[string]int)
	str := "dog dog dog cat dog"
	for _, s := range strings.Split(str, " ") {
		counts[s]++
	}
	for a, c := range counts {
		fmt.Println("The number of", a, "is", c)
	}
}
