package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	entries := map[int]bool{}
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		entries[n] = true
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for n := range entries {
		m := 2020 - n
		if entries[m] {
			fmt.Println(m * n)
		}
	}
}
