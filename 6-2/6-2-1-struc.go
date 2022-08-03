package main

import "fmt"

type Hoge struct { //インタフェース型はメソッドの呼び出しだけができる
	N int
}

type Fuga struct {
	Hoge //構造体に構造体を埋め込める
}

func main() {
	f := Fuga{Hoge{N: 100}}
	fmt.Println(f.N)
	fmt.Println(f.Hoge.N) //アクセス可能
}
