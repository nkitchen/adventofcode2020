package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"

func main() {
	p, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	wq := []int{}
	ws := map[int]bool{}
	for line := range inputLines() {
		z, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

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
			fmt.Println(z)
			return
		}
		delete(ws, wq[0])
		wq = append(wq[1:], z)
		ws[z] = true
	}
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
