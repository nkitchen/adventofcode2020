package main

import "fmt"
import "io/ioutil"
import "log"
import "os"
import "regexp"
import "strings"
import "strconv"

func main() {
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(inFile)
	if err != nil {
		log.Fatal(err)
	}

	inText := string(buf)

	fieldValidator := map[string]func(string)bool{
		"byr": yearInRange(1920, 2002),
		"iyr": yearInRange(2010, 2020), 
		"eyr": yearInRange(2020, 2030),
		"hgt": checkHeight,
		"hcl": checkHairColor,
		"ecl": checkEyeColor,
		"pid": checkPassportID,
		//"cid": checkCountryID,
	}

	validPassports := 0
	passports := strings.Split(inText, "\n\n")
	for _, pp := range passports {
		fields := strings.Fields(pp)
		goodFields := 0
		for _, f := range fields {
			s := strings.Split(f, ":")
			label := s[0]
			v := s[1]
			validate := fieldValidator[label]
			if validate == nil {
				continue
			}
			if validate(v) {
				goodFields++
			}
		}
		if goodFields == len(fieldValidator) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

func yearInRange(min, max int) (func(string) bool) {
	return func(s string) bool {
		re := regexp.MustCompile(`^\d{4}$`)
		m := re.FindString(s)
		if m == "" {
			return false
		}
		y, _ := strconv.Atoi(m)
		return min <= y && y <= max
	}
}

func checkHeight(s string) bool {
	re := regexp.MustCompile(`^(\d+)(cm|in)$`)
	m := re.FindStringSubmatch(s)
	if len(m) == 0 {
		return false
	}
	h, _ := strconv.Atoi(m[1])
	switch m[2] {
	case "cm":
		return 150 <= h && h <= 193
	case "in":
		return 59 <= h && h <= 76
	default:
		panic("Unexpected unit: " + m[2])
	}
	return false
}

func checkHairColor(s string) bool {
	re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	m := re.FindString(s)
	return m != ""
}

func checkEyeColor(s string) bool {
	switch s {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func checkPassportID(s string) bool {
	re := regexp.MustCompile(`^\d{9}$`)
	m := re.FindString(s)
	return m != ""
}
