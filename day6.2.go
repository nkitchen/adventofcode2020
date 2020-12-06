package main

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "strings"

func main() {
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(inFile)
	if err != nil {
		log.Fatal(err)
	}

	inText := string(buf)

	sum := 0
	groups := strings.Split(inText, "\n\n")
	for _, g := range groups {
		yeses := map[rune]int{}
		people := strings.Fields(g)
		n := len(people)
		for _, p := range people {
			for _, c := range p {
				if 'a' <= c && c <= 'z' {
					yeses[c]++
				}
			}
		}

		for _, k := range yeses {
			if k == n {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
