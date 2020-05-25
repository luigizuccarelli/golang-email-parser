package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"gopkg.in/jdkato/prose.v2"
)

type ActionItem struct {
	Action string `json:"action"`
	Stock  string `json:"stock"`
	Date   string `json:"date"`
	Price  string `json:"price"`
	Time   string `json:"time"`
}

func main() {
	file, err := ioutil.ReadFile("email2.txt")
	if err != nil {
		fmt.Printf("ERROR", err)
		os.Exit(-1)
	}

	// Create a new document with the default configuration:
	doc, err := prose.NewDocument(string(file))
	if err != nil {
		log.Fatal(err)
	}
	/*
		fmt.Println("TOKENS -----------")
		// Iterate over the doc's tokens:
		for _, tok := range doc.Tokens() {
			fmt.Println(tok.Text, tok.Tag, tok.Label)
		}
		fmt.Println("END TOKENS -----------\n")

		fmt.Println("ENTITIES -----------")
		// Iterate over the doc's named-entities:
		for _, ent := range doc.Entities() {
			fmt.Println(ent.Text, ent.Label)
		}
		fmt.Println("END ENTITIES -----------\n")
	*/

	template := `buy-to-open|buy-to-close|sell-to-close|sell-to-open|limit|limit order|buy|sell|\([a-z:0-9:\s]*\)|(january|february|march|april|may|june|july|august|september|october|november|december)|(\sat-the-money\s|\sup to \$[0-9\.]*\s|pay no more than \$[0-9\.]*|\$[0-9.]* or less|up to \$[0-9\.]*|\$[0-9\.]* or better|\s[0-9]{2}:[0-9]{2})`

	var items []ActionItem
	var item ActionItem

	// Iterate over the doc's sentences:
	for _, sent := range doc.Sentences() {
		text := strings.ToLower(sent.Text)

		var cmd = regexp.MustCompile(template)
		s := cmd.FindAllString(text, -1)
		if len(s) > 2 && strings.Trim(s[1], " ") != "" && strings.Trim(s[2], " ") != "" {
			for x, _ := range s {
				switch x {
				case 0:
					item = ActionItem{Action: s[x]}
					break
				case 1:
					item.Stock = s[x]
					break
				case 2:
					item.Date = s[x]
					break
				case 3:
					item.Price = s[x]
					break
				case 4:
					item.Time = s[x]
					break
				}
			}
			items = append(items, item)
			item = ActionItem{}
		}
	}
	for y, _ := range items {
		fmt.Println(fmt.Sprintf("ACTION: %s", items[y].Action))
		fmt.Println(fmt.Sprintf("STOCK: %s", items[y].Stock))
		fmt.Println(fmt.Sprintf("DATE: %s", items[y].Date))
		fmt.Println(fmt.Sprintf("PRICE: %s", items[y].Price))
		fmt.Println(fmt.Sprintf("DATE: %s", items[y].Time))
		fmt.Println("")
	}
}
