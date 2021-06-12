package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Uncommented because we're using 'unstructured result' via map
// type Users struct {
// 	Users []User `json:"users"`
// }

// User struct maps valsues from our Users.json to fiels within the backend
// type User struct {
// 	Name   string `json:"name"`
// 	Type   string `json:"type"`
// 	Age    int    `json:"age"`
// 	Social Social `json:"social"`
// }

// Social maps values from users.json to the User struct within Backend
// type Social struct {
// 	Facebook string `json:"facebook"`
// 	Twitter  string `json:"twitter"`
// }

func main() {
	//Open our JSON file
	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened our JSON File")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	//var users Users
	//json.Unmarshal(byteValue, &users)
	var result map[string]interface{}
	if err := json.Unmarshal(byteValue, &result); err != nil {
		fmt.Println("Error")
	}

	fmt.Println(result["users"])
	fmt.Println(result)
}
