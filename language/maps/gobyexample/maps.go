package main

import "fmt"

func main() {
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