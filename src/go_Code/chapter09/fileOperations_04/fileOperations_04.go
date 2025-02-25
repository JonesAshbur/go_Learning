package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	filePath := "E:/go_Learning/fileDemo/demo_06.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("读取文件错误：", err)
			break
		}

		// 打印调试信息（确认读取到数据）
		fmt.Printf("[DEBUG] Read line: %q\n", str)

		for _, value := range str {
			switch {
			case value >= 'a' && value <= 'z', value >= 'A' && value <= 'Z':
				count.ChCount++
			case value == ' ':
				count.SpaceCount++
			case value == '\t':
				count.SpaceCount++
			case value >= '0' && value <= '9':
				count.NumCount++
			case value == '\n':
				// 忽略换行符
			default:
				count.OtherCount++
			}
		}

		if err == io.EOF {
			break
		}
	}

	fmt.Printf("英文字符：%v，数字：%v，空格/制表符：%v，其他字符：%v\n",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
