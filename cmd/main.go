package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Print usage of cat if length of argument is not equal to 2
	if len(os.Args) != 2 {
		fmt.Println("Usage: cat fileName")
	} else {
		// Open the file specified by the argument
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Use a scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		fmt.Print(file.Name(), ":\n\n")
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
