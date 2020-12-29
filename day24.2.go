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

	for p := range tileColor {
		if tileColor[p] == white {
			delete(tileColor, p)
		}
	}

	//show(tileColor)
	for day := 1; day <= 100; day++ {
		whiteTiles := map[pos]bool{}
		changes := map[pos]color{}

		for p := range tileColor {
			b := 0
			for _, nbr := range neighbors(p) {
				nc := tileColor[nbr]
				b += int(nc)
				if nc == white {
					whiteTiles[nbr] = true
				}
			}
			if tileColor[p] == black && (b == 0 || b > 2) {
				changes[p] = white
			}
		}

		for p := range whiteTiles {
			b := 0
			for _, nbr := range neighbors(p) {
				b += int(tileColor[nbr])
			}
			if b == 2 {
				changes[p] = black
			}
		}

		for p, col := range changes {
			if col == black {
				tileColor[p] = black
			} else {
				delete(tileColor, p)
			}
		}

		//fmt.Printf("Day %d: %d\n", day, len(tileColor))
		//show(tileColor)
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

func neighbors(p pos) []pos {
	return []pos{
		{p.x - 2, p.y},
		{p.x + 2, p.y},
		{p.x - 1, p.y - 1},
		{p.x - 1, p.y + 1},
		{p.x + 1, p.y - 1},
		{p.x + 1, p.y + 1},
	}
}

func show(tiles map[pos]color) {
	xMin := 10000
	xMax := -10000
	yMin := 10000
	yMax := -10000

	for p := range tiles {
		if p.x < xMin {
			xMin = p.x
		}
		if p.x > xMax {
			xMax = p.x
		}
		if p.y < yMin {
			yMin = p.y
		}
		if p.y > yMax {
			yMax = p.y
		}
	}

	for y := yMax; y >= yMin; y-- {
		for x := xMin; x <= xMax; x++ {
			switch {
			case x == 1 && y == 0:
				fmt.Print("<")
			case (x + y) % 2 != 0:
				fmt.Print(" ")
			case tiles[pos{x, y}] == black:
				fmt.Print("#")
			case tiles[pos{x, y}] == white:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
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
