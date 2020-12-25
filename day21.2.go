package main

import "bufio"
import "fmt"
import "log"
import "os"
import "regexp"
import "sort"
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
		for _, allerg := range allergens {
			foodsByAllergen[allerg] = append(foodsByAllergen[allerg], food)
		}
	}

	ingredientsByAllergen := map[string]*stringset.Set{}
	badIngredients := stringset.New()
	for allerg := range foodsByAllergen {
		var ings stringset.Set
		for i, food := range foodsByAllergen[allerg] {
			if i == 0 {
				ings = ingredientsByFood[food]
			} else {
				ings = ings.Intersect(ingredientsByFood[food])
			}
		}
		ingredientsByAllergen[allerg] = &ings
		badIngredients.Update(ings)
	}

	type danger struct {
		allergen string
		ingredient string
	}
	dangers := []danger{}
	for len(ingredientsByAllergen) > 0 {
		for allerg := range ingredientsByAllergen {
			if ingredientsByAllergen[allerg].Len() == 1 {
				elts := ingredientsByAllergen[allerg].Elements()
				ingred := elts[0]
				dangers = append(dangers, danger{allerg, ingred})
				delete(ingredientsByAllergen, allerg)
				for _, s := range ingredientsByAllergen {
					s.Discard(ingred)
				}
				break
			}
		}
	}

	sort.Slice(dangers, func(i, j int) bool {
		return dangers[i].allergen < dangers[j].allergen
	})

	badList := []string{}
	for _, d := range dangers {
		badList = append(badList, d.ingredient)
	}
	fmt.Println(strings.Join(badList, ","))
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
