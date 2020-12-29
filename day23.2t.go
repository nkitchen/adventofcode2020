package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"

func main() {
	ch := inputLines()
	line := <-ch

	million := 1_000_000
	circle := make([]int, 0, million * 10)
	for _, c := range line {
		circle = append(circle, int(c - '0'))
	}
	for i := 10; i <= million; i++ {
		circle = append(circle, i)
	}

	turns := 10_000_000
	if len(os.Args) > 2 {
		turns, _ = strconv.Atoi(os.Args[2])
	}

	for turn := 0; turn < turns; turn++ {
		pickedUp := circle[:3]
		circle = append(append(circle[3:], pickedUp[1:3]...), pickedUp[0])
	}
	fmt.Println(circle[0])
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
