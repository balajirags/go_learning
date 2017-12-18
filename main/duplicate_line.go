package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	lineCount := make(map[string]int)
	for (input.Scan()) {
		if(input.Text() == "stop"){
			break
		}
		lineCount[input.Text()]++
	}

	for key, value := range lineCount{
		if(value > 1){
			fmt.Printf("%s : %d\n", key, value)
		}
	}
}
