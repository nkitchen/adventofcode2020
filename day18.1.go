package main

import "bufio"
import "fmt"
import "log"
import "os"
import "regexp"
import "strconv"

func main() {
	tokenRe := regexp.MustCompile(`\d+|[()*+]`)
	s := 0
	for line := range inputLines() {
		tokens := tokenRe.FindAllString(line, -1)
		v, _, err := evalExpr(tokens)
		if err != nil {
			log.Fatal(err)
		}
		s += v
	}
	fmt.Println(s)
}

func evalExpr(tokens []string) (value int, nextTokens []string, err error) {
	debugf("evalExpr %v\n", tokens)

	value, tokens, err = evalArg(tokens)
	if err != nil {
		return 0, tokens, err
	}

	var rhs int
	for len(tokens) > 0 {
		if tokens[0] == ")" {
			return value, tokens, nil
		}

		op := tokens[0]
		rhs, tokens, err = evalArg(tokens[1:])
		if err != nil {
			return 0, tokens, err
		}
		switch op {
		case "+":
			value += rhs
		case "*":
			value *= rhs
		default:
			return 0, tokens, fmt.Errorf("Unexpected token %q", op)
		}
	}

	return value, tokens, nil
}

func evalArg(tokens []string) (value int, nextTokens []string, err error) {
	debugf("evalArg %v\n", tokens)

	if tokens[0] == "(" {
		value, tokens, err = evalExpr(tokens[1:])
		if err != nil {
			return 0, tokens, err
		}

		t := tokens[0]
		if t != ")" {
			return 0, tokens, fmt.Errorf(`Expected ")", got %q`, t)
		}
		return value, tokens[1:], nil
	}

	value, err = strconv.Atoi(tokens[0])
	return value, tokens[1:], err
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

var debug = false
func debugf(format string, a ...interface{}) {
	if debug {
		fmt.Printf(format, a...)
	}
}
