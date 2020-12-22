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
	mm := rs.matchRule(0, []byte(t))
	return len(mm) > 0
}

type Match struct {
	matched []byte
	rest []byte
}

func (rs RuleSet) matchRule(n int, t []byte) []Match {
	//if n == 8 {
	//	return rs.matchRule8(t)
	//}

	//if n == 11 {
	//	return rs.matchRule11(t)
	//}

	rule := rs[n]
	if rule.lit != "" {
		k := len(rule.lit)
		litb := []byte(rule.lit)
		if bytes.Equal(litb, t[:k]) {
			return []Match{{litb, t[k:]}}
		} else {
			return nil
		}
	}

	mm := []Match{}
	for _, branch := range rule.branches {
		mm = append(mm, rs.matchBranch(branch, t))
	}
	return mm
}

//// 42+
//func (rs RuleSet) matchRule8(t []byte) (bool, []byte) {
//	ok, u := rs.matchRule(42, t)
//	if !ok {
//		return false, t
//	}
//
//	for {
//		if len(u) == 0 {
//			return true, u
//		}
//
//		ok, v := rs.matchRule(42, u)
//		if !ok {
//			return true, u
//		}
//		u = v
//	}
//}
//
//func (rs RuleSet) matchRule11(t []byte) (bool, []byte) {
//	ok, u := rs.matchRule(42, t)
//	if !ok {
//		return false, t
//	}
//
//	n := 1
//	for {
//		if len(u) == 0 {
//			break
//		}
//
//		ok, v := rs.matchRule(42, u)
//		if ok {
//			n++
//			u = v
//		} else {
//			break
//		}
//	}
//
//	m := n
//	for m > 0 {
//		if len(u) == 0 {
//			return false, t
//		}
//
//		ok, v := rs.matchRule(31, u)
//		if ok {
//			m--
//			u = v
//		} else {
//			return false, t
//		}
//	}
//	return true, u
//}

func (rs RuleSet) matchBranch(branch []int, t []byte) []Match {
	mm := []Match{{nil, t}}
	for _, r := range branch {
		res := []Match{}
		for _, m := range mm {
			nn := rs.matchRule(r, m.rest)
			for _, n := range nn {
				res = append(res, Match{
					append(m.matched, n.matched...),
					n.rest,
				})
			}
		}
		mm = res
	}
	return mm
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
