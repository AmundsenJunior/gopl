// fetch prints the content returned from a URL request
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		check(err)

		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		check(err)

		fmt.Printf("%s", b)
	}
}

func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", e)
		os.Exit(1)
	}
}
