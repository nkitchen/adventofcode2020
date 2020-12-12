package main

import "bufio"
import "fmt"
import "log"
import "os"
//import "strconv"

func main() {
	grid := [][]byte{}
	for line := range inputLines() {
		row := []byte(line)
		grid = append(grid, row)
	}

	type rc struct {
		r, c int
	}

	show(grid)
	for {
		next := make([][]byte, len(grid))

		changed := map[rc]bool{}
		for r := range grid {
			next[r] = make([]byte, len(grid[r]))
			for c := range grid[r] {
				next[r][c] = grid[r][c]

				if grid[r][c] == 'L' &&
					occupiedVisible(grid, r, c) == 0 {
					next[r][c] = '#'
					changed[rc{r, c}] = true
				}

				if grid[r][c] == '#' &&
					occupiedVisible(grid, r, c) >= 5 {
					next[r][c] = 'L'
					changed[rc{r, c}] = true
				}
			}
		}

		if len(changed) == 0 {
			break
		}

		grid = next
		show(grid)
	}

	occ := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '#' {
				occ++
			}
		}
	}
	fmt.Println(occ)
}

var verbose = false
func show(grid [][]byte) {
	if !verbose {
		return
	}

	for r := range grid {
		fmt.Println(string(grid[r]))
	}
	fmt.Println()
}

func occupiedVisible(grid [][]byte, row, col int) int {
	k := 0
	for drc := range rowColDels() {
		dr := drc[0]
		dc := drc[1]
		r := row + dr
		c := col + dc
		for 0 <= r && r < len(grid) &&
			0 <= c && c < len(grid[r]) {
			if grid[r][c] == 'L' {
				break
			}
			if grid[r][c] == '#' {
				k++
				break
			}
			r += dr
			c += dc
		}
	}
	return k
}

func rowColDels() <-chan [2]int {
	ch := make(chan [2]int)

	go func() {
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dr == 0 && dc == 0 {
					continue
				}
				ch <- [2]int{dr, dc}
			}
		}
		close(ch)
	}()

	return ch
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
