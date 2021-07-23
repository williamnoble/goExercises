package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "William")

	got := buffer.String()
	want := "Hello, William"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}