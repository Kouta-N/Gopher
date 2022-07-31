package main

import "fmt"

func main() {
	slice := []struct {
		N int
	}{{N: 10}, {N: 20}}

	for _, v := range slice {
		v.N *= 10
		fmt.Printf("%v,", v)
	}
	fmt.Println(slice)//スライス自体は値渡しなので変更はない

	for i := range slice {
		slice[i].N *= 10
	}
	fmt.Println(slice)//slice自体を10倍するので値が変わる

}
