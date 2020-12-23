package main

import "bufio"
import "fmt"
import "io/ioutil"
import "log"
import "os"
import "strings"

type Tile struct {
	id int
	pixels [][]byte
}

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
	tiles := []Tile{}
	for _, tileText := range tileTexts {
		if tileText == "" {
			continue
		}

		id := -1
		pixels := [][]byte{}
		for line := range stringLines(tileText) {
			n, _ := fmt.Sscanf(line, "Tile %d:", &id)
			if n == 1 {
				continue
			}

			row := []byte{}
			for _, c := range line {
				switch c {
				case '.':
					row = append(row, 0)
				case '#':
					row = append(row, 1)
				default:
					log.Fatalf("Unexpected character: %c", c)
				}
			}
			pixels = append(pixels, row)
		}
		t := Tile{id, pixels}
		tiles = append(tiles, t)
	}

	// tile ID => edge IDs
	tileEdges := map[int][]int{}
	// edge ID => tile IDs
	edgeTiles := map[int][]int{}

	for _, tile := range tiles {
		ee := tile.EdgeIDs()
		tileEdges[tile.id] = ee
		for _, e := range ee {
			edgeTiles[e] = append(edgeTiles[e], tile.id)
		}
	}

	cornerProd := uint64(1)
	for _, tile := range tiles {
		neighbors := 0
		for _, e := range tileEdges[tile.id] {
			if len(edgeTiles[e]) > 1 {
				neighbors++
			}
		}
		if neighbors == 2 {
			cornerProd *= uint64(tile.id)
		}
	}
	fmt.Println(cornerProd)
}

// EdgeIDs returns the numbers whose bits are on the edges of the tile.
// Each edge's bits can be used in two orientations; the orientation
// with the lower value is the one returned.
func (t Tile) EdgeIDs() []int {
	n := len(t.pixels)
	topPixels := t.pixels[0]
	bottomPixels := t.pixels[n-1]

	leftPixels := []byte{}
	rightPixels := []byte{}
	for i := range t.pixels {
		leftPixels = append(leftPixels, t.pixels[i][0])
		rightPixels = append(rightPixels, t.pixels[i][n-1])
	}

	return []int{
		edgeID(topPixels),
		edgeID(bottomPixels),
		edgeID(leftPixels),
		edgeID(rightPixels),
	}
}

func edgeID(pixels []byte) int {
	n := len(pixels)
	f := 0
	r := 0
	for i := range pixels {
		r |= int(pixels[i]) << i
		f |= int(pixels[n-1-i]) << i
	}
	//fmt.Println(pixels)
	//fmt.Printf("fwd %08x\n", f)
	//fmt.Printf("rev %08x\n", r)
	if f < r {
		//fmt.Println(">>> fwd")
		return f
	} else {
		//fmt.Println(">>> rev")
		return r
	}
}

func (t Tile) String() string {
	w := &strings.Builder{}
	fmt.Fprintf(w, "Tile %d:\n", t.id)
	for i := range t.pixels {
		fmt.Fprintln(w, t.pixels[i])
	}
	return w.String()
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
