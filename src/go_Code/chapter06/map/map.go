package main

import (
	"fmt"
	"sort"
)

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

	// map的遍历
	for key, value := range studentInformation {
		fmt.Println("key:", key)
		for key, value := range value {
			fmt.Printf("\tkey:%v,value:%v\n", key, value)
		}
	}

	// 获取map中键值对的个数
	fmt.Println(len(names))

	// map切片，切片的数据类型如果是map，称为slice of map，map的个数可以动态变化

	heroes := make([]map[string]string, 2)
	if heroes[0] == nil {
		heroes[0] = make(map[string]string, 2)
		heroes[0]["name"] = "牛魔王"
		heroes[0]["age"] = "500"
	}
	newhero := map[string]string{
		"name": "红孩儿",
		"age":  "200",
	}
	heroes = append(heroes, newhero)
	fmt.Println(heroes)

	// map排序
	// map中的key是无序的
	mapDemo := make(map[int]int, 10)
	mapDemo[9] = 1
	mapDemo[0] = 5
	mapDemo[4] = 7
	mapDemo[7] = 9
	mapDemo[1] = 6
	mapDemo[3] = 2
	fmt.Println(mapDemo)
	// 按照map的key排序输出
	// 1.先将map的key放入切片中
	// 2.对切片进行排序
	// 3.遍历切片，按照key来输出map的value
	var keys []int
	for index := range mapDemo {
		keys = append(keys, index)
	}
	// 排序
	sort.Ints(keys)
	fmt.Println(keys)
	for _, keys := range keys {
		fmt.Printf("mapDemo[%v]=%v\n", keys, mapDemo[keys])
	}

	// map注意事项：
	// 1.map是引用类型，遵守引用类型传递机制
	// 2.map达到一定容量时会自动扩容，不会发生panic，可以的动态增长键值对
	// 3.map的value也经常使用struct类型
}
