package main

func Sum(nums []int) (total int) {
	total = 0

	// Option 1: For loop
	//for i := 0; i < len(nums); i++ {
	//	total += nums[i]
	//}

	// Option 2: Range
	for _, value := range nums {
		total += value
	}

	return total
}

func SumAll(numsToSum ...[]int) ([]int) {
	//sums = make([]int, len(numsToSum))
	var sums []int

	for _, numbers := range numsToSum {
		//sums[i] = Sum(v)
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}