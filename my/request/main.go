package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func printBody(r *http.Response) {
	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func requetsByParams() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values) //参数
	params.Add("name", "haha")
	params.Add("age", "18")
	request.URL.RawQuery = params.Encode()
	request.Header.Add("user-agent", "chrome") //请求头
	fmt.Printf("%s", params.Encode())

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	printBody(r)
}

func main() {
	requetsByParams()
}
