package main

// func main(){
// 	var sum int
// 	sum = 14
// 	avg := sum / 3
// 	println(avg)
// 	if avg > 4.5{ //キャストしなければエラーになる
// 		println("good")
// 	}
// }

func main(){
	var sum int
	sum = 14
	avg := float64(sum) / 3
	println(avg)
	if avg > 4.5{ //キャストしなければエラーになる
		println("good")
	}
}
