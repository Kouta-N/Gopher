package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := &Person{Name: "tenntenn", Age: 32}
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}
