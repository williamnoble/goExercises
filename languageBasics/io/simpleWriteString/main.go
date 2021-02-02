package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fout, err := os.Create("./file.dat")
	if err != nil {
		fmt.Println(err)
	}

	defer fout.Close()
	io.WriteString(fout, "If only one could easily write a string to a file, alas it seems too complicated\n")

}
