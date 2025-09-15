package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Print("Too many arguments")
		return
	}
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Println(err.Error())
	}
}
