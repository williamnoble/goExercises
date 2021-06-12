package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("1	----- map声明 -----")
	m1 := make(map[string]int, 10) // map必须初始化才能使用

	m1["A"] = 10
	m1["H"] = 6
	m1["D"] = 5
	m1["B"] = 20
	m1["K"] = 11
	m1["F"] = 12
	m1["Z"] = 13
	m1["V"] = 23

	fmt.Println("m1", m1)
	fmt.Println("A：", m1["A"])

	// map声明并赋值
	m2 := map[string]string{
		"A": "AA",
		"B": "BB",
	}
	fmt.Println("m2", m2)

	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	fmt.Println("2	----- map key判断 -----")
	v, ok := m1["A"]
	if !ok {
		fmt.Println("无此Key值")
	} else {
		fmt.Println("存在Key值")
	}
	fmt.Println(v)

	// 删除map键值
	fmt.Println("3	----- map删除 -----")
	delete(m1, "A")
	fmt.Println("删除后map值：", m1)

	// map遍历，遍历map时的元素顺序与添加键值对的顺序无关
	fmt.Println("4	----- map遍历 -----")

	fmt.Println("map遍历1")
	for k, v := range m1 {
		fmt.Printf("key：%v	value：%d \n", k, v)
	}

	fmt.Println("map遍历2")
	for _, v := range m2 {
		fmt.Printf("value：%v \n", v)
	}

	fmt.Println("只要key map遍历")
	for k := range m2 {
		fmt.Printf("key：%v \n", k)
	}

	fmt.Println("按照key指定顺序map遍历")
	var keys = make([]string, 0, 10)
	for key := range m1 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, m1[key])
	}

	fmt.Println("4	----- 元素为map类型的切片 -----")

	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index：%d，value：%v \n", index, value)
	}
	fmt.Println("init...")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 5)
	mapSlice[0]["name"] = "aaron"
	mapSlice[0]["email"] = "aaron@qq.com"
	mapSlice[0]["address"] = "福田区"
	for index, value := range mapSlice {
		fmt.Printf("index：%d，value：%v \n", index, value)
	}

	fmt.Println("5	----- 值为切片类型的map -----")
	var sliceMap = make(map[string][]string, 3)
	fmt.Println("init...")
	key := "广东"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "中山", "广州")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
