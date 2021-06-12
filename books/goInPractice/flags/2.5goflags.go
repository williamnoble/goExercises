package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short:"n" long:"name" default:"world" description:"A name to say hello to"`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish Language"`
}

func main() {
	flags.Parse(&opts)
	if opts.Spanish == true {
		fmt.Printf("Hola %s\n", opts.Name)

	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}

/*
go run 2.5goflags.go --> Hello world!
go run 2.5goflags.go -s false
*/
