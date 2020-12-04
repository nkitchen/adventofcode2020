package main

import "bufio"
import "fmt"
import "log"
import "os"

func main() {
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()


	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
