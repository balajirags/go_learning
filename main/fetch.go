package main


import(
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)

func main(){

	if len(os.Args) < 2{
		fmt.Println(" Please enter a url")
		os.Exit(1)
	}

	url := os.Args[1]
	resp, err := http.Get(url)
	if(err!=nil){
		fmt.Fprintf(os.Stderr, "Fetching url - %s - got error - %v", url, err)
		os.Exit(1)
	}
	bytes, e := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if(e!=nil){
		fmt.Fprintf(os.Stderr, "Parsing response and got error - %v", e)
		os.Exit(1)
	}
	fmt.Printf("%v", string(bytes))
}
