package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"
import "strings"

func main() {
	ch := inputLines()
	line1 := <-ch
	t, err := strconv.Atoi(line1)
	if err != nil {
		log.Fatal(err)
	}

	type bus struct {
		id int
		wait int
	}

	buses := []bus{}
	line2 := <-ch
	for _, s := range strings.Split(line2, ",") {
		if s == "x" {
			continue
		}
		b, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		wait := (b - t % b) % b
		buses = append(buses, bus{b, wait})
	}

	earliest := buses[0]
	for _, b := range buses[1:] {
		if b.wait < earliest.wait {
			earliest = b
		}
	}
	fmt.Println(earliest.id * earliest.wait)
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
