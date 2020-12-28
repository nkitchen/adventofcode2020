package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"

func main() {
	decks := [2][]int32{}
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
			decks[i] = append(decks[i], int32(n))
		}
	}

	winner := play(&decks[0], &decks[1])
	winning := decks[winner]

	score := 0
	for k := 1; k <= len(winning); k++ {
		score += k * int(winning[len(winning)-k])
	}
	fmt.Println(score)
}


func play(deck0, deck1 *[]int32) int {
	history := map[[2]string]bool{}

	for len(*deck0) > 0 && len(*deck1) > 0 {
		h := [2]string{string(*deck0), string(*deck1)}
		if history[h] {
			return 0
		}
		history[h] = true

		card0 := (*deck0)[0]
		card1 := (*deck1)[0]
		*deck0 = (*deck0)[1:]
		*deck1 = (*deck1)[1:]

		var w int
		if card0 <= int32(len(*deck0)) &&
		    card1 <= int32(len(*deck1)) {
				subdeck0 := make([]int32, int(card0))
				subdeck1 := make([]int32, int(card1))
				copy(subdeck0, (*deck0)[:int(card0)])
				copy(subdeck1, (*deck1)[:int(card1)])
				w = play(&subdeck0, &subdeck1)
		} else if card0 > card1 {
			w = 0
		} else {
			w = 1
		}

		if w == 0 {
			*deck0 = append(*deck0, card0, card1)
		} else {
			*deck1 = append(*deck1, card1, card0)
		}
	}

	if len(*deck0) > 0 {
		return 0
	} else {
		return 1
	}
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
