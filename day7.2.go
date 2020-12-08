package main

import "bufio"
import "fmt"
import "log"
import "os"
import "regexp"
import "strconv"

func main() {
	holds := map[string]map[string]int{}
	holderRe := regexp.MustCompile(`(\S+ \S+) bags contain`)
	heldRe := regexp.MustCompile(`(\d+) (\S+ \S+) bags?`)
	for line := range inputLines() {
		mHolder := holderRe.FindStringSubmatch(line)
		if len(mHolder) == 0 {
			panic("Unexpected format: " + line)
		}
		holder := mHolder[1]

		holds[holder] = map[string]int{}

		for _, mHeld := range heldRe.FindAllStringSubmatch(line, -1) {
			held := mHeld[2]
			k, _ := strconv.Atoi(mHeld[1])
			holds[holder][held] = k
		}
	}

	totalContainedMemo := map[string]int{}
	var totalContained func( string) int
	totalContained = func(holder string) int {
		if n, ok := totalContainedMemo[holder]; ok {
			return n
		}

		n := 0
		for held, k := range holds[holder] {
			r := totalContained(held)
			n += k * (1 + r)
		}
		totalContainedMemo[holder] = n
		return n
	}
	myBag := "shiny gold"
	fmt.Println(totalContained(myBag))
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
