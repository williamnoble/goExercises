package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	s := strings.NewReader("Some wildly redundant text, maybe I should be a politician")
	file, _ := os.Create("./file.dat")
	io.Copy(file, s)
}
