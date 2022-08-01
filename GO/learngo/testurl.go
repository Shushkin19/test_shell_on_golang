package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		url = "http://" + url
		resp, err := http.Get(url)
		fmt.Println(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s\t %v\n",resp.Status, err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:чтение %s:  %v\n", url, err)
		}
		fmt.Printf("%s", b)
	}
}
