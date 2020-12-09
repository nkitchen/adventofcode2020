package main

import "bufio"
import "fmt"
import "log"
import "os"

type inst struct {
	op string
	arg int
}

func main() {
	insts := []inst{}
	for line := range inputLines() {
		ii := inst{}
		fmt.Sscanf(line, "%s %d", &ii.op, &ii.arg)
		insts = append(insts, ii)
	}

	visited := map[int]bool{}
	ip := 0
	acc := 0
	for {
		if visited[ip] {
			break
		}
		visited[ip] = true

		switch insts[ip].op {
		case "acc":
			acc += insts[ip].arg
			fallthrough
		case "nop":
			ip++
		case "jmp":
			ip += insts[ip].arg
		}
	}
	fmt.Println(acc)
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
