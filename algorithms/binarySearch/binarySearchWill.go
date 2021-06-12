package main

import "fmt"

// Is !Mid >= to our desired integer
// If the Midpoint Greater than X
// If it is then the endInt = MidPoint
// Else Midpoint is the start

func binarySearch(arrLength int, callback func(int) bool) int {
	startInt, endInt := 0, arrLength
	for startInt < endInt {
		mid := startInt + (endInt-startInt)/2
		if !callback(mid) {
			startInt = mid + 1
		} else {
			endInt = mid
		}
	}

	return startInt
}

// Is 28 >= 10 ==> True.. Is the Midpoint > x

func main() {

	arr := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 45
	arrLength := len(arr)
	i := binarySearch(arrLength, func(i int) bool {
		return arr[i] >= x
	})
	fmt.Println(i)

	if i < len(arr) && arr[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, arr)
	} else {
		fmt.Printf("%d not found in %v\n", x, arr)
	}

}
