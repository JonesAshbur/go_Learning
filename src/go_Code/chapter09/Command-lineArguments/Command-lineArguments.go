package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("命令行的参数个数：", len(os.Args))
	// 遍历os.Args切片，就可以得到所有的命令行输入参数值
	for index, value := range os.Args {
		fmt.Printf("args[%v]=%v\n", index, value)
	}

	// flag包解析命令行参数
	// 定义变量用于接收命令行参数值
	var user string
	var password string
	var host string
	var port int
	// 接收用户命令行输入的 -u 后面的参数值
	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&password, "password", "", "密码默认为空")
	flag.StringVar(&host, "host", "localhost", "主机名默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口默认为3306")
	flag.Parse()
	fmt.Printf("user=%v,password=%v,host=%v,port=%v", user, password, host, port)
}
