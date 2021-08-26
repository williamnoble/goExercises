package main

import "fmt"

func main() {

	// declared but not initialized. Attempting to use will panic
	var m map[string]int

	// initialized
	m = make(map[string]int)
	m["ONE"] = 1
	m["TWO"] = 2

	// initialize a map literal
	n := map[string]int{
		"a": 1,
		"b": 2,
	}
	_ = n

	// create a map type
	type wrap map[string]interface{}
	a := wrap{"error": "error"}
	fmt.Println(a)

	// read a map
	println("Value of first map: ", m["ONE"])

	// range: KEY
	fmt.Print("Keys: ")
	for k := range m {
		fmt.Printf("%s, ", k)
	}
	fmt.Println("")
	// range: KEY VALUE
	for i, v := range m {
		println("KEY | VALUE = : ", i, v)
	}
	m["THREE"] = 3

	// is a value in a map? panic if we don't use ok and m["THREE"] not in map
	if v, ok := m["THREE"]; ok {
		println("Map with KEY 'THREE' - okay? ", v, ok)
	}
}
