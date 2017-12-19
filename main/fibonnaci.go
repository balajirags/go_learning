package main

import (
	"os"
	"strconv"
	"fmt"
)

func main() {
	input, err := strconv.ParseInt(os.Args[1], 10, 64)
	if (err != nil) {
		os.Exit(1)
	}
	var x int64 = 0
	var y int64 = 1
	for i := int64(0); i < input; i++ {
		x, y = y,x+y
		fmt.Println(y-x)
	}

}
