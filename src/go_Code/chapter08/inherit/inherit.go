package main

import "fmt"

// 共同特征
type Student struct {
	Name  string
	Age   int
	Score float64
}

// Pupil,Graduate共有的方法
// 展示信息
func (stu *Student) ShowInfo() {
	fmt.Printf("name:%v,age:%v,score:%v\n", (*stu).Name, (*stu).Age, (*stu).Score)
}

// 设置分数
func (stu *Student) SetScore(set_score float64) (return_score float64) {
	(*stu).Score = set_score
	return return_score
}

// 格式化字符串
func (stu *Student) String() string {
	str := fmt.Sprintf("姓名：%v\n年龄：%v\n分数：%v\n\n", (*stu).Name, (*stu).Age, (*stu).Score)
	return str
}

type Pupil struct {
	Student //嵌入student匿名结构体
}

func (p *Pupil) testing() {
	fmt.Println("小学生考试中...(特性保留)")
}

type Graduate struct {
	Student //嵌入student匿名结构体
}

func (p *Graduate) testing() {
	fmt.Println("大学生考试中...(特性保留)")
}
func main() {

	// 继承可以解决代码复用问题
	// 当多个结构体存在相同的字段和方法时，可以在这些结构体中抽象出结构体，该结构体中定义相同的属性和方法
	// 其他结构体不需要重新定义这些属性和方法，只需要嵌套一个该结构体的匿名结构体即可

	// 小学生演示
	pupil := &Pupil{}
	pupil.Student.Name = "jones"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(89.98)
	pupil.Student.ShowInfo()
	fmt.Println(pupil) //pupil已经是一个地址，所以不需要加&
	// 大学生演示
	graduate := &Graduate{}
	graduate.Student.Name = "jack"
	graduate.Student.Age = 26
	graduate.testing()
	graduate.Student.SetScore(99.18)
	graduate.Student.ShowInfo()
	fmt.Println(graduate) //graduate已经是一个地址，所以不需要加&

	// 继承注意事项：
	// 1.结构体可以使用匿名结构体的所有字段和方法，即首字母大写或者小写的字段和方法都可以使用
}
