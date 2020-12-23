package main

import "bufio"
import "fmt"
import "io/ioutil"
import "log"
import "math"
import "os"
import "strconv"
import "strings"

func main() {
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(inFile)
	if err != nil {
		log.Fatal(err)
	}
	inText := string(buf)

	tileTexts := strings.Split(inText, "\n\n")
	tiles := []*Tile{}
	for _, tileText := range tileTexts {
		if tileText == "" {
			continue
		}

		id := -1
		pixels := []string{}
		for line := range stringLines(tileText) {
			n, _ := fmt.Sscanf(line, "Tile %d:", &id)
			if n == 1 {
				continue
			}

			pixels = append(pixels, line)
		}
		t := &Tile{id, pixels}
		tiles = append(tiles, t)
	}

	// left edge => oriented tiles
	orientedByLeft := map[string][]*OrientedTile{}
	// top edge => oriented tiles
	orientedByTop := map[string][]*OrientedTile{}
	for _, tile := range tiles {
		for _, ot := range tile.Orientations() {
			left := ot.LeftEdge()
			orientedByLeft[left] = append(orientedByLeft[left], ot)
			top := ot.TopEdge()
			orientedByTop[top] = append(orientedByTop[top], ot)
		}
	}


	// Get id and edge of seed tile from arguments.
	// We should know them from part 1 (e.g., Tile 1987, edge .#.#.....#).
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s FILE TILE EDGE\n", os.Args[0])
		os.Exit(1)
	}
	seedID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	seedEdge := os.Args[3]

	var seedTile *OrientedTile
	for _, ot := range orientedByLeft[seedEdge] {
		if ot.tile.id == seedID {
			seedTile = ot
			break
		}
	}
	if seedTile == nil {
		log.Fatalf("Tile %d with edge %q not found", seedID, seedEdge)
	}

	m := int(math.Sqrt(float64(len(tiles))))
	grid := make([][]*OrientedTile, m)

	for i := 0; i < m; i++ {
		grid[i] = make([]*OrientedTile, m)
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				grid[i][j] = seedTile
				continue
			}

			if j == 0 {
				topNeighbor := grid[i-1][j]
				top := topNeighbor.BottomEdge()
				for _, ot := range orientedByTop[top] {
					if ot.tile.id == topNeighbor.tile.id {
						// Reflection of that tile
						continue
					}
					grid[i][j] = ot
					break
				}
			} else {
				leftNeighbor := grid[i][j-1]
				left := leftNeighbor.RightEdge()
				for _, ot := range orientedByLeft[left] {
					if ot.tile.id == leftNeighbor.tile.id {
						// Reflection of that tile
						continue
					}
					grid[i][j] = ot
					break
				}
			}
			if grid[i][j] == nil {
				log.Fatalf("Matching tile not found: i=%d j=%d", i, j)
			}
		}
	}

	// Assemble the image.
	n := len(grid[0][0].tile.pixels)
	img := []string{}
	for i := 0; i < m * n; i++ {
		if i % n == 0 || i % n == n - 1 {
			// edge
			continue
		}
		row := []byte{}
		for j := 0; j < m * n; j++ {
			if j % n == 0 || j % n == n - 1 {
				// edge
				continue
			}
			b := grid[i / n][j / n].At(i % n, j % n)
			row = append(row, b)
		}
		img = append(img, string(row))
	}
	m = len(img)
	n = len(img[0])

	monster := Tile{
		id: 0,
		pixels: []string{
			"                  # ",
			"#    ##    ##    ###",
			" #  #  #  #  #  #   ",
		},
	}

	matchesByTemplate := map[*OrientedTile][]Pos{}
	for _, om := range monster.Orientations() {
		matches := []Pos{}
		// Start position of template
		for i := 0; i + om.M() - 1 < m; i++ {
			for j := 0; j + om.N() - 1 < n; j++ {
				if matchTemplate(om, img, i, j) {
					matches = append(matches, Pos{i, j})
				}
			}
		}
		matchesByTemplate[om] = matches
	}
	
	bestMatches := 0
	for om := range matchesByTemplate {
		k := len(matchesByTemplate[om])
		if k > bestMatches {
			bestMatches = k
		}
	}

	monsterPixels := 0
	for i := 0; i < monster.M(); i++ {
		for j := 0; j < monster.N(); j++ {
			if monster.pixels[i][j] == '#' {
				monsterPixels++
			}
		}
	}

	roughPixels := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if img[i][j] == '#' {
				roughPixels++
			}
		}
	}

	roughness := roughPixels - bestMatches * monsterPixels
	fmt.Println(roughness)
}

