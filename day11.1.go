package main

import "bufio"
import "fmt"
import "log"
import "os"

func main() {
	grid := [][]byte{}
	for line := range inputLines() {
		if len(grid) == 0 {
			n := len(line) + 2
			row := make([]byte, n)
			for i := range row {
				row[i] = '.'
			}
			grid = append(grid, row)
		}

		row := []byte("." + line + ".")
		grid = append(grid, row)
	}

	row := []byte(string(grid[0]))
	grid = append(grid, row)

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
					occupiedAdjacent(grid, r, c) == 0 {
					next[r][c] = '#'
					changed[rc{r, c}] = true
				}

				if grid[r][c] == '#' &&
					occupiedAdjacent(grid, r, c) >= 4 {
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

func occupiedAdjacent(grid [][]byte, row, col int) int {
	k := 0
	for rr := row - 1; rr <= row+1; rr++ {
		for cc := col - 1; cc <= col+1; cc++ {
			if rr == row && cc == col {
				continue
			}
			if grid[rr][cc] == '#' {
				k++
			}
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
