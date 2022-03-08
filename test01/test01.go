package test01

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
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
	year, month, day := 1990, time.February, 1
	time1990 := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	fmt.Println(time1990.Format("2006-01-02 15:04:05"))

}

func TestBase64() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("我是林克！")))
	txt, _ := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString([]byte("我是林克！")))
	fmt.Println(string(txt))
	fmt.Println(base64.URLEncoding.EncodeToString([]byte("我是林克！")))
}

func TestHex() {
	fmt.Printf("%x\n", []byte("我是林克！"))
	fmt.Println(hex.EncodeToString([]byte("我是林克！")))
	ds, _ := hex.DecodeString("e68891e698afe69e97e5858befbc81")
	fmt.Println(string(ds))

}

func randString(n int) (string, string) {
	rt1 := make([]byte, 0, n)
	rt2 := make([]byte, n, n)
	chars := "01234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~`!@#$%^&*()_+-=/\\|\"'?,.<>"
	for i := 0; i < n; i++ {
		rt1 = append(rt1, chars[rand.Intn(len(chars))])
		rt2[i] = chars[rand.Intn(len(chars))]
	}
	return string(rt1), string(rt2)
}

func md5String(text string, salt string) string {
	bytes := []byte(salt)
	bytes = append(bytes, ':')
	bytes = append(bytes, []byte(text)...)
	return fmt.Sprintf("%x", md5.Sum(bytes))
}

func TestHash() {
	fmt.Printf("%x\n", md5.Sum([]byte("我是林克！")))
	hasher := md5.New()
	hasher.Write([]byte("我是林克！"))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	salt1, salt2 := randString(6)
	fmt.Println(salt1, salt2)
	fmt.Println("我是林克！", salt1, md5String("我是林克！", salt1))
}

func TestSha() {
	fmt.Printf("%x\n", sha1.Sum([]byte("我是林克！")))
	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte("我是林克！"))
	fmt.Println(hex.EncodeToString(sha256Hasher.Sum(nil)))

}

func TestCmd() {
	cmd := exec.Command("/sbin/ping", "-c", "2", "www.baidu.com")
	bytes, err := cmd.Output()
	fmt.Println(string(bytes), err)
}

func NgFlow(count int, dict map[int]float64) map[int]float64 {
	for k := range dict {
		fmt.Println(dict[k])
	}
	return dict
}
