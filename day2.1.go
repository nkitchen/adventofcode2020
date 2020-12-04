package main

import "bufio"
import "fmt"
import "log"
import "os"

func main() {
	nValid := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var min int
		var max int
		var letter rune
		var password string
		fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &letter, &password)

		k := map[rune]int{}
		for _, c := range password {
			k[c]++
		}

		if min <= k[letter] && k[letter] <= max {
			nValid++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(nValid)
}
