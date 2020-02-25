package main

import (
	"fmt"
	"github.com/golang/go/src/pkg/io/ioutil"
	"net/http"
	"os"
)

func main(){
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "error accessing url %s :: %v",url, err)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			fmt.Fprintf(os.Stderr, "unable to read url content for %s :: %v", url, err)
			continue
		}
		fmt.Printf("%s", string(body))
	}
}
