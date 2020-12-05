package main

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "strings"

func main() {
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(inFile)
	if err != nil {
		log.Fatal(err)
	}

	expectedFields := map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
		//"cid": true,
	}

	inText := string(buf)

	nValid := 0
	passports := strings.Split(inText, "\n\n")
	for _, pp := range passports {
		fields := strings.Fields(pp)
		fieldLabels := map[string]bool{}
		for _, f := range fields {
			s := strings.Split(f, ":")
			fieldLabels[s[0]] = true
		}

		nPresent := 0
		for f := range expectedFields {
			if fieldLabels[f] {
				nPresent++
			}
		}

		nExpected := len(expectedFields)
		if fieldLabels["cid"] {
			nExpected++
			nPresent++
		}

		if nPresent == nExpected && nPresent == len(fieldLabels) {
			nValid++
		}
	}

	fmt.Println(nValid)
}
