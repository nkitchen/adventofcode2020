package main

import "bufio"
import "fmt"
import "log"
import "math/big"
import "os"
import "strings"

// n_1, n_2, ..., n_k
// a_1 = 0, a_2, ..., a_k
// min t :
//     t = a_1 (mod n_1)
//     t = a_2 (mod n_2)
//     ...
//
// Chinese remainder theorem:
//     ∃ unique t : 0 <= t < n_1 * ... * n_k = N
//
// Bézout's identity:
//     m_1 * n_1 + m_2 * n_2 = 1
//     t = a_1 * m_2 * n_2 + a_2 * m_1 * n_1 (mod n_1 * n_2)
//
//     m_1,2 * (n_1 * n_2) + m_3 * n_3 = 1
//     t = a_1,2 * m_3 * n_3 + a_3 * m_1,2 * (n_1 * n_2) (mod n_1 * n_2 * n_3)
//
//     ...

func main() {
	ch := inputLines()
	line1 := <-ch
	_ = line1

	nn := []*big.Int{}
	aa := []*big.Int{}

	line2 := <-ch
	for i, s := range strings.Split(line2, ",") {
		if s == "x" {
			continue
		}

		n := &big.Int{}
		_, ok := n.SetString(s, 10)
		if !ok {
			log.Fatalf("Parse error for %q", s)
		}
		a := big.NewInt(-int64(i))
		a.Mod(a, n)
		nn = append(nn, n)
		aa = append(aa, a)
	}

	if len(nn) < 2 {
		panic("Unexpectedly few buses")
	}

	n1 := nn[0]
	a1 := aa[0]
	t := &big.Int{}
	for i := 1; i < len(nn); i++ {
		n2 := nn[i]
		a2 := aa[i]

		//fmt.Printf("n1=%v\ta1=%v\tn2=%v\ta2=%v", n1, a1, n2, a2)

		z := &big.Int{}
		m1 := &big.Int{}
		m2 := &big.Int{}
		z.GCD(m1, m2, n1, n2)
		if z.String() != "1" {
			panic(z.String())
		}

		mn2a1 := &big.Int{}
		mn1a2 := &big.Int{}
		mn2a1.Mul(m2, n2)
		mn2a1.Mul(mn2a1, a1)
		mn1a2.Mul(m1, n1)
		mn1a2.Mul(mn1a2, a2)
		t.Add(mn2a1, mn1a2)

		n1.Mul(n1, n2)
		t.Mod(t, n1)
		a1 = t
		//fmt.Printf("\tm1=%v\tm2=%v\tt=%v\n", m1, m2, t)
	}
	fmt.Println(t)
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
