package main

import "bufio"
import "fmt"
import "log"
import "os"

//import "strconv"

type rc struct {
	r, c int
}

type layout struct {
	grid         [][]byte
	visibleSeats map[rc][]rc
}

func main() {
	grid := [][]byte{}
	for line := range inputLines() {
		row := []byte(line)
		grid = append(grid, row)
	}
	lay := newLayout(grid)
	lay.show()

	toCheck := map[rc]bool{}
	for r := range lay.grid {
		for c := range lay.grid[r] {
			toCheck[rc{r, c}] = true
		}
	}

	for {
		changed := map[rc]byte{}
		for rc := range toCheck {
			r := rc.r
			c := rc.c
			if lay.grid[r][c] == 'L' &&
				lay.occupiedVisible(r, c) == 0 {
				changed[rc] = '#'
			}

			if lay.grid[r][c] == '#' &&
				lay.occupiedVisible(r, c) >= 5 {
				changed[rc] = 'L'
			}
		}

		if len(changed) == 0 {
			break
		}

		for rc, b := range changed {
			lay.grid[rc.r][rc.c] = b
		}
		lay.show()

		toCheck = map[rc]bool{}
		for rc := range changed {
			for _, rrcc := range lay.visibleSeats[rc] {
				toCheck[rrcc] = true
			}
		}
	}

	occ := 0
	for r := range lay.grid {
		for c := range lay.grid[r] {
			if lay.grid[r][c] == '#' {
				occ++
			}
		}
	}
	fmt.Println(occ)
}

var verbose = false

func (lay *layout) show() {
	if !verbose {
		return
	}

	for r := range lay.grid {
		fmt.Println(string(lay.grid[r]))
	}
	fmt.Println()
}

var rowColVecs = func() [][2]int {
	vecs := [][2]int{}
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			vecs = append(vecs, [2]int{dr, dc})
		}
	}
	return vecs
}()

func newLayout(grid [][]byte) *layout {
	lay := &layout{
		grid:         grid,
		visibleSeats: map[rc][]rc{},
	}

	m := len(lay.grid)
	for r := range lay.grid {
		n := len(lay.grid[r])
		for c := range lay.grid[r] {
			vis := []rc{}
			for _, drc := range rowColVecs {
				dr := drc[0]
				dc := drc[1]
				rr := r + dr
				cc := c + dc
			SightLine:
				for 0 <= rr && rr < m &&
					0 <= cc && cc < n {
					switch lay.grid[rr][cc] {
					case 'L', '#':
						vis = append(vis, rc{rr, cc})
						break SightLine
					}
					rr += dr
					cc += dc
				}
			}
			lay.visibleSeats[rc{r, c}] = vis
		}
	}
	return lay
}

func (lay *layout) occupiedVisible(row, col int) int {
	k := 0
	for _, rc := range lay.visibleSeats[rc{row, col}] {
		r := rc.r
		c := rc.c
		if lay.grid[r][c] == '#' {
			k++
		}
	}
	return k
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
