package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

var LoginUrl, CodeUrl string

func init() {
	url := "http://127.0.0.1:8081"
	login := "login"
	webcode := "captcha/captchaImage?type=math"
	LoginUrl = url + login
	CodeUrl = url + webcode
}

func main() {
	Fdata := NewLoginData("1899998888", "123456", "0", "false")
	JsonLogin(Fdata)
}

func JsonLogin(datajson *LoginData) {
	pyload, _ := json.Marshal(datajson)
	for i := 0; i <= 20; i++ {
		jar, _ := cookiejar.New(nil)
		client := http.Client{
			Jar: jar,
		}
		rPost, _ := client.Post(LoginUrl, "application/json", bytes.NewReader(pyload))
		rGet, _ := client.Get(CodeUrl)

		defer func() {
			_ = rPost.Body.Close()
			_ = rGet.Body.Close()
		}()

		content, _ := ioutil.ReadAll(rPost.Body)

		fmt.Printf("%s\n", content)

	}
}

type LoginData struct {
	UserName     string
	Password     string
	ValidateCode string
	RememberMe   string
}

func NewLoginData(Name string, Password string, Code string, Reme string) *LoginData {
	return &LoginData{
		UserName:     Name,
		Password:     Password,
		ValidateCode: Code,
		RememberMe:   Reme,
	}
}
