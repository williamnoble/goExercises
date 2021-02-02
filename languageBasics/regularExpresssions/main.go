package main

import (
	"fmt"
	"regexp"
)

func main() {
	// "AAB" in "DFAAGBJ" :: true
	match, _ := regexp.MatchString("AAB", "DFAABGJ")
	fmt.Println(match)

	// "a.g" A<any character>G in "argkz" :: true
	between, _ := regexp.MatchString("a.g", "argkz")
	fmt.Println(between)

	// match a full string
	fullString, _ := regexp.MatchString("^As{2}Z$", "AssZ")
	fmt.Println(fullString)
}
