package main

import (
	"fmt"
	"os"

	"github.com/AlperAkca79/cof/internal/cof"
)

func main() {
	// Print usage of cat if length of argument is not equal to 2
	if len(os.Args) < 2 {
		fmt.Println(`View the contents of your file in your terminal with 'cof'.

USAGE
cof <filename>`)
	} else {
		for i := 1; i < len(os.Args); i++ {
			cof.ViewContent(os.Args[i])
			fmt.Print("\n")
		}
	}
}
