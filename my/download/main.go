package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "123.com"
	filename := "123.jpg"
	downloadFileProcess(url, filename)

}

func downloadFile(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	n, err := io.Copy(f, r.Body)
	fmt.Println(n, err)

}

type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Current += int64(n)
	// 行首
	fmt.Printf("\r进度 %.2f%% ", float64(r.Current*10000/r.Total)*100)
	return
}

func downloadFileProcess(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	reader := &Reader{
		Reader: r.Body,
		Total:  r.ContentLength,
	}
	n, err := io.Copy(f, reader)
	fmt.Println(n, err)

}
