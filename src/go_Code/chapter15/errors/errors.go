package main

import (
	"errors"
	"fmt"
	"log"
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
func ErrorsCase() {
	err := errors.New("程序错误")
	fmt.Println(err)
	var a, b int = -1, -2
	res, err := sum(a, b)
	fmt.Println(res, err)
	if err != nil {
		log.Println(err)
		cusError, ok := err.(customerError)
		if ok {
			fmt.Println("打印自定义错误信息：", cusError.Code, cusError.Message)
		}
	}
}
func sum(a, b int) (int, error) {
	if a <= 0 && b <= 0 {
		return 0, GetError("500", "参数都小于等于0")
	}
	return a + b, nil
}
func main() {
	ErrorsCase()
}
