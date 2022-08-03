package main

import "fmt"

func f(xp *int) {
	var y int
	xp = &y//&はアドレス
}

func main() {
	var x int
	var xp *int = &x
	f(xp)
	*xp = 100//*は格納物
	fmt.Println(x)//xpにはxのアドレスが入っているので値も連動して変化する
}
