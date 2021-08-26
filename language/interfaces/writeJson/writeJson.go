package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Customer changed json tags to emphasize tags.
type Customer struct {
	Name string `json:"customer_name"`
	Age  int    `json:"customer_age"`
}

// Implement a WriteJSON method that takes an io.Writer as the parameter.
// It marshals the customer struct to json, and if the marshal worked
// successfully, then calls the relevant io.Writer's Write() method.
func (c *Customer) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = w.Write(js)
	return err
}

func main() {
	// Initialize a customer struct.
	fmt.Println("start")
	c := &Customer{Name: "Alice", Age: 21}

	// We can then call the WriteJSON method using a buffer...
	var buf bytes.Buffer
	// Note we pass a bytes.Buffer because (b *Buffer) Write takes a
	// pointer receiver.
	err := c.WriteJSON(&buf)
	if err != nil {
		log.Fatal(err)
	}

	// Or using a file.
	f, err := os.Create("/tmp/customer")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	err = c.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}
}
