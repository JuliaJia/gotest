package test01

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func Test1() {
	fmt.Println("我是来自github的Test1！")
}

func Test2() {
	var (
		host string
		port int
		h    bool
		help bool
	)
	flag.StringVar(&host, "H", "127.0.0.1", "连接地址")
	flag.IntVar(&port, "P", 22, "连接端口")
	flag.BoolVar(&h, "h", false, "帮助")
	flag.BoolVar(&help, "help", false, "帮助")
	flag.Usage = func() {
		fmt.Println("usage: testflag [-H 127.0.0.1] [-P 22]")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println(host, port, h, help)
	if h || help {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println(flag.NArg())
	fmt.Printf("%#v\n", flag.Args())
}

func TestLog() {
	log.SetFlags(log.Flags() | log.Ldate | log.Lshortfile)
	log.SetPrefix("TestLog: ")
	log.Println("这是TestLog的第一条日志！")
	logger := log.New(os.Stdout, "Logger:", log.Flags())
	logger2 := log.New(os.Stdout, "Logger2:", log.Flags())
	logger.Println("我是Logger日志")
	logger2.Println("我是Logger2日志")
}

func TestTime() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
}
