package main

import "bufio"
import "fmt"
import "log"
import "os"
import "regexp"
import "strconv"
import "strings"

type Range struct {
	first, last int
}

type MultiRange []Range

type ticket []int

func main() {
	ch := inputLines()

	fieldNames := []string{}
	fieldRanges := []MultiRange{}
	myTicket := ticket{}
	_ = myTicket
	nearbyTickets := []ticket{}

	fieldRe := regexp.MustCompile(`^(.*): (\d+-\d+ or \d+-\d+)`)
Input:
	for {
		line := <-ch
		m := fieldRe.FindStringSubmatch(line)
		if len(m) != 0 {
			fieldNames = append(fieldNames, m[1])

			var a, b, c, d int
			_, err := fmt.Sscanf(m[2], "%d-%d or %d-%d", &a, &b, &c, &d)
			if err != nil {
				log.Fatal(err)
			}
			mr := MultiRange{}
			mr = append(mr, Range{a, b})
			mr = append(mr, Range{c, d})
			fieldRanges = append(fieldRanges, mr)
		}

		if line == "your ticket:" {
			line = <-ch
			myTicket = newTicket(line)
		}

		if line == "nearby tickets:" {
			for {
				line, ok := <-ch
				if !ok {
					break Input
				}
				nearbyTickets = append(nearbyTickets, newTicket(line))
			}
		}
	}

	// validTickets[f][pos]
	validTickets := map[[2]int]int{}
	for _, t := range nearbyTickets {
		// validWhere[pos]: fields for which t[pos] is valid
		validWhere := make([][]int, len(t))
		nValid := 0
		for pos, n := range t {
			for f, mr := range fieldRanges {
				if mr.Contains(n) {
					validWhere[pos] = append(validWhere[pos], f)
				}
			}
			if len(validWhere[pos]) > 0 {
				nValid++
			}
		}
		if nValid < len(t) {
			continue
		}

		for pos := range validWhere {
			for _, f := range validWhere[pos] {
				k := [2]int{f, pos}
				validTickets[k] = 1 + validTickets[k]
			}
		}
	}

	m := 0
	for _, n := range validTickets {
		if n > m {
			m = n
		}
	}

	// poss[pos][f]
	N := len(fieldRanges)
	poss := map[int]map[int]bool{}
	for pos := 0; pos < N; pos++ {
		poss[pos] = map[int]bool{}
		for f := 0; f < N; f++ {
			k := [2]int{f, pos}
			if validTickets[k] == m {
				poss[pos][f] = true
			}
		}
	}

	removePoss := func(pos, f int) {
		delete(poss[pos], f)
		if len(poss[pos]) == 0 {
			delete(poss, pos)
		}
	}

	fieldPos := make([]int, N)
	for len(poss) > 0 {
		for pos := 0; pos < N; pos++ {
			if len(poss[pos]) == 1 {
				for f := range poss[pos] {
					fieldPos[f] = pos
					removePoss(pos, f)
					for pp := 0; pp < N; pp++ {
						removePoss(pp, f)
					}
					delete(poss, pos)
					break
				}
			}
		}
	}

	p := 1
	for f, name := range fieldNames {
		if strings.Index(name, "departure") != 0 {
			continue
		}
		n := myTicket[fieldPos[f]]
		fmt.Println(n)
		p *= n
	}
	fmt.Println(p)
}

func newTicket(csv string) ticket {
	t := ticket{}
	for _, s := range strings.Split(csv, ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		t = append(t, n)
	}
	return t
}

func (mr MultiRange) Contains(x int) bool {
	for _, r := range mr {
		if r.first <= x && x <= r.last {
			return true
		}
	}
	return false
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
