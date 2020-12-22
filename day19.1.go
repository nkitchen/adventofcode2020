package main

import "bufio"
import "bytes"
import "fmt"
import "log"
import "os"
import "strconv"
import "strings"

type Rule struct {
	lit string
	branches [][]int
}

type RuleSet map[int]Rule

func main() {
	ruleSet := RuleSet{}
	messages := []string{}
	for line := range inputLines() {
		i := strings.Index(line, ": ")
		if i != -1 {
			rule := Rule{}
			
			n, err := strconv.Atoi(line[:i])
			if err != nil {
				log.Fatal(err)
			}

			body := line[i+2:]
			k := len(body)
			if body[0] == '"' && body[k-1] == '"' {
				rule.lit = body[1:k-1]
			} else {
				branches := [][]int{}
				for _, branchStr := range strings.Split(body, " | ") {
					refs := []int{}
					for _, w := range strings.Fields(branchStr) {
						m, err := strconv.Atoi(w)
						if err != nil {
							log.Fatal(err)
						}
						refs = append(refs, m)
					}
					branches = append(branches, refs)
				}
				rule.branches = branches
			}
			ruleSet[n] = rule
			continue
		} 

		if line == "" {
			continue
		}

		messages = append(messages, line)
	}

	n := 0
	for _, msg := range messages {
		if ruleSet.match(msg) {
			n++
		}
	}
	fmt.Println(n)
}

func (rs RuleSet) match(t string) bool {
	ok, u := rs.matchRule(0, []byte(t))
	return ok && len(u) == 0
}

func (rs RuleSet) matchRule(n int, t []byte) (bool, []byte) {
	rule := rs[n]
	if rule.lit != "" {
		k := len(rule.lit)
		if bytes.Equal([]byte(rule.lit), t[:k]) {
			return true, t[k:]
		} else {
			return false, t
		}
	}

	for _, branch := range rule.branches {
		ok, u := rs.matchBranch(branch, t)
		if ok {
			return ok, u
		}
	}
	return false, t
}

func (rs RuleSet) matchBranch(branch []int, t []byte) (bool, []byte) {
	u := t
	for _, r := range branch {
		ok, v := rs.matchRule(r, u)
		if !ok {
			return false, t
		}
		u = v
	}
	return true, u
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
