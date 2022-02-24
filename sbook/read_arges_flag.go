package main

import (
	"flag"
	"fmt"
	"os"
)

var h bool
var v, V bool
var t, T bool
var q *bool
var s string
var p string
var c *string
var g int

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")
	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")
	//另一种绑定方式,空值
	flag.Bool("q", false, "suppress non-error messages during configuration testing")
	// `内容` 会在参数后面加 内容
	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	//注意c值
	flag.String("c", "etc.ini/nginx.etc.ini", "set configuration `file`")
	flag.IntVar(&g, "g", 0, "set global `type` ")
	//os.Args[0]
	flag.Usage = usage
}

func main() {
	//扫描参数列表
	flag.Parse()

	if h {
		flag.Usage()
	}

	fmt.Println()
	fmt.Printf("v:%t, V:%t, t:%t, T:%t, q:%v, s:%s, p:%s, c:%v, g:%d\n", v, V, t, T, q, s, p, c, g)
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("参数名:%v, 参数值:%v, 默认值:%v, 描述值:%v\n", f.Name, f.Value, f.DefValue, f.Usage)
	})
	fmt.Println("---------------------------")
	flag.Visit(func(f *flag.Flag) {
		fmt.Printf("参数名:%v, 参数值:%v, 默认值:%v, 描述值:%v\n", f.Name, f.Value, f.DefValue, f.Usage)
	})
	fmt.Println("---------------------------")
	// 参数  参数个数
	fmt.Println(flag.Args(), flag.Arg(0), flag.NArg())

}

func usage() {
	// os.Stderr 无缓冲标准输出 os.Stdout 缓存标准输出，遇到换行
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
    Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

    Options:
`)
	flag.PrintDefaults()
}
