package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"
import "strings"

func main() {
	maskSet := uint64(0)
	maskFloating := []int{}
	mem := map[uint64]uint64{}

	for line := range inputLines() {
		var bits string
		n, err := fmt.Sscanf(line, "mask = %s", &bits)
		if n == 1 {
			s := strings.ReplaceAll(bits, "X", "0")
			maskSet, err = strconv.ParseUint(s, 2, 64)
			if err != nil {
				log.Fatal(err)
			}

			m := len(bits)
			maskFloating = []int{}
			for i := 0; i < 36; i++ {
				if bits[m - 1 - i] == 'X' {
					maskFloating = append(maskFloating, i)
				}
			}

			continue
		}

		var addr, val uint64
		n, err = fmt.Sscanf(line, "mem[%d] = %d", &addr, &val)
		if n == 2 {
			addr |= maskSet
			floated := []uint64{addr}
			for _, k := range maskFloating {
				b := uint64(1) << k
				expanded := []uint64{}
				for _, f := range floated {
					expanded = append(expanded, f &^ b)
					expanded = append(expanded, f | b)
				}
				floated = expanded
			}
			for _, f := range floated {
				mem[f] = val
			}
		}
	}

	s := uint64(0)
	for _, val := range mem {
		s += val
	}
	fmt.Println(s)
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
