package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"
import "strings"

func main() {
	maskWhere := uint64(0)
	maskVal := uint64(0)
	mem := map[uint64]uint64{}

	for line := range inputLines() {
		var bits string
		n, err := fmt.Sscanf(line, "mask = %s", &bits)
		if n == 1 {
			w := strings.ReplaceAll(bits, "0", "1")
			w = strings.ReplaceAll(w, "X", "0")
			maskWhere, err = strconv.ParseUint(w, 2, 64)
			if err != nil {
				log.Fatal(err)
			}

			v := strings.ReplaceAll(bits, "X", "0")
			maskVal, err = strconv.ParseUint(v, 2, 64)
			if err != nil {
				log.Fatal(err)
			}

			continue
		}

		var addr, val uint64
		n, err = fmt.Sscanf(line, "mem[%d] = %d", &addr, &val)
		if n == 2 {
			val = val &^ maskWhere | maskVal
			mem[addr] = val
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
