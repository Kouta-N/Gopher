package main

import "fmt"

func main() {
	var names = make(map[int]string)

	names[0] = "Suzuki" // インデックスで挿入
	names[1] = "Yoshida"
	fmt.Println(names[1]) // Yoshida

	var m = make(map[string]int)
	m["ab"] = 10 //""でないと文字のキーは指定できない
	fmt.Println(m["ab"])
}
