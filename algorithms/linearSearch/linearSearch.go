package main

import "fmt"

func linearSearch(dataList []int, key int) bool {
	for _, item := range dataList {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	searchArray := []int{1, 23, 56, 102, 245}
	var key = 22
	result := linearSearch(searchArray, key)
	if result {
		fmt.Printf("%v - Searched Array %v and found Key %d", result, searchArray, key)
	} else {
		fmt.Println("Key not found in given array")
	}

}
