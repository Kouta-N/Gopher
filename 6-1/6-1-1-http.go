package main

import (
	"fmt"
	"net/http"
)

type HTTPClient interface { ///インタフェース型はメソッドの呼び出しだけができる
	Do(req *http.Request) (*http.Response, error)
}

func main() {
	var cli *http.Client = http.DefaultClient
	var cli2 HTTPClient = cli
	fmt.Println(cli2)
}
