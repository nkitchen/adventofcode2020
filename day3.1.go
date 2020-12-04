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

	dCol := 3
	dRow := 1

	nHit := 0
	i := 0
	j := 0
	for i < m {
		if forestAt(i, j) == '#' {
			nHit++
		}
		i += dRow
		j += dCol
	}

	fmt.Println(nHit)
}
