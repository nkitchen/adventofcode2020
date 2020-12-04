package main

import "bufio"
import "fmt"
import "log"
import "os"

func main() {
	treeMap := []string{}

	inFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			treeMap = append(treeMap, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m := len(treeMap)
	n := len(treeMap[0])

	forestAt := func(i, j int) byte {
		return treeMap[i][j % n]
	}

	nHits := []int{}
	for _, slope := range ([][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}) {
		dCol := slope[0]
		dRow := slope[1]

		i := 0
		j := 0
		h := 0
		for i < m {
			if forestAt(i, j) == '#' {
				h++
			}
			i += dRow
			j += dCol
		}

		nHits = append(nHits, h)
	}

	p := 1
	for _, h := range nHits {
		p *= h
	}

	fmt.Println(p)
}
