package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"
import "strings"

func main() {
	ch := inputLines()

	line := <-ch

	starting := []int{}
	for _, s := range strings.Split(line, ",") {
		x, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		starting = append(starting, x)
	}

	spoken := map[int][2]int{}
	last := 0
	for turn := 1; turn <= 2020; turn++ {
		var x int
		if turn <= len(starting) {
			x = starting[turn - 1]
		} else {
			hist := spoken[last]
			if hist[1] == 0 {
				x = 0
			} else {
				x = hist[0] - hist[1]
			}
		}
		 

		p := spoken[x][0]
		spoken[x] = [2]int{turn, p}
		last = x
	}
	fmt.Println(last)
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
