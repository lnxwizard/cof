package main

import (
	"fmt"
	"os"

	"github.com/AlperAkca79/cof/internal/cof"
)

func main() {
	// Print usage of cat if length of argument is not equal to 2
	if len(os.Args) < 2 {
		fmt.Println("Usage: cof <filename>")
	} else {
		for i := 1; i < len(os.Args); i++ {
			cof.COF(os.Args[i])
			fmt.Print("\n")
		}
	}
}
