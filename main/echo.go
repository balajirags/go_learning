package main

import (
	"os"
	"fmt"
	"strings"
	"time"
)

func main() {


	fmt.Println("Command Name", os.Args[0])

	fv := firstVersion()
	fmt.Println("First Version", fv)

	sv := secondVersion()
	fmt.Println("Second Version", sv)

	tv := thirdVersion()
	fmt.Println("Third Version", tv)

}
func firstVersion() string {
	var s string
	startTime := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += " " + os.Args[i]
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("First Version time taken ", elapsedTime)
	return s
}

func secondVersion() string {
	startTime := time.Now()
	var s string
	for _, value := range os.Args[1:] {
		s += " " + value
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Second Version time taken ", elapsedTime)
	return s
}

func thirdVersion() string {
	startTime := time.Now()
	var s string
	s = strings.Join(os.Args[1:], " ")
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Third Version time taken ", elapsedTime)
	return s
}
