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
	adapters = append(adapters, 3+adapters[len(adapters)-1])

	joltages := append([]int{0}, adapters...)
	d1j := 0
	d3j := 0
	for i := 0; i + 1 < len(joltages); i++ {
		d := joltages[i+1] - joltages[i]
		switch d {
		case 1:
			d1j++
		case 3:
			d3j++
		}
	}

	fmt.Println(d1j * d3j)
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
