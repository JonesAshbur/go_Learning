package model

// 首字母大写
type Student_01 struct {
	Name string
	Age  int
}

// 首字母小写
type student_02 struct {
	Name string
	age  int
}

// 此时student结构体是私有的，不能被其他包调用
// 通过工厂模式解决
func NewStudent(name string, age int) *student_02 {
	return &student_02{
		Name: name,
		age:  age,
	}

}

// age是私有的，要在其他包访问：
func (s *student_02) GetAge() int {
	return (*s).age
}
