package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("File Open Error: %q\n", err)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	s := ""
	for scanner.Scan() {
		if len(s) == 0 {
			s = scanner.Text()
		} else {
			s = s + "," + scanner.Text()
		}
	}
	fmt.Printf("%s\n", s)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %q\n", err)
	}
}
