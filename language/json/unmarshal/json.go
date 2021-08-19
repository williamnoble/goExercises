package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name            string           `json:"name,omitempty"`
	Language        string           `json:"language,omitempty"`
	Interests       []string         `json:"interests,omitempty"`
	Characteristics []Characteristic `json:"characteristics,omitempty"`
}

type Characteristic struct {
	Height    string `json:"height,omitempty"`
	EyeColour string `json:"eye_colour,omitempty"`
}

func main() {
	//f, _ := os.Open("./data.json")
	//defer f.Close()

	b, _ := ioutil.ReadFile("./data.json")
	var users Users
	//goland:noinspection GoUnhandledErrorResult
	json.Unmarshal(b, &users)
	fmt.Println(string(b))
}
