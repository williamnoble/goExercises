package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute() {
	out, err := exec.Command("ls", "-ltr").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	// output := string(out[:])
	output := string(out)
	fmt.Println(output)
}

func execute2() {
	out2, err := exec.Command("pwd").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Successfully excecuted the pwd command")
	output2 := string(out2[:])
	fmt.Println(output2)
}

func main() {
	if runtime.GOOS == "Windows" {
		fmt.Println("Can't excute on a Windows Machine")
	} else {
		execute()
		execute2()
	}
}
