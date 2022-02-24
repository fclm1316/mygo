package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := CopyFile("dst.txt", "src.txt")
	if err != nil {
		fmt.Printf("copy faile %v", err)
		return
	}
	fmt.Println("copy done")

	flag.Parse() //解析命令行参数
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
		return
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed , err %v \n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			fmt.Fprintf(os.Stdout, "%s", buf)
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func CopyFile(dstname, srcnmae string) (written int64, err error) {
	// 拷贝文件 copy
	src, err := os.Open(srcnmae)
	if err != nil {
		fmt.Printf("open %s failed : err: %v", srcnmae, err)
		return
	}
	defer src.Close()

	dst, err := os.Open(dstname)
	if err != nil {
		fmt.Printf("open %s failed : err: %v", dstname, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
