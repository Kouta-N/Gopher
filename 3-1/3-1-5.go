package main

import "fmt"

func main() {
	a := []int{10,20}
	fmt.Println(a,cap(a))

	b := append(a,30)
	a[0] = 100
	fmt.Println(b,cap(b))//aを参照しているわけではないので、上の命令はbには反映されない
	fmt.Println(a,cap(a))
}
