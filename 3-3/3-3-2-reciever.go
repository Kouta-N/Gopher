package main

type MyInt int //レシーバとはメソッドのための変数

func (n *MyInt) Inc() { //メソッドとはユーザー定義関数を使用する関数
	*n++//ポインタ型を渡さなければ、値のコピーで終わってしまう
}

func main() {
	var n MyInt
	println(n)
	n.Inc()
	println(n)
}
