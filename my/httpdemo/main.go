package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }() //关闭

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func post() {
	r, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }() //关闭

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)

}
func put() {
	//没有直接提供方法,post,get 二次封装
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil) //构造请求
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request) //执行请求
	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func del() {
	//没有直接提供方法,post,get 二次封装
	request, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil) //构造请求
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request) //执行请求
	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func main() {
	get()
	post()
	put()
	del()

}
