package main

import "fmt"

type integer int

func main() {
	x := 21
	fmt.Printf("starting x: %d\n", x) // => x: 21

	addWithoutPointer(x)
	fmt.Printf("(fn w/o pointer) x: %d\n", x) // => x: 21

	addWithPointer(&x)
	fmt.Printf("(fn w pointer) x: %d\n", x) // => x:22

	y := integer(21)
	y.addWithPointer()
	fmt.Printf("(value rec) y: %d\n", y) // => y: 22

	y = integer(21)
	newValue := int(y.addWithoutPointer())
	fmt.Printf("(value rec) y: %d - %d\n", y, newValue) // => y: 21 - 22

	// copy
	simpleString := "Hello"
	pointerTo := &simpleString
	pointedDeference := *pointerTo
	pointerDereferenceTo := *&simpleString

	fmt.Printf("basic value of the simple string is: %s\n", simpleString)
	fmt.Printf("location in memory: %v\n", pointedDeference)
	fmt.Printf("location in memory: %v\n", pointerTo)
	fmt.Printf("type of value : %T\n", pointerTo)
	fmt.Printf("value of dereferenced pointer: %v", pointerDereferenceTo)
	fmt.Println("\n\nkorea")

	b := new(int)
	b1 := new(bool)
	fmt.Printf("%T\n", b)  // *int
	fmt.Printf("%T\n", b1) // *bool
	fmt.Println(*b)        // 0
	fmt.Println(*b1)
	var c map[string]int
	c = make(map[string]int, 10)
	c["onehundred"] = 100
	fmt.Println(c)

	//3
	var num int
	fmt.Println(&num) // => address of num
	var ptr *int
	ptr = &num // ptr points to address of num
	*ptr = 55  // dereference ptr, value of num = 55
	fmt.Println(num)
}

func addWithoutPointer(i int) int {
	i = i + 1
	return i
}

func addWithPointer(i *int) int {
	*i = *i + 1
	return *i // requires pointer dereference
}

// not signature not required
func (in *integer) addWithPointer() {

	*in = *in + 1
}

func (in integer) addWithoutPointer() integer {
	result := in + 1
	return result
}
