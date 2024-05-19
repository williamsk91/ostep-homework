package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file := flag.String("f", "seek", "file to read")

	nFlag := flag.Int("n", 10, "seek distance from end of file")
	n := int64(*nFlag)

	flag.Parse()

	f, err := os.Open(*file)
	defer f.Close()
	check(err)

	o3, err := f.Seek(-n, io.SeekEnd)
	check(err)

	b3 := make([]byte, n)
	n3, err := io.ReadAtLeast(f, b3, int(n))
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

}
