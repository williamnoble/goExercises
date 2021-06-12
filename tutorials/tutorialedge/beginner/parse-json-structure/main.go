package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

// User struct maps valsues from our Users.json to fiels within the backend
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

// Social maps values from users.json to the User struct within Backend
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened our JSON File")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook URI: " + users.Users[i].Social.Facebook)
	}
}
