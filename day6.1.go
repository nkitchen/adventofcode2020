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

	groupCounts := []int{}
	groups := strings.Split(inText, "\n\n")
	for _, g := range groups {
		answered := map[rune]bool{}
		for _, c := range g {
			if 'a' <= c && c <= 'z' {
				answered[c] = true
			}
		}
		groupCounts = append(groupCounts, len(answered))
	}

	s := 0
	for _, n := range groupCounts {
		s += n
	}
	fmt.Println(s)
}
