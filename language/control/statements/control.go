package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("**COMPARISONS**\n\n")
	Comparisons()

	fmt.Printf("**Interface**\n\n")
	Interfaces("Hi")
	Interfaces(false)
	Interfaces(293)
	Interfaces(21.02)
	Interfaces(struct{}{})
	for i := 1; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(3)
		Switches(r)
	}

	fmt.Printf("**FALLTHROUGH**\n\n")
	Fallthrough()

	fmt.Printf("**MULTIPLE SWITCHES**\n\n")
	MultipleSwitches()

	fmt.Printf("**GOTO**\n\n")
	Goto()
	TimesTable()
}

func Comparisons() {

	a, b, c := 10, 20, 30
	if a == 10 {
		fmt.Println("'a' is = 10")
		if b == 20 && c == 30 {
			fmt.Println("All integers correct")
		}
		if a+b == c {
			fmt.Println("'a' + 'b' = 'c'")
		}
	} else {
		fmt.Println("Have you altered the value of 'a' yet")
	}
}

func Interfaces(ex interface{}) {

	var x interface{}
	x = ex
	switch i := x.(type) {
	case nil:
		fmt.Printf("x: %T\n", i)
	case int:
		fmt.Printf("x: int\n")
	case float64:
		fmt.Printf("x: float64\n")
	case bool:
		fmt.Printf("x: bool\n")
	case string:
		fmt.Printf("x: string\n")
	case func(int) float64:
		fmt.Printf("x func(int)\n")
	default:
		fmt.Printf("No Value Given\n")
	}
}

func Switches(x int) {
	switch x {
	case 0:
		fmt.Println("ZERO")
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	default:
		fmt.Println("No match")
	}
}

func Fallthrough() {
	x := 5
	switch x {
	case 0:
		fmt.Println("f.zero")
	case 1:
		fmt.Println("f.one")
		fallthrough
	case 2, 4, 5:
		fmt.Println("f.two")
	default:
		fmt.Println("f.Default case")
	}
}

func MultipleSwitches() {

	x := 1
	switch x {
	case 0, 1:
		fmt.Println("Selected either 0 or 1")
	case 2:
		fmt.Println("Only 2 remains")
	default:
		fmt.Println("default case")
	}
}

// Note: switching without X because we are comparing bools not ints.
func Comparison() {
	x := 13
	if x < 1 {
		fmt.Println("please supply a non-negative integer")
	}
	// begin
	switch {
	case (x >= 1 && x <= 10):
		fmt.Println("Between ONE And TEN")
	case (x >= 11 && x <= 20):
		fmt.Println("Between TEN and TWENTY")
	default:
		fmt.Println("default case")
	}
}

func Goto() {

	var found rune

	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'F' {
				found = 'F'
				goto over
			}
			fmt.Printf("%v ", string(j))
		}
	}
over:
	fmt.Printf("found %s\n\n", string(found))
}

func TimesTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
	fmt.Println()
}
