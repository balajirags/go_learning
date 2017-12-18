package main

import (
	"os"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"time"
)

const (
	HTTP = "http://"
)

func main() {

	start := time.Now()
	if (len(os.Args) < 2) {
		fmt.Println("Error")
	}

	queue := make(chan string, len(os.Args)-1)
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, HTTP) {
			url = HTTP + url
		}
		go fetch(url, queue)
	}
	for range os.Args[1:] {
		fmt.Println(<-queue) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(myUrl string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(myUrl)
	if (err != nil) {
		ch <- fmt.Sprint("error fetching url - ", err)
	}
	nbytes, e := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if (e != nil) {
		ch <- fmt.Sprint("error parsing response - ", e)
	}
	writeToFile(myUrl, nbytes)

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %s", secs, myUrl)
}
func writeToFile(myUrl string, nbytes []byte) {
	fileName := strings.Replace(myUrl, "//", "_", 100)
	fileName = strings.Replace(fileName, ".", "_", 100)
	fileName = strings.Replace(fileName, ":", "_", 100)
	writeError := ioutil.WriteFile(fileName, nbytes, 0644)
	if writeError != nil {
		panic(writeError)
	}
}
