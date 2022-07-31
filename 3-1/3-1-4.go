package main

import "fmt"

func main() {
	ns := []int{10,20,30,40,50}
	println(ns[3])
	println(len(ns))
	ns = append(ns,60,70)
	fmt.Printf("%v,", ns)
	println(cap(ns))
}
