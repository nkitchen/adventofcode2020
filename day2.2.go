package main

import "bufio"
import "fmt"
import "log"
import "os"

func main() {
	nValid := 0

	inFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()

		var pos1 int
		var pos2 int
		var letter rune
		var password string
		fmt.Sscanf(line, "%d-%d %c: %s", &pos1, &pos2, &letter, &password)

		k := 0
		if rune(password[pos1-1]) == letter {
			k++
		}
		if rune(password[pos2-1]) == letter {
			k++
		}
		if k == 1 {
			nValid++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(nValid)
}
