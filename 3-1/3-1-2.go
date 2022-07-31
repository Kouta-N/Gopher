package main

func main(){
	// var ns []int //[]はスライス、intは要素

	// var m map[string]int //mapは連想配列で、キーはstring,intが要素

	// var p struct{//構造体の定義方法
	// 	name string
	// 	age int
	// }

	p2 := struct{//構造体リテラル(初期化も済ます)
	name string
	age int
	}{
		name:"Gopher",
		age:10,
	}

	p2.age++ //構造体へのアクセスは.を使用する
	println(p2.name,p2.age);

	var ns [5]int//配列の要素数を宣言し、ゼロ値で初期化(要素数が異なると別の型になるので代入などは不可能)
	//Goの配列は可変長配列

	ns2 := [...]int{10,20,30,40,50}//勝手に型も要素数(5)も定義してくれる 
	ns3 := [...]int{5:50,10:100} //5番目が50,10番目が100の要素数11の配列を作成(指定していないところはゼロ値)

}
