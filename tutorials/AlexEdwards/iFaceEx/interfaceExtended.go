package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Customer struct {
	Name string
	Age  int
}

func (c *Customer) WriteJSON(w io.Writer) error {
	jsonified, err := json.Marshal(c)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonified)
	return err
}
func main() {
	c := &Customer{Name: "Alice", Age: 21}
	var buf bytes.Buffer
	err := c.WriteJSON(&buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf)
	f, err := os.Create("/tmp/customer")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = c.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}
}
