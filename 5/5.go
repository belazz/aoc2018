package main

import (
	"2018/helpers"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	contents := helpers.ReadFileToString("5/input.txt")
	part1 := contents
	result := collapseScan(part1)
	fmt.Printf("Part1 answer: %v\n", len(result))
	minPolymerLength := len(contents)
	for _, item := range strings.Split("A B C D E F G H I K L M N O P Q R S T V X Y Z", " ") {
		part2 := contents
		withoutUnit := removeUnit(item, part2)
		fmt.Printf("Removing unit %v...\n", item)
		collapsedLength := len(collapseScan(withoutUnit))
		if minPolymerLength > collapsedLength {
			minPolymerLength = collapsedLength
		}
	}
	fmt.Printf("Part2 answer: %v\n", minPolymerLength)
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

func collapseScan(polymer string) string {
	splitPolymer := strings.Split(polymer, "")
	polymerLength := len(splitPolymer)
	for i := 0; i < polymerLength-1; i++ {
		if strings.ToLower(splitPolymer[i]) == strings.ToLower(splitPolymer[i+1]) && splitPolymer[i] != splitPolymer[i+1] {
			splitPolymer = append(splitPolymer[:i], splitPolymer[i+2:]...)
			i = 0
			polymerLength -= 2
		}
	}

	return strings.Join(splitPolymer, "")
}

func removeUnit(unit string, polymer string) string {
	remover := regexp.MustCompile(fmt.Sprintf("((?i)[%v])", unit))
	return remover.ReplaceAllLiteralString(polymer, "")
}
