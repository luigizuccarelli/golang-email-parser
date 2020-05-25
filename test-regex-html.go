package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	//"strings"
)

func main() {

	file, err := ioutil.ReadFile("body.html")
	if err != nil {
		fmt.Printf("ERROR", err)
		os.Exit(-1)
	}

	// use regex to extract the table
	// regex breakdown
	// table with attributes
	// add whitespaces and check fro one instance of <tbody>
	// add whitespeces
	// zero or one instance of <tr> and whitespaces
	// one or more instances of <td>data..</td> with whitespaces
	// zero or one instance of </tr> adn whitespaces
	// one instance of </tbody> and whitespaces
	// one instance of </table> and whitespaces
	var tableData = regexp.MustCompile(`<table [0-9a-zA-Z=\_\-\"\s\:\;\,]*>[\s]*(<tbody>){1}[\s]*([\s]*(<tr>)?[\s]*(<td>[\w\d\-]*</td>[\s]*)+[\s]*(</tr>)?[\s]*[\s]*)*(</tbody>){1}[\s]*(</table>){1}`)
	s := tableData.FindAllString(string(file), -1)
	fmt.Println(s[0])
	//r := strings.NewReplacer("{", "", "}", "", "subs", "", "\"", "")
	//result := r.Replace(string(codes))
	//data := strings.Split(s[0], "</tr>")
	//for x, _ := range data {
	//	i := strings.Index(data[x], "<td>")
	//}
	//fmt.Println(data[1])
}
