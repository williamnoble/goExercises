package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v,kind:%v\n", t, t.Kind())
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64,value is %f\n", float64(v.Float()))
	}
}

func main() {
	//str := `{"name":"aaron","age":25}`

	// TypeOf
	fmt.Println("1、TypeOf")
	var a int64 = 100
	reflectType(a)
	var b float64 = 9.62
	reflectType(b)
	var c = person{}
	reflectType(c)

	// ValueOf
	fmt.Println("2、ValueOf")
	reflectValue(b)

	// SetValueOf
	fmt.Println("3、SetValueOf")
	reflectSetValue2(&a)
	fmt.Println(a)

	fmt.Println("4、TypeOf结构体")
	per := person{
		Name: "aaron",
		Age:  25,
	}
	t := reflect.TypeOf(per)
	fmt.Println(t.Name(), t.Kind())
	// 遍历结构体字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s,index:%d,type:%v,jsonTag：%v\n", field.Name, field.Index, field.Type, field.Tag)
	}
	// 通过字段名称获取指定结构体字段信息
	if ageField, ok := t.FieldByName("Age"); ok {
		fmt.Printf("name:%s,index:%d,type:%v,jsonTag：%v\n", ageField.Name, ageField.Index, ageField.Type, ageField.Tag.Get("json"))
	}
}
