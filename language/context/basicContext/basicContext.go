package main

import (
	"context"
	"fmt"
)

type user struct {
	name string
}

type userKey string

func main() {
	u := user{
		name: "William",
	}

	var k userKey
	ctx := context.WithValue(context.Background(), k, u)

	value := ctx.Value(k).(user)
	fmt.Println(value.name)

	// Or we Pass it a pointer to &u then dereference .(*user)
}
