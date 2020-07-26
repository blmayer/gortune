// +build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const url = "http://www.bsdfortune.com/fortunes/"

var fortunes = map[string][]string{}
var categories = []string{}

func die(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	// Read categories file
	categoryFile, err := os.Open("gen/categories.txt")
	die(err)
	defer categoryFile.Close()

	scanner := bufio.NewScanner(categoryFile)
	for scanner.Scan() {
		categories = append(categories, scanner.Text())
	}
	die(scanner.Err())

	// Download and parse categories
	for _, category := range categories {
		res, err := http.Get(url + category)
		if err != nil {
			log.Println("request", err.Error())
			continue
		}
		content := make([]byte, res.ContentLength)
		n, err := res.Body.Read(content)
		for int64(n) < res.ContentLength {
			var k int
			if n == 0 && err != nil {
				log.Println(err)
				break
			}
			k, err = res.Body.Read(content[n:])
			n += k
		}
		res.Body.Close()
		
		fortuneList := strings.Split(string(content), "\n%")
		for i, f := range fortuneList {
			fortuneList[i] = strings.TrimSpace(f)
		}
		fortunes[category] = fortuneList
	}

	// Write fortunes
	fortunesFile, err := os.Create("fortunes/fortunes.go")
	die(err)
	defer fortunesFile.Close()

	fortunesFile.WriteString("package fortunes\n")
	fortunesFile.WriteString(fmt.Sprintf("var Categories = %#v\n", categories))
	fortunesFile.WriteString(fmt.Sprintf("var List = %#v\n", fortunes))
}
