package main

import "bufio"
import "bytes"
import "fmt"
import "log"
import "os"
import "strconv"
import "strings"

type Rule struct {
	lit []byte
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
				rule.lit = []byte(body[1:k-1])
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
	for _, m := range mm {
		if len(m.rest) == 0 {
			return true
		}
	}
	return false
}

type Match struct {
	matched []byte
	rest []byte
}

func (rs RuleSet) matchRule(n int, t []byte) []Match {
	if n == 8 {
		return rs.matchRule8(t)
	}

	if n == 11 {
		return rs.matchRule11(t)
	}

	rule := rs[n]
	if len(rule.lit) != 0 {
		if bytes.HasPrefix(t, rule.lit) {
			k := len(rule.lit)
			return []Match{{rule.lit, t[k:]}}
		} else {
			return nil
		}
	}

	mm := []Match{}
	for _, branch := range rule.branches {
		mm = append(mm, rs.matchBranch(branch, t)...)
	}
	return mm
}

// 42+
func (rs RuleSet) matchRule8(t []byte) []Match {
	q := rs.matchRule(42, t)
	res := []Match{}
	for len(q) > 0 {
		m := q[0]
		res = append(res, m)
		q = q[1:]

		nn := rs.matchRule(42, m.rest)
		for _, n := range nn {
			q = append(q, Match{
				append(m.matched, n.matched...),
				n.rest,
			})
		}
	}
	return res
}

// 42^k 31^k
func (rs RuleSet) matchRule11(t []byte) []Match {
	// mm42[k]: matches of rule 42, k times in a row
	mm42 := make([][]Match, 1)
	mm42[0] = []Match{{nil, t}}
	for k := 1; len(mm42[k-1]) > 0; k++ {
		mm42 = append(mm42, nil)
		for _, m := range mm42[k-1] {
			nn := rs.matchRule(42, m.rest)
			for _, n := range nn {
				r := Match{
					append(m.matched, n.matched...),
					n.rest,
				}
				mm42[k] = append(mm42[k], r)
			}
		}
	}

	res := []Match{}
	for k := range mm42 {
		if k == 0 {
			continue
		}
		q := mm42[k]
		for j := 1; j <= k; j++ {
			rr := []Match{}
			for _, m := range q {
				nn := rs.matchRule(31, m.rest)
				for _, n := range nn {
					r := Match{
						append(m.matched, n.matched...),
						n.rest,
					}
					rr = append(rr, r)
				}
			}
			// Each of rr is now 42^k 31^j
			q = rr
		}
		res = append(res, q...)
	}
	return res
}

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
