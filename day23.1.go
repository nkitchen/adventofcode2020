package main

import "bufio"
import "bytes"
import "fmt"
import "log"
import "os"
import "strconv"

func main() {
	ch := inputLines()
	line := <-ch

	circle := []byte(line)

	turns := 100
	if len(os.Args) > 2 {
		turns, _ = strconv.Atoi(os.Args[2])
	}

	n := len(circle)
	for turn := 0; turn < turns; turn++ {
		// Invariant: The circle is rotated so that the current item is [0].
		//fmt.Printf("%s\n", circle)

		pickedUp := circle[1:4]

		destLabel := circle[0] - 1
		for {
			if destLabel < '1' {
				destLabel = '9'
			}
			if bytes.IndexByte(pickedUp, destLabel) == -1 {
				break
			} else {
				destLabel--
			}
		}
		dest := bytes.IndexByte(circle, destLabel)

		newCircle := make([]byte, 0, n)
		newCircle = append(newCircle, circle[4:dest+1]...)
		newCircle = append(newCircle, pickedUp...)
		newCircle = append(newCircle, circle[dest+1:]...)
		newCircle = append(newCircle, circle[0])
		circle = newCircle
	}

	i := bytes.IndexByte(circle, '1')
	s := string(circle[i+1:]) + string(circle[:i])
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
