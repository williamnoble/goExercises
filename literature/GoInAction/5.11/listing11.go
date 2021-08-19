package main

// edited version.

import "fmt"

type User struct {
	name  string
	email string
}

func (u *User) notify() {
	fmt.Printf("Sending an email "+
		"to: %s<%s>  ", u.name, u.email)
}

func main() {
	bill := User{
		name:  "Bill",
		email: "User@mail.com",
	}

	tom := User{
		name:  "Timothy",
		email: "Timbothegreat@mail.com",
	}

	bill.notify()
	tom.notify()

	users := []User{bill, tom}
	fmt.Println(users)

	for _, v := range users {
		fmt.Printf("%s - %s", v.email, v.name)
	}
}
