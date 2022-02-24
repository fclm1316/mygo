package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

func main() {
	// post form
	// post json
	postForm()
	fmt.Println("----------------------")
	postJson()
	fmt.Println("----------------------")
	postFile()
	jar, _ := cookiejar.New(nil)
	login(jar)
	center(jar)
}

func postForm() {
	data := make(url.Values)
	data.Add("name", "bat")
	data.Add("age", "14")
	// 转换
	pyload := data.Encode()
	//fmt.Printf("%v",reflect.TypeOf(pyload))
	//fmt.Printf("%v",reflect.TypeOf(data))
	r, _ := http.Post(
		"http://httpbin.org/post",
		"application/x-www-form-urlencoded", //form 表单
		strings.NewReader(pyload),
	)
	defer func() { _ = r.Body.Close() }()
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)

}

type User struct {
	Name string
	Age  int8
}

func newUser(Name string, Age int8) *User {
	return &User{
		Name: Name,
		Age:  Age,
	}

}
func postJson() {
	u := newUser("lili", 88)
	pyload, _ := json.Marshal(u)
	//fmt.Printf("%v",reflect.TypeOf(pyload))
	//fmt.Printf("%v",reflect.TypeOf(data))
	r, _ := http.Post(
		"http://httpbin.org/post",
		"application/json", //json 表单
		bytes.NewReader(pyload),
	)
	defer func() { _ = r.Body.Close() }()
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func postFile() {
	//二进制流上传
	body := &bytes.Buffer{}
	write := multipart.NewWriter(body)
	_ = write.WriteField("words", "123")

	update1Write, _ := write.CreateFormFile("updatefile1", "filename1")
	uploadFile1, _ := os.Open("filename1")
	defer func() { _ = uploadFile1.Close() }()
	_, _ = io.Copy(update1Write, uploadFile1)

	update2Write, _ := write.CreateFormFile("updatefile2", "filename2")
	uploadFile2, _ := os.Open("filename2")
	defer func() { _ = uploadFile2.Close() }()
	_, _ = io.Copy(update2Write, uploadFile2)

	_ = write.Close()

	fmt.Printf("%s\n", write.FormDataContentType())
	fmt.Printf("%s", body.String())
	r, _ := http.Post("http://httpbin.org/post",
		write.FormDataContentType(),
		body)
	defer func() { _ = r.Body.Close() }()
	connect, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", connect)
}

func login(jar http.CookieJar) {
	client := http.Client{
		Jar: jar,
	}
	r, _ := client.PostForm(
		"http://localhost:8080/login",
		url.Values{"username": {"haha"}, "password": {"123456"}},
	)
	defer func() { _ = r.Body.Close() }()
	_, _ = io.Copy(os.Stdout, r.Body)

}

func center(jar http.CookieJar) {
	client := &http.Client{
		Jar: jar,
	}
	r, _ := client.Get("http://localhost:8080/center")
	defer func() { _ = r.Body.Close() }()
	_, _ = io.Copy(os.Stdout, r.Body)
}
