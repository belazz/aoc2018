package main

import (
	"awesomeProject/helpers"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	contents := helpers.ReadFileToString("5/input.txt")
	result := collapse(contents)
	fmt.Println(len(result))
	fmt.Println(result)
}

func collapse(polymer string) string {
	oldPolymer := polymer
	twiceRegex := regexp.MustCompile(`([a-z][A-Z])`)
	twiceRegexAlt := regexp.MustCompile(`([A-Z][a-z])`)

	for {
		matches := twiceRegex.FindAllString(polymer, -1)
		matches = append(matches, twiceRegexAlt.FindAllString(polymer, -1)...)
		oldPolymer = polymer
		for _, match := range matches {
			chars := strings.Split(match, "")
			first, second := chars[0], chars[1]
			if strings.ToLower(first) == strings.ToLower(second) && first != second {
				polymer = strings.Replace(polymer, match, "", 1)
			}
		}
		if oldPolymer == polymer {
			return polymer
		}
	}
}
