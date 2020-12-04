package main

import "bufio"
import "log"
import "os"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
