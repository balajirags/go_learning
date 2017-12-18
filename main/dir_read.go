package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	directory := os.Args[1]
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Error - ", err)
	}

	for _, f := range files{
		fmt.Printf("%s - %d \n", f.Name(), f.Size())
	}
}
