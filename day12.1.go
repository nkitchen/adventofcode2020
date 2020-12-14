package main

import "bufio"
import "fmt"
import "log"
import "os"

func main() {
	x, y := 0, 0
	vx, vy := 1, 0

	turnLeft := func(n int) {
		for n > 0 {
			vx, vy = -vy, vx
			n -= 90
		}
	}

	for line := range inputLines() {
		var c rune
		var n int
		_, err := fmt.Sscanf(line, "%c%d", &c, &n)
		if err != nil {
			log.Fatal(err)
		}

		switch c {
		case 'N':
			y += n
		case 'S':
			y -= n
		case 'E':
			x += n
		case 'W':
			x -= n
		case 'F':
			x += vx * n
			y += vy * n
		case 'L':
			turnLeft(n)
		case 'R':
			turnLeft(360 - n)
		}
	}
	fmt.Println(abs(x) + abs(y))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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
