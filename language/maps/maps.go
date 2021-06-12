package main

import "fmt"

func main() {

	newMap := make(map[string]int, 4)
	newMap["A"] = 1
	newMap["B"] = 2
	newMap["C"] = 3
	newMap["D"] = 4

	fmt.Println("newMapA", newMap["A"])
	fmt.Println(newMap, newMap)

	newMap2 := map[string]string{
		"A": "Alpha",
		"O": "Omega",
	}

	fmt.Println(newMap2)
	v, okay := newMap2["O"]
	if !okay {
		fmt.Println("Not found newMap2 with Index O")
	} else {
		fmt.Println("Found value", v)
	}

	for key, value := range newMap2 {
		fmt.Printf("%v:%s,", key, value)
	}

	for key := range newMap2 {
		fmt.Printf("%v", key)
	}

	mappy := make(map[string]int)
	mappy["ex1"] = 21
	mappy["ex2"] = 43
	mappy["ex3"] = 75
	fmt.Printf("mappy: %v, len(%d)\n", mappy, len(mappy))
	fmt.Printf("mappy[ex1] %v\n", mappy["ex1"])
	mappy["ex4"] = 88
	delete(mappy, "ex4")
	value, present := mappy["ex3"] // value, bool. Checks existence of key
	fmt.Printf("key ex3. value:%v  present:%v\n", value, present)
}
