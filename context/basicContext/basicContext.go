package main

import (
	"context"
	"fmt"
)

type user struct {
	name string
}

type userKey int

func main() {
	u := user{
		name: "William",
	}

	const uk userKey = 0

	ctx := context.WithValue(context.Background(), uk, &u)

	if u, ok := ctx.Value(uk).(*user); ok {
		fmt.Println("User", u.name)
		fmt.Println(uk)
	}

	if _, ok := ctx.Value(0).(*user); !ok {
		fmt.Println("User not found")
	}
}
