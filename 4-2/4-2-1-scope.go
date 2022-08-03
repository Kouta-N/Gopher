package main

func f(){//ブロックスコープは関数やif,for文のみで有効なスコープ
	n := 100//ブロック外では無効となる
	println(n)
}

func main() {

}
