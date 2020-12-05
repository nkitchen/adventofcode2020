package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strings"
import "strconv"

func main() {
	maxSeat := -1

	inFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		s := line
		s = strings.ReplaceAll(s, "F", "0")
		s = strings.ReplaceAll(s, "B", "1")
		s = strings.ReplaceAll(s, "L", "0")
		s = strings.ReplaceAll(s, "R", "1")

		id64, err := strconv.ParseInt(s, 2, 32)
		if err != nil {
			log.Fatal(err)
		}

		id := int(id64)
		if id > maxSeat {
			maxSeat = id
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(maxSeat)
}
