package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"

func main() {
	data := []int{}
	for line := range inputLines() {
		x, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, x)
	}

	iv := findInvalid(data)

	type Range struct {
		first, last int
	}
	rangeSum := map[Range]int{}
	for i, x := range data {
		s := x
		for j := i + 1; j < len(data); j++ {
			s += data[j]
			rangeSum[Range{i, j}] = s
		}
	}

	for r, s := range rangeSum {
		if s == iv {
			min := data[r.first]
			max := data[r.last]
			for i := r.first + 1; i <= r.last; i++ {
				if data[i] < min {
					min = data[i]
				}
				if data[i] > max {
					max = data[i]
				}
			}
			fmt.Println(min + max)
			return
		}
	}
}

func findInvalid(data []int) int {
	p, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	wq := []int{}
	ws := map[int]bool{}
	for _, z := range data {
		if len(ws) < p {
			wq = append(wq, z)
			ws[z] = true
			continue
		}

		j := -1
		for i, x := range wq {
			y := z - x
			if ws[y] {
				j = i
				break
			}
		}
		if j == -1 {
			return z
		}
		delete(ws, wq[0])
		wq = append(wq[1:], z)
		ws[z] = true
	}

	panic("Invalid data not found")
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
