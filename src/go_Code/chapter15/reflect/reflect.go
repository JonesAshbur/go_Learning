package main

import (
	"errors"
	"fmt"
	"reflect"
)

func ReflectCase() {
	type user struct {
		ID    int64
		Name  string
		Hobby []string
	}
	type outUser struct {
		ID    int64
		Name  string
		Hobby []string
	}
	u := &user{
		ID:    1,
		Name:  "ash",
		Hobby: []string{"football", "coding", "reading"},
	}
	out := outUser{}
	//接收错误信息
	err := reflectCopy(&out, u)
	fmt.Println(err, out)
	listUser := []user{
		{ID: 1, Name: "jones", Hobby: []string{"basketball", "coding", "reading"}},
		{ID: 2, Name: "jack", Hobby: []string{"baseball", "writing", "singing"}},
		{ID: 3, Name: "lucks", Hobby: []string{"volleyball", "running", "playing"}},
	}
	list := sliceCollection(listUser, "Hobby")
	fmt.Println(list)
}

func reflectCopy(dest interface{}, source interface{}) error {
	sType := reflect.TypeOf(source)
	sValue := reflect.ValueOf(source)
	//	如果为指针类型则获取到值
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
		sValue = sValue.Elem()
	}
	dType := reflect.TypeOf(dest)
	dValue := reflect.ValueOf(dest)
	//判断目标对象类别
	if dType.Kind() == reflect.Ptr {
		dType = dType.Elem()
		dValue = dValue.Elem()
	} else {
		return errors.New("目标对象必须是struct指针")
	}
	if sValue.Kind() != reflect.Struct {
		return errors.New("原对象必须为struct或struct指针")
	}
	if dValue.Kind() != reflect.Struct {
		return errors.New("目标对象必须为struct或struct指针")
	}
	destObj := reflect.New(dType)
	for i := 0; i < dType.NumField(); i++ {
		destField := dType.Field(i)
		if sourceField, ok := sType.FieldByName(destField.Name); ok {
			if destField.Type != sourceField.Type {
				continue
			}
			value := sValue.FieldByName(destField.Name)
			destObj.Elem().FieldByName(destField.Name).Set(value)
		}
	}
	dValue.Set(destObj.Elem())
	return nil
}
func sliceCollection(slice interface{}, collection string) interface{} {
	sliceType := reflect.TypeOf(slice)
	sliceValue := reflect.ValueOf(slice)
	if sliceType.Kind() == reflect.Ptr {
		sliceType = sliceType.Elem()
		sliceValue = sliceValue.Elem()
	}
	if sliceValue.Kind() == reflect.Struct {
		value := sliceValue.FieldByName(collection)
		return value.Interface()
	}
	if sliceValue.Kind() != reflect.Slice {
		return nil
	}
	sliceType = sliceType.Elem()
	if sliceType.Kind() == reflect.Ptr {
		sliceType = sliceType.Elem()
	}
	field, _ := sliceType.FieldByName(collection)
	Stype := reflect.SliceOf(field.Type)
	sliceGather := reflect.MakeSlice(Stype, 0, 0)
	for i := 0; i < sliceValue.Len(); i++ {
		object := sliceValue.Index(i)
		if object.Kind() == reflect.Struct {
			value := object.FieldByName(collection)
			sliceGather = reflect.Append(sliceGather, value)
		}
		if object.Kind() == reflect.Ptr {
			objValue := object.Elem()
			value := objValue.FieldByName(collection)
			sliceGather = reflect.Append(sliceGather, value)
		}
	}
	return sliceGather.Interface()
}
func main() {
	ReflectCase()
}
