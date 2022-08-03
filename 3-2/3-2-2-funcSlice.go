package main

import "fmt"

func main() {
	// fs := make([]func() string, 3) //文字型を返す関数を2つスライスで用意
	// //makeはlengthを指定しないとcapのみを定義する。
	// fs[0] = func() string { return "hoge" }
	// fs[1] = func() string { return "fuga" }
	// fs[2] = func() string { return "gaga" }
	// for _, f := range fs {
	// 	fmt.Println(f())
	// }

	// fs := make([]func(), 3)
	// for i := range fs {
	// 	fs[i] = func() { fmt.Println(i) }
	// }
	// for _, f := range fs {
	// 	f()//iの値は2になったところで終わってるので出力は全て2
	// }

	fs := make([]func(), 3)
	for i := range fs {
		i := i                 //スコープが変わり、変数が毎回新規作成されるようになった
		fmt.Printf("%p\n", &i) //iのポインタは毎回変わる
		fs[i] = func() { fmt.Println(i) }
	}
	for _, f := range fs {
		f()
	}

	for i := range fs {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
}
