package main

import "bufio"
import "fmt"
import "log"
import "os"
import "regexp"
import "strconv"

func main() {
	heldBy := map[string]map[string]int{}
	holderRe := regexp.MustCompile(`(\S+ \S+) bags contain`)
	heldRe := regexp.MustCompile(`(\d+) (\S+ \S+) bags?`)
	for line := range inputLines() {
		mHolder := holderRe.FindStringSubmatch(line)
		if len(mHolder) == 0 {
			panic("Unexpected format: " + line)
		}
		holder := mHolder[1]

		for _, mHeld := range heldRe.FindAllStringSubmatch(line, -1) {
			held := mHeld[2]
			k, _ := strconv.Atoi(mHeld[1])
			if heldBy[held] == nil {
				heldBy[held] = map[string]int{}
			}
			heldBy[held][holder] = k
		}
	}

	myBag := "shiny gold"
	checked := map[string]bool{}
	q := []string{myBag}
	for len(q) > 0 {
		bag := q[0]
		q = q[1:]
		checked[bag] = true
		for holder := range heldBy[bag] {
			if !checked[holder] {
				q = append(q, holder)
			}
		}
	}
	delete(checked, myBag)
	fmt.Println(len(checked))
}

func inputLines() <-chan string {
	ch := make(chan string)

	go func() {
		inFile, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(inFile)
		for scanner.Scan() {
			line := scanner.Text()
			ch <- line
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		close(ch)
	}()

	return ch
}
