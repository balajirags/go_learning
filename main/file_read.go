package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	count := make(map[string]int)
	for _, inputFile := range os.Args[1:] {
		file, _ := os.Open(inputFile)
		inputScanner := bufio.NewScanner(file)
		for inputScanner.Scan() {
				count[inputFile]++
		}
		fmt.Printf("File name - %s has %d lines", inputFile, count[inputFile])
	}
}
