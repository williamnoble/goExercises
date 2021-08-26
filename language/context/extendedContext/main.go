package main

import (
	"context"
	"fmt"
)

type contextKey string

func (c contextKey) String() string {
	return "my pakcage context key " + string(c)
}

var (
	contextKeyAuthToken = contextKey("auth-token")
	contextKeyAnother   = contextKey("another")
)

func main() {
	token := "123"
	ctx := context.WithValue(context.Background(), contextKeyAuthToken, token)
	s, ok := GetAuthToken(ctx)
	fmt.Println("1st part: ", s, ok)
	user, ok := LookupUserFromToken(ctx, s)
	fmt.Println(user)
}

func GetAuthToken(ctx context.Context) (string, bool) {
	tokenStr, ok := ctx.Value(contextKeyAuthToken).(string)
	return tokenStr, ok
}

type User struct {
	name string
}

func LookupUserFromToken(ctx context.Context, token string) (User, bool) {

	users := make(map[string]User)
	users["123"] = User{name: "William"}
	//fmt.Println("User: ", users["123"])
	u, ok := users[token]
	return u, ok
}
