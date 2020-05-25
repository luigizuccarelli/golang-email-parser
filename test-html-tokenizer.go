package main

import (
	"bytes"
	"fmt"
	//"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := ioutil.ReadFile("body.html")
	if err != nil {
		fmt.Printf("ERROR", err)
		os.Exit(-1)
	}
	r := bytes.NewReader(file)
	tokenizer := html.NewTokenizer(r)
	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				//end of the file, break out of the loop
				break
			}
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if "table" == token.Data {
				for a, _ := range token.Attr {
					if token.Attr[a].Key == "id" && strings.Contains(token.Attr[a].Val, "actions-to-take") {
						fmt.Printf("LMZ attr %v\n", token)

						tokenizer.Next()
						tokenizer.Next()
						tokenizer.Next()
						tokenizer.Next()
						tokenizer.Next()
						tokenizer.Next()
						tokenizer.Next()
						t := tokenizer.Token()
						fmt.Printf("LMZ next element %v\n", t)
					}
				}
			}
		}
	}
}

//if strings.Contains(buf.String(), "actions-to-take") {
//	html.Render(os.Stdout, buf.String())
//}
// Compile the expression once, usually at init time.
// Use raw strings to avoid having to quote the backslashes.
//var validID = regexp.MustCompile(`https:\/\/^[a-z]+\[[0-9.-]+\]*`)
//var validID = regexp.MustCompile(`style="display:none"`)
//s := validID.FindAllString(string(file), -1)
// remove the <loc> and </loc> tags then iterate through each url
//for index, _ := range s {
//url := s[index][5 : len(s[index])-6]
//id := strings.LastIndex(url, "/")
//name := url[id+1:]
//if name == "" {
//	name = "home.html"
//}
//	fmt.Println(s[index])
//}
