package main

import "bufio"
import "fmt"
import "log"
import "math/big"
import "os"
import "strconv"

func main() {
	ch := inputLines()

	line := <-ch
	pubKeyCard64, err := strconv.ParseInt(line, 0, 64)
	if err != nil {
		log.Fatal(err)
	}
	pubKeyCard := big.NewInt(pubKeyCard64)

	line = <-ch
	pubKeyDoor64, err := strconv.ParseInt(line, 0, 64)
	if err != nil {
		log.Fatal(err)
	}
	pubKeyDoor := big.NewInt(pubKeyDoor64)

	seven := big.NewInt(7)
	mod := big.NewInt(20201227)

	loopSizeCard := discLog(pubKeyCard, seven, mod)
	loopSizeDoor := discLog(pubKeyDoor, seven, mod)
	
	n := &big.Int{}
	n.Mul(loopSizeCard, loopSizeDoor)
	encKey := &big.Int{}
	encKey.Exp(big.NewInt(7), n, mod)
	fmt.Println(encKey)
}

func discLog(a, b, m *big.Int) *big.Int {
	one := big.NewInt(1)
	for k := big.NewInt(1); k.Cmp(m) < 0; k = k.Add(k, one) {
		t := (&big.Int{}).Exp(b, k, m)
		if t.Cmp(a) == 0 {
			return k
		}
	}
	return nil
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
