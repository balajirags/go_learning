package main

import (
	"net/http"
	"os"
	"fmt"
	"io"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(" Please enter a url")
		os.Exit(1)
	}

	const prefix = "http://"
	var url string
	if (!strings.HasPrefix(os.Args[1], prefix)) {
		url = prefix + os.Args[1]
	}else {
		url = os.Args[1]
	}

	resp, err := http.Get(url)
	responseCode := resp.StatusCode
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "Fetching url - %s - got error - %v", url, err)
		os.Exit(1)
	}
	buff := make([]byte, 10)
	_, e := io.CopyBuffer(os.Stdout, resp.Body, buff)
	if (e != nil) {
		fmt.Fprintf(os.Stderr, "Parsing response and got error - %v", e)
		os.Exit(1)
	}
	fmt.Println(responseCode)
}
