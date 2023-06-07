package cof

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/cyucelen/marker"
	"github.com/fatih/color"
)

func ViewContent(filename string) {
	// Open the file specified by the argument
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// filename
	completeMessage := file.Name() + ":\n"

	// marking filename
	emphasizedFilename := marker.Mark(completeMessage, marker.MatchAll(file.Name()), color.New(color.FgGreen))

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Printing colorful filename
	fmt.Print(emphasizedFilename)
	fmt.Println("-------------------------------------------------------------")

	for scanner.Scan() {
		fmt.Print("# | ", scanner.Text(), "\n")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
