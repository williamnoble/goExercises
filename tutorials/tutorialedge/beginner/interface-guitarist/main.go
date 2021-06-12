package main

import "fmt"

type Guitarist interface {
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Base Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

func main() {
	var player BaseGuitarist
	player.Name = "Paul"
	player.PlayGuitar()

	var player2 AcousticGuitarist
	player2.Name = "Star"
	player2.PlayGuitar()

	var guitarists []Guitarist
	guitarists = append(guitarists, player)
	guitarists = append(guitarists, player2)
}
