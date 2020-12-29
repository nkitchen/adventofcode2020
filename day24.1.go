package main

import "bufio"
import "fmt"
import "log"
import "os"

func main() {
	tileColor := map[pos]color{}
	for line := range inputLines() {
		p := find(line)
		tileColor[p] = 1 ^ tileColor[p]
	}

	n := 0
	for _, col := range tileColor {
		n += int(col)
	}
	fmt.Println(n)
}

type pos struct {
	x, y int
}

type color int
const (
	white color = 0
	black color = 1
)

func find(path string) pos {
	p := pos{0, 0}
	for len(path) > 0 {
		switch path[0] {
		case 'w':
			p.x -= 2
			path = path[1:]
			continue
		case 'e':
			p.x += 2
			path = path[1:]
			continue
		}

		switch path[:2] {
		case "nw":
			p.x--
			p.y++
		case "sw":
			p.x--
			p.y--
		case "ne":
			p.x++
			p.y++
		case "se":
			p.x++
			p.y--
		}
		path = path[2:]
	}
	return p
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
