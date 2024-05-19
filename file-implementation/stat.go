package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	pathToCheck := os.Args[1]
	fmt.Println(pathToCheck)

	var stat syscall.Stat_t
	syscall.Stat(pathToCheck, &stat)
	fmt.Println(stat)
	fmt.Println(stat.Ino)
	fmt.Println(stat.Blocks)
	fmt.Println(stat.Nlink)
}
