package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	pExecutable, _ := os.Executable()
	pExecutableDir := filepath.Dir(pExecutable)
	_ = pExecutableDir
	f, _ := os.Create("file.txt")
	defer f.Close()
	fmt.Println("Got this far")

}
