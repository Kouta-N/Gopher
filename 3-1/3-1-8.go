package main

import (
	"fmt"
	"strings"
)

func main() {
	counts := map[string]int{}
	str := "dog dog dog cat dog"
	for _, s := range strings.Split(str, " ") {
		counts[s]++
	}
	fmt.Println(counts)
}
