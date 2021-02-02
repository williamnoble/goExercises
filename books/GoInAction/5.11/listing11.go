package main

import "fmt"

type User struct {
	name  string
	email string
}

func (u *User) notify() {
	fmt.Printf("Sending email to: %s<%s>\n", u.email, u.name,
	)

}
func main() {
	bill := User{"Bill", "Bill@email.com"}
	tom := User{"Tom", "Tom@email.com"}

	bill.notify()
	tom.notify()
}
