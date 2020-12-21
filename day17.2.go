package main

import "bufio"
import "fmt"
import "log"
import "os"

var Debug = false

type Coords struct {
	x, y, z, w int
}

func cellIter(min, max Coords, f func(Coords)) {
	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			for z := min.z; z <= max.z; z++ {
				for w := min.w; w <= max.w; w++ {
					f(Coords{x, y, z, w})
				}
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minCoords(c, d Coords) Coords {
	return Coords{
		min(c.x, d.x),
		min(c.y, d.y),
		min(c.z, d.z),
		min(c.w, d.w),
	}
}

func maxCoords(c, d Coords) Coords {
	return Coords{
		max(c.x, d.x),
		max(c.y, d.y),
		max(c.z, d.z),
		max(c.w, d.w),
	}
}

func coordsInc(c Coords, k int) Coords {
	return Coords{c.x + k, c.y + k, c.z + k, c.w + k}
}

func main() {
	cell := map[Coords]int{}

	y := 0
	bboxMin := Coords{0, 0, 0, 0}
	bboxMax := Coords{0, 0, 0, 0}
	for line := range inputLines() {
		for x, r := range line {
			if r == '#' {
				c := Coords{x, y, 0, 0}
				cell[c] = 1
				bboxMin = minCoords(bboxMin, c)
				bboxMax = maxCoords(bboxMax, c)
			}
		}
		y++
	}

	show := func() {
		fmt.Println("==")
		for z := bboxMin.z; z <= bboxMax.z; z++ {
			fmt.Printf("z=%v\n", z)
			for y := bboxMin.y; y <= bboxMax.y; y++ {
				b := []rune{}
				for x := bboxMin.x; x <= bboxMax.x; x++ {
					if cell[Coords{x, y, z, 0}] == 1 {
						b = append(b, '#')
					} else {
						b = append(b, '.')
					}
				}
				fmt.Println(string(b))
			}
			fmt.Println()
		}
		fmt.Println()
	}

	for cycle := 0; cycle < 6; cycle++ {
		changes := map[Coords]int{}
		cellIter(coordsInc(bboxMin, -1), coordsInc(bboxMax, +1), func(c Coords) {
			neighbors := 0
			cellIter(coordsInc(c, -1), coordsInc(c, +1), func(d Coords) {
				if d == c {
					return
				}
				neighbors += cell[d]
			})
			if cell[c] == 1 && !(2 <= neighbors && neighbors <= 3) {
				changes[c] = 0
			}
			if cell[c] == 0 && neighbors == 3 {
				changes[c] = 1
			}
		})
		for c, v := range changes {
			if v == 1 {
				cell[c] = v
				bboxMin = minCoords(bboxMin, c)
				bboxMax = maxCoords(bboxMax, c)
			} else {
				delete(cell, c)
			}
		}
		if Debug {
			show()
		}
	}

	fmt.Println(len(cell))
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
