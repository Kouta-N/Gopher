package main

func main() {
	//パッケージ変数はパッケージ内にて有効。関数を跨いでも使用できる。
	//go routine では競合が発生しやすいので避けるべき
	var Hoge string//大文字はpublic
	var hoge string//小文字はprivate
}