type Tile struct {
	id int
	pixels []string
}

// A Tile as seen after rotation and flipping
type OrientedTile struct {
	tile *Tile
	// i: 0..n-1 => n-1..0 ?
	revI bool
	// j: 0..n-1 => n-1..0 ?
	revJ bool
	// i, j => j, i ?
	swap bool
}

func (t *Tile) Orientations() []*OrientedTile {
	orientations := []*OrientedTile{}
	for b := 0; b <= 7; b++ {
		ot := &OrientedTile{
			tile: t,
			revI: b & 1 == 1,
			revJ: b & 2 == 2,
			swap: b & 4 == 4,
		}
		orientations = append(orientations, ot)
	}
	return orientations
}

func (t *Tile) M() int {
	return len(t.pixels)
}

func (t *Tile) N() int {
	return len(t.pixels[0])
}

func (ot *OrientedTile) M() int {
	if ot.swap {
		return ot.tile.N()
	} else {
		return ot.tile.M()
	}
}

func (ot *OrientedTile) N() int {
	if ot.swap {
		return ot.tile.M()
	} else {
		return ot.tile.N()
	}
}

func (ot *OrientedTile) At(i, j int) byte {
	ti := i
	tj := j
	m := ot.M()
	n := ot.N()
	if ot.swap {
		ti, tj = tj, ti
		m, n = n, m
	}
	if ot.revI {
		ti = m - 1 - ti
	}
	if ot.revJ {
		tj = n - 1 - tj
	}
	return ot.tile.pixels[ti][tj]
}

func (ot *OrientedTile) TopEdge() string {
	n := ot.N()
	b := make([]byte, n)
	for j := 0; j < n; j++ {
		b[j] = ot.At(0, j)
	}
	return string(b)
}

func (ot *OrientedTile) BottomEdge() string {
	n := ot.N()
	b := make([]byte, n)
	for j := 0; j < n; j++ {
		b[j] = ot.At(n - 1, j)
	}
	return string(b)
}

func (ot *OrientedTile) LeftEdge() string {
	m := ot.M()
	b := make([]byte, m)
	for i := 0; i < m; i++ {
		b[i] = ot.At(i, 0)
	}
	return string(b)
}

func (ot *OrientedTile) RightEdge() string {
	m := ot.M()
	b := make([]byte, m)
	for i := 0; i < m; i++ {
		b[i] = ot.At(i, m - 1)
	}
	return string(b)
}

func (ot *OrientedTile) String() string {
	w := &strings.Builder{}
	fmt.Fprintf(w, "OrientedTile{id=%d,\n", ot.tile.id)
	m := ot.M()
	n := ot.N()
	for i := 0; i < m; i++ {
		fmt.Fprint(w, "  ")
		for j := 0; j < n; j++ {
			fmt.Fprintf(w, "%c", ot.At(i, j))
		}
		fmt.Fprintln(w)
	}
	fmt.Fprint(w, "}")
	return w.String()
}

type Pos struct {
	i, j int
}

func matchTemplate(template *OrientedTile, img []string, i, j int) bool {
	m := template.M()
	n := template.N()
	for u := 0; u < m; u++ {
		for v := 0; v < n; v++ {
			if template.At(u, v) != '#' {
				continue
			}
			if img[i + u][j + v] != '#' {
				return false
			}
		}
	}
	return true
}

func stringLines(s string) <-chan string {
	ch := make(chan string)

	go func() {
		scanner := bufio.NewScanner(strings.NewReader(s))
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
