package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	//"gopkg.in/jdkato/prose.v2"
)

type ActionItem struct {
	Action string `json:"action"`
	Stock  string `json:"stock"`
	Date   string `json:"date"`
	Price  string `json:"price"`
	Time   string `json:"time"`
}

func buildActionItem(in string) ([]ActionItem, error) {

	action := `buy-to-open|buy-to-close|sell-to-close|sell-to-open|limit order to buy|buy|sell`
	stock := `\([a-z:0-9:\s]*\)`
	date := `(january|february|march|april|may|june|july|august|september|october|november|december)(\s[0-9,\s]*)`
	price := `[0-9]*%\sgains or at market|at-the-money|up to \$[0-9\.]*\s|pay no more than \$[0-9\.]*|\$[0-9.]* or less|up to \$[0-9\.]*|\$[0-9\.]* or better|[0-9]{2}:[0-9]{2}`
	time := `\s[0-9]{2}:[0-9]{2}|([0-9]*pm|[0-9]*am)(\son\s)(monday|tuesday|wednesday|thursday|friday)(,\s)(january|february|march|april|may|june|july|august|september|october|november|december)(\s[0-9,\s]*)`
	items := make([]ActionItem, 2)

	// fmt.Println(fmt.Sprintf("DEBUG %s %v ", in, s))
	cmd := regexp.MustCompile(action)
	s := cmd.FindAllString(in, -1)
	if len(s) >= 2 {
		items[0].Action = s[0]
		items[1].Action = s[1]
	} else if len(s) == 1 {
		items[0].Action = s[0]
	}

	cmd = regexp.MustCompile(stock)
	s = cmd.FindAllString(in, -1)
	if len(s) >= 2 {
		items[0].Stock = s[0]
		items[1].Stock = s[1]
	} else if len(s) == 1 {
		items[0].Stock = s[0]
	}

	cmd = regexp.MustCompile(date)
	s = cmd.FindAllString(in, -1)
	if len(s) >= 2 {
		items[0].Date = s[0]
		items[1].Date = s[1]
	} else if len(s) == 1 {
		items[0].Date = s[0]
	}

	cmd = regexp.MustCompile(price)
	s = cmd.FindAllString(in, -1)
	if len(s) >= 1 {
		items[0].Price = s[0]
		items[1].Price = s[0]
	}

	cmd = regexp.MustCompile(time)
	s = cmd.FindAllString(in, -1)
	if len(s) >= 1 {
		items[0].Time = s[0]
	}
	return items, nil
}

func main() {
	file, err := ioutil.ReadFile("email6.txt")
	if err != nil {
		fmt.Printf("ERROR", err)
		os.Exit(-1)
	}

	var items, tmp []ActionItem
	hld := strings.Split(string(file), "\n")
	for x, line := range hld {
		text := strings.ToLower(line)

		tmp, _ = buildActionItem(text)
		for y, _ := range tmp {
			if tmp[y].Action != "" && tmp[y].Stock != "" {
				if tmp[y].Price == "" {
					// try add the price - should be next line
					cmd := regexp.MustCompile(`at-the-money|up to \$[0-9\.]*\s|pay no more than \$[0-9\.]*|\$[0-9.]* or less|up to \$[0-9\.]*|\$[0-9\.]* or better|[0-9]{2}:[0-9]{2}`)
					s := cmd.FindAllString(strings.ToLower(hld[x+1]), -1)
					if len(s) >= 1 {
						tmp[y].Price = s[0]
					}
				}
				items = append(items, tmp[y])
			}
		}
	}
	for y, _ := range items {
		fmt.Println(fmt.Sprintf("ACTION: %s", items[y].Action))
		fmt.Println(fmt.Sprintf("STOCK: %s", items[y].Stock))
		fmt.Println(fmt.Sprintf("DATE: %s", items[y].Date))
		fmt.Println(fmt.Sprintf("PRICE: %s", items[y].Price))
		fmt.Println(fmt.Sprintf("TIME: %s", items[y].Time))
		fmt.Println("")
	}
}
