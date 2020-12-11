package main

import "bufio"
import "fmt"
import "log"
import "os"
import "sort"
import "strconv"

func main() {
	adapters := []int{}
	for line := range inputLines() {
		a, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		adapters = append(adapters, a)
	}

	sort.IntSlice(adapters).Sort()
	n := len(adapters)
	joltageMax := 3 + adapters[n-1]
	adapters = append(adapters, joltageMax)

	// ways[j] is the number of valid ways to get joltage j
	ways := make([]int64, joltageMax+1)

	ways[0] = 1
	for _, a := range adapters {
		w := int64(0)
		for j := a - 3; j <= a - 1; j++ {
			if j < 0 {
				continue
			}
			w += ways[j]
		}
		ways[a] = w
	}
	fmt.Println(ways[joltageMax])
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
