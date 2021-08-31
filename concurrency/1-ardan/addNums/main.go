package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func nonConcurrentSum(numbers []int) int {
	var v int
	for _, number := range numbers {
		v += number
	}
	return v
}

func concurrentSum(goroutines int, numbers []int) int {
	var v int64
	totalNumbers := len(numbers)
	lastGoRoutine := goroutines - 1
	stride := totalNumbers / goroutines

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride
			if g == lastGoRoutine {
				end = totalNumbers
			}

			var lv int
			for _, n := range numbers[start:end] {
				lv += n
			}

			atomic.AddInt64(&v, int64(lv))
			wg.Done()
		}(g)
	}
	wg.Wait()
	return int(v)
}

func main() {
	fmt.Println(nonConcurrentSum([]int{2, 3, 4, 5, 6, 6}))
	fmt.Println(concurrentSum(3, []int{21321, 3342334, 23213, 1233}))
}
