package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	host := os.Args[1]

	resp, err := http.Get("http://" + host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		defer resp.Body.Close()
		r := resp.Body
		bytes, err := ioutil.ReadAll(r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("HTML:\n\n", string(bytes))

		z := html.NewTokenizer(r)

		for {
			tt := z.Next()

			switch {
			case tt == html.ErrorToken:
				return
			case tt == html.StartTagToken:
				t := z.Token()

				isAnchor := t.Data == "a"
				if isAnchor {
					fmt.Println("We found a link!")
				}
			}
		}
	}
}
