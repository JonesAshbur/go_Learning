package model

import "fmt"

type person struct {
	Name   string
	age    int
	salary float64
}

// 工厂模式函数
func Newperson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了访问age和salary，编写Set和Get方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		(*p).age = age
	} else {
		fmt.Println("年龄输入不合法")
	}
}

func (p *person) GetAge() (age int) {
	return (*p).age
}

func (p *person) SetSalary(salary float64) {
	if salary > 0 {
		(*p).salary = salary
	} else {
		fmt.Println("薪水输入不合法")
	}
}

func (p *person) GetSalary() (salary float64) {
	return (*p).salary
}
