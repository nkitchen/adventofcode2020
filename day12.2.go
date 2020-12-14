package main

import "bufio"
import "fmt"
import "log"
import "os"

func main() {
	sx, sy := 0, 0
	wx, wy := 10, 1

	for line := range inputLines() {
		var c rune
		var n int
		_, err := fmt.Sscanf(line, "%c%d", &c, &n)
		if err != nil {
			log.Fatal(err)
		}

		switch c {
		case 'N':
			wy += n
		case 'S':
			wy -= n
		case 'E':
			wx += n
		case 'W':
			wx -= n
		case 'F':
			sx += wx * n
			sy += wy * n
		case 'L':
			wx, wy = rotate(wx, wy, n)
		case 'R':
			wx, wy = rotate(wx, wy, -n)
		}
	}
	fmt.Println(abs(sx) + abs(sy))
}

func rotate(vx, vy, deg int) (int, int) {
	for deg < 0 {
		deg += 360
	}

	for deg > 0 {
		vx, vy = -vy, vx
		deg -= 90
	}

	return vx, vy
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
