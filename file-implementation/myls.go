package main

import (
	"fmt"
	"os"
)

func main() {

	pathToCheck := "." // default to current directory
	if len(os.Args) >= 2 {
		pathToCheck = os.Args[1]
	}

	fmt.Println("pathToCheck: ", pathToCheck)

	files, err := os.ReadDir(pathToCheck)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}
