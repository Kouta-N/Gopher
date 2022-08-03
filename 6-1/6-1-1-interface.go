package main

import "fmt"

type Stringer interface {///インタフェース型はメソッドの呼び出しだけができる
	String() string
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	var s Stringer = Hex(100)
	fmt.Println(s.String())
	//インタフェースとはメソッドの集まり
}
