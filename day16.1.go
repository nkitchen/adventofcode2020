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

type ticket []int

func main() {
	ch := inputLines()

	valid := []Range{}
	myTicket := ticket{}
	_ = myTicket
	nearbyTickets := []ticket{}

	fieldRe := regexp.MustCompile(`^.*: (\d+-\d+ or \d+-\d+)`)
Input:
	for {
		line := <-ch
		m := fieldRe.FindStringSubmatch(line)
		if len(m) != 0 {
			var a, b, c, d int
			_, err := fmt.Sscanf(m[1], "%d-%d or %d-%d", &a, &b, &c, &d)
			if err != nil {
				log.Fatal(err)
			}
			valid = append(valid, Range{a, b})
			valid = append(valid, Range{c, d})
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

	errorRate := 0
	for _, t := range nearbyTickets {
		for _, n := range t {
			v := false
			for _, r := range valid {
				if r.first <= n && n <= r.last {
					v = true
					break
				}
			}
			if !v {
				errorRate += n
			}
		}
	}
	fmt.Println(errorRate)
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
