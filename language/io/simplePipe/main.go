package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Create("./pipedFile.dat")
	pr, pw := io.Pipe()
	go func() {
		fmt.Fprint(pw, "Pipe Streaming Baby!!")
		pw.Close()
	}()

	wait := make(chan struct{})
	go func() {
		io.Copy(file, pr)
		pr.Close()
		close(wait)
	}()

	<-wait
}
