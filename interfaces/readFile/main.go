package main

import (
	"fmt"
	"os"
)

func main() {
	q := shiftArgs(0)
	fn := q()
	file, _ := os.Open(fn)

	// os.Stdout -> implements the Writer interface
	// os.Open -> implements the Read interface
	//	io.Copy(os.Stdout, file)

	// or
	stat, _ := file.Stat()
	bs := make([]byte, stat.Size())
	file.Read(bs)
	fmt.Println(string(bs))
}

func printError(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(0)
	}
}

func shiftArgs(s int) func() string {
	i := s
	return func() string {
		i++
		return os.Args[i]
	}
}
