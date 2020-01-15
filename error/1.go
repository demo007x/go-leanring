package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./time.go")
	fmt.Println(err)
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")

}
