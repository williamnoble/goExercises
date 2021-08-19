package main

import "fmt"

type TeamMember interface {
	Language() string
}

type Developer struct {
	devName string
	lang    string
}

func (d Developer) Language() string {
	return fmt.Sprintf("SuperDev %s devs in %q", d.devName, d.lang)
}

func main() {
	devA := Developer{"William", "go"}
	devB := Developer{"Joseph", "cobalt"}

	SuperDevs := [2]TeamMember{devA, devB}
	// can just write i.. no need for i , _ :=
	for i := range SuperDevs {
		s := SuperDevs[i].Language()
		fmt.Println(s)
	}

}
