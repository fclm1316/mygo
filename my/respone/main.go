package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil) //构造请求
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request) //执行请求
	defer func() { _ = r.Body.Close() }()
	responeBody(r)
	status(r)
	header(r)
}

func responeBody(r *http.Response) {
	context, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", context)
}
func status(r *http.Response) {
	fmt.Printf("状态 字符串:%s \n", r.Status)
	fmt.Printf("状态 整型:%d\n", r.StatusCode)
}

func header(r *http.Response) {
	fmt.Printf("%s", r.Header.Get("Content-type"))
}

/*func encoding(r *http.Response) {

}
*/
