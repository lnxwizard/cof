package cof

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func COF(filename string) {
	// Open the file specified by the argument
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// complete message
	completeMessage := file.Name() + ":\n"

	// marking filename
	emphasized := marker.Mark(completeMessage, marker.MatchAll(file.Name()), color.New(color.FgGreen))

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Printing colorful filename
	fmt.Println(emphasized)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
