package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strconv"
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

	nn := []int64{}
	aa := []int64{}

	line2 := <-ch
	for i, s := range strings.Split(line2, ",") {
		if s == "x" {
			continue
		}

		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		nn = append(nn, n)
		aa = append(aa, n - int64(i))
	}

	if len(nn) < 2 {
		panic("Unexpectedly few buses")
	}

	n1 := nn[0]
	a1 := aa[0]
	var t int64
	for i := 1; i < len(nn); i++ {
		n2 := nn[i]
		a2 := aa[i]

		m1, m2 := BézoutCoeffs(n1, n2)
		t = a1*m2*n2 + a2*m1*n1
		for t < 0 {
			t += n1*n2
		}
		for t >= n1*n2 {
			t -= n1*n2
		}

		fmt.Printf("n1=%v a1=%v n2=%v a2=%v m1=%v m2=%v t=%v\n",
		n1, a1, n2, a2, m1, m2, t)
		n1 *= n2
		a1 = t
	}
	fmt.Println(t)
}

// From https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm
func BézoutCoeffs(a, b int64) (int64, int64) {
	rOld, r := a, b
	sOld, s := int64(1), int64(0)
	tOld, t := int64(0), int64(1)

	for r != 0 {
		q := rOld / r
		rOld, r = r, rOld-q*r
		sOld, s = s, sOld-q*s
		tOld, t = t, tOld-q*t
	}

	return sOld, tOld
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
