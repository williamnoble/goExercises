package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int
	fmt.Println(s == nil)

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

	s10 := make([]int, 5, 10)
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s10=%v	len=%d	cap=%d \n", s1, len(s10), cap(s10))

	stringSlice := []string{"ONE", "TWO", "THREE"}
	stringSlice2 := []string{"FOUR", "FIVE"}
	stringSlice = append(stringSlice, stringSlice2...)
	fmt.Println(stringSlice)

	unsortedSlice := []int{8, 9, 7, 3, 4, 2, 1}
	// Numeric > Capital > lowercase
	sort.Ints(unsortedSlice[:])
	fmt.Printf("Sorted: %v \n", unsortedSlice)
	sort.Sort(sort.Reverse(sort.IntSlice(unsortedSlice)))
	fmt.Printf("Sorted(Reversed): %v \n", unsortedSlice)

	unsortedStringSlice := []string{"lowercase", "777", "alphabet", "London", "Paris", "NY", "123", "03"}
	sort.Strings(unsortedStringSlice)
	fmt.Printf("SortedString:%v \n", unsortedStringSlice)
	sort.Sort(sort.Reverse(sort.StringSlice(unsortedStringSlice)))
	fmt.Printf("SortedString(reversed)%v \n", unsortedStringSlice)
}
