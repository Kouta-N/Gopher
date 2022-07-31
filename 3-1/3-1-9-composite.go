package main

import "fmt"

type composite struct{
	A struct{
		N int
	}
}

func main() {//コンポジット型の定義
	var a [][]int
	var b map[string][]int
	var c composite
	fmt.Println(a,b,c)
}
