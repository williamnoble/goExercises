package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	MIN := 0
	MAX := 0
	TOTAL := 0
	if len(os.Args) > 3 {
		MIN, _ = strconv.Atoi(os.Args[1])
		MAX, _ = strconv.Atoi(os.Args[2])
		TOTAL, _ = strconv.Atoi(os.Args[3])
	} else {
		fmt.Println("Usage: ", os.Args[0], "MIN MAX TOTAL")
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < TOTAL; i++ {
		myrand := random(MIN, MAX)
		fmt.Print(myrand)
		fmt.Print("  -  ")
	}
	fmt.Println()
}

func random(min, max int) int {
	return rand.Intn((max - min) + min)
}

// (3-2) + 2 = 3
