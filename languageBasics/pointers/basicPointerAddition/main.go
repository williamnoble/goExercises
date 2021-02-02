package main

import "fmt"

type integer int

func main() {
	x := 21
	fmt.Printf("starting x: %d\n", x)

	addWithoutPointer(x)
	fmt.Printf("(fn w/o pointer) x: %d\n", x)

	addWithPointer(&x)
	fmt.Printf("(fn w pointer) x: %d\n", x)

	y := integer(21)
	y.addWithPointer()
	fmt.Printf("(value rec) y: %d\n", y)

	y = integer(21)
	newValue := int(y.addWithoutPointer())
	fmt.Printf("(value rec) y: %d - %d\n", y, newValue)

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
