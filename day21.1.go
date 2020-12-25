package main

import "bufio"
import "fmt"
import "log"
import "os"
import "regexp"
import "strings"

import "bitbucket.org/creachadair/stringset"

func main() {
	ingredientsByFood := []stringset.Set{}
	foodsByAllergen := map[string][]int{}
	re := regexp.MustCompile(`(\S+(?: \S+)*) [(]contains (\S+(?:, \S+)*)[)]`)
	for line := range inputLines() {
		m := re.FindStringSubmatch(line)
		if len(m) == 0 {
			log.Fatalf("No match for line: %s", line)
		}

		food := len(ingredientsByFood)

		ings := stringset.New(strings.Fields(m[1])...)
		ingredientsByFood = append(ingredientsByFood, ings)

		allergens := strings.Split(m[2], ", ")
		for _, all := range allergens {
			foodsByAllergen[all] = append(foodsByAllergen[all], food)
		}
	}

	allergenIngredients := stringset.New()
	for all := range foodsByAllergen {
		var ings stringset.Set
		for i, food := range foodsByAllergen[all] {
			if i == 0 {
				ings = ingredientsByFood[food]
			} else {
				ings = ings.Intersect(ingredientsByFood[food])
			}
		}
		allergenIngredients.Update(ings)
	}

	occs := 0
	for _, ings := range ingredientsByFood {
		occs += ings.Diff(allergenIngredients).Len()
	}
	fmt.Println(occs)
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
