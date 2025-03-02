package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type customerError struct {
	Code    string
	Message string
	Time    time.Time
}

func (err customerError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s, Time: %s", err.Code, err.Message, err.Time.Format("2006-01-02 15:04:05Z07:00"))
}

func GetError(code, message string) error {
	return customerError{
		Code:    code,
		Message: message,
		Time:    time.Now(),
	}
}
func sum(a, b int) (int, error) {
	if a <= 0 && b <= 0 {
		return 0, GetError("500", "参数都小于等于0")
	}
	return a + b, nil
}

// 打印报错位置
func init() {
	log.SetFlags(log.Llongfile)
	log.SetOutput(os.Stderr)
}
func LogCase() {
	var a, b int = -2, -3
	_, err := sum(a, b)
	if err != nil {
		log.Println(err)
	}
	log.Printf("a: %d, b: %d,两数求和发生错误：%s\n", a, b, err)
	//打印错误并且退出程序
	log.Fatalf("a: %d, b: %d,两数求和发生错误：%s\n", a, b, err)
}
func main() {
	LogCase()
}
