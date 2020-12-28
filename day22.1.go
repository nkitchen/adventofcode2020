package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"

func main() {
	decks := [2][]int{}
	i := -1
	for line := range inputLines() {
		switch line {
		case "Player 1:":
			i = 0
		case "Player 2:":
			i = 1
		case "":
			// nothing
		default:
			n, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			decks[i] = append(decks[i], n)
		}
	}

	for len(decks[0]) > 0 && len(decks[1]) > 0 {
		card0 := decks[0][0]
		decks[0] = decks[0][1:]
		card1 := decks[1][0]
		decks[1] = decks[1][1:]
		if card0 > card1 {
			decks[0] = append(decks[0], card0, card1)
		} else {
			decks[1] = append(decks[1], card1, card0)
		}
	}

	var winning []int
	if len(decks[0]) > 0 {
		winning = decks[0]
	} else {
		winning = decks[1]
	}
	score := 0
	for k := 1; k <= len(winning); k++ {
		score += k * winning[len(winning)-k]
	}
	fmt.Println(score)
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
