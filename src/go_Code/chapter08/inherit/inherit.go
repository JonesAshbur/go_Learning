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

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}
type Phone struct {
	Goods
	Brand
}

// 也可以嵌套匿名结构体的指针
type Pad struct {
	*Goods
	*Brand
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
	// 2.访问简化：b.A.name => b.name,优先调用自己的属性，如果没有则查找匿名结构体，如果没有匿名结构体则报错
	// 3.当结构体和匿名结构体有相同的字段或方法时，编译器采用就近访问原则，如希望访问匿名结构体的字段或方法，可以通过匿名结构体名来区分（b.A.xxx）
	// 4.结构体嵌入两个或多个匿名结构体，如果两个匿名结构体有相同的字段和方法时，同时结构体本身没有同名的字段和方法，在访问时，就必须明确匿名结构体名字，否则报错
	// 5.如果一个结构体嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么访问组合结构体的字段或方法时，必须带上结构体名字（b.a.name)a是结构体b中的一个有名结构体
	// 6.嵌套匿名结构体后，也可以在创建结构体变量时，直接指定各个匿名结构体字段的值

	phoneList_01 := Phone{Goods{"Phone__1", 8999.9}, Brand{"Apple", "America"}} //字段顺序固定
	phoneList_02 := Phone{
		Goods{
			Name:  "Phone__2",
			Price: 5999.9, //字段顺序不固定
		},
		Brand{
			Name:    "Xiaomi",
			Address: "China",
		},
	}
	fmt.Println(phoneList_01)
	fmt.Println(phoneList_02)

	padList_01 := Pad{
		&Goods{
			Name:  "Pad_01",
			Price: 1999.9, //字段顺序不固定
		},
		&Brand{
			Name:    "Huawei",
			Address: "China",
		},
	}
	padList_02 := Pad{
		&Goods{
			Name:  "Pad_02",
			Price: 3999.9, //字段顺序不固定
		},
		&Brand{
			Name:    "Google",
			Address: "America",
		},
	}
	fmt.Println(*(padList_01.Goods), *padList_01.Brand)
	fmt.Println(*padList_02.Goods, *padList_02.Brand)
}
