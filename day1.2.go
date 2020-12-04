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

	for x := range entries {
		for y := range entries {
			z := 2020 - x - y
			if entries[z] {
				fmt.Println(x * y * z)
			}
		}
	}
}
