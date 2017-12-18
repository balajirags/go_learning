package main

import (
	"bufio"
	"strings"
	"fmt"
)

func main(){

	input := "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"

	inputScanner := bufio.NewScanner(strings.NewReader(input))

	inputScanner.Split(bufio.ScanWords)

	count := 0

	for inputScanner.Scan(){
		count++
	}

	fmt.Println("count", count)
}
