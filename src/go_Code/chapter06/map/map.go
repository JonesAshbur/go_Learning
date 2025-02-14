package main

import "fmt"

func main() {

	// map是key，value数据结构，又称为字段或者关联数组，类似其他语言的集合
	// map中的key可以是多种数据类型，bool，int，float，string，指针，channel，接口，结构体，数组
	// slice,map,function不可以做key
	// valuetype的值通常是数字（整数，浮点数），string，map，struct
	// 声明map不会分配内存，初始化需要make，分配内存后才能赋值和使用
	// map中的key不能重复，如果重复了以最后的key，value为准
	// value可以重复
	// map中的key，value是无序的
	var str1 map[string]string = make(map[string]string, 2)
	str1["no1"] = "jones"
	str1["no2"] = "jack"
	str1["no3"] = "ashbur"
	for index, value := range str1 {
		fmt.Printf("index=%v,value=%v\n", index, value)
	}

	// map使用方式
	cities := make(map[string]string)
	cities["first"] = "beijing"
	cities["second"] = "tianjin"
	cities["third"] = "shanghai"
	fmt.Println(cities)

	names := map[string]string{
		"1": "mary",
		"2": "jacky",
		"3": "alfred",
	}
	names["4"] = "jone"
	fmt.Println(names)

	studentInformation := make(map[string]map[string]string)
	studentInformation["ifo_01"] = make(map[string]string, 2)
	studentInformation["ifo_01"]["name"] = "tom"
	studentInformation["ifo_01"]["sex"] = "male"

	studentInformation["ifo_02"] = make(map[string]string, 2)
	studentInformation["ifo_02"]["name"] = "ashbur"
	studentInformation["ifo_02"]["sex"] = "fmale"

	studentInformation["ifo_03"] = make(map[string]string, 2)
	studentInformation["ifo_03"]["name"] = "elon"
	studentInformation["ifo_03"]["sex"] = "male"

	fmt.Println(studentInformation["ifo_01"])

	// map的CRUD操作
	// map的增改操作对应定义key value，修改key对应的value
	// map的删除操作delete(map,"key")，delete是一个内置函数，如果key存在，就删除key-value，如果key不存在就不操作也不报错
	// 通过遍历删除所有key，或者重新make一个新空间，让原有数据直接被gc回收
	delete(studentInformation, "ifo_01")
	fmt.Println(studentInformation)
	// studentInformation = make(map[string]map[string]string)
	// fmt.Println(studentInformation)

	// map查找操作
	// res返回true，false
	// val返回查找的key对应的value
	val, res := names["1"]
	if res {
		fmt.Println("value=", val)
	} else {
		fmt.Println("没有找到")
	}
}
