package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("length of 'Alphabet'：", len("Alphabet"))
	fmt.Println(fmt.Sprintf("Nicely Formmated String: %v %s %q", "Hello", "There", "Mr Noname"))
	fmt.Println("Onomatopoeia contains the substring mat:", strings.Contains("Onomatopoeia", "mat")) // => bool
	fmt.Println(strings.HasPrefix("PlanetEarth", "planet"))                                          // => False (case sensitive)
	fmt.Println(strings.HasSuffix("PlanetEarth", "Earth"))
	fmt.Println("Indexed substring:", strings.Index("Alphabet", "ha")) // => True (case sensitive)
	s := "A.L.P.H.A.B.E.T"
	splitS := strings.Split(s, ".")
	fmt.Println("string sliced by '.' yielding:", splitS)
	for index, value := range splitS {
		if index == len(splitS)-1 {
			fmt.Printf("Split String: %d: %s\n", index, value)
		} else {
			fmt.Printf("Split String: %d: %s\n", index, value)
		}
	}
}
