package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Printf("%v\n", slice)
	fmt.Print("Slice: ")
	for i := range slice {
		fmt.Print(slice[i])
	}
	fmt.Println("\n--")

	a1 := []int{1, 2, 3, 4}
	fmt.Printf("data: %v, length: %d, cap: %d\n", a1, cap(a1), len(a1))
	a2 := append(a1, 5, 6)
	fmt.Printf("data: %v, length: %d, cap: %d\n", a2, cap(a2), len(a2))

}
