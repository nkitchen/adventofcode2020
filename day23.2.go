package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"

func main() {
	ch := inputLines()
	line := <-ch

	million := 1_000_000
	circle := []int{}
	for _, c := range line {
		circle = append(circle, int(c - '0'))
	}

	succ := make([]int, million + 1)
	for i := 0; i < len(circle) - 1; i++ {
		succ[circle[i]] = circle[i+1]
	}
	succ[circle[len(circle)-1]] = 10
	for i := 10; i < million; i++ {
		succ[i] = i + 1
	}
	succ[million] = circle[0]

	turns := 10_000_000
	if len(os.Args) > 2 {
		turns, _ = strconv.Atoi(os.Args[2])
	}

	cur := circle[0]
	for turn := 0; turn < turns; turn++ {
		pickedUp := make([]int, 3)
		pickedUp[0] = succ[cur]
		pickedUp[1] = succ[pickedUp[0]]
		pickedUp[2] = succ[pickedUp[1]]
		//fmt.Println(cur, pickedUp)

		dest := cur - 1
		for {
			if dest < 1 {
				dest = million
			}
			if rindex(pickedUp, dest) == -1 {
				break
			} else {
				dest--
			}
		}

		succ[cur] = succ[pickedUp[2]]
		succ[pickedUp[2]] = succ[dest]
		succ[dest] = pickedUp[0]

		cur = succ[cur]
	}
	s := succ[1]
	p := s * succ[s]
	fmt.Println(p)
}

func rindex(a []int, c int) int {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == c {
			return i
		}
	}
	return -1
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
