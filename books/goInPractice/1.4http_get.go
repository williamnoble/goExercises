package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://www.google.com")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	_ = resp.Body.Close()
}
