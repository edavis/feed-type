package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func feedType(elem xml.StartElement) string {
	switch elem.Name.Local {
	case "rss":
		return "rss"
	case "feed":
		return "atom"
	default:
		return "unknown"
	}
}

func main() {
	url := flag.String("url", "", "URL to check")
	flag.Parse()

	if *url == "" {
		fmt.Fprintf(os.Stderr, "no URL given\n")
		return
	}
	resp, err := http.Get(*url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "http error: %v\n", err)
	}

	decoder := xml.NewDecoder(resp.Body)
	for {
		token, err := decoder.Token()
		if token == nil {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "xml error: %v\n", err)
		}
		if elem, ok := token.(xml.StartElement); ok {
			fmt.Println(feedType(elem))
			break
		}
	}
}
