package main

import "fmt"

func swap(x, y int) (x2, y2 int) {
	y2, x2 = x, y
	return //名前つき戻り値のおかげでいちいち定義しなくてもよい
	//何をreturnしているのか分かりにくくなるのであまり実務では使用しない
}

func main() {
	a := 20
	b := 10
	fmt.Println(swap(a, b))
}
