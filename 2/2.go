package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("-----------------PART 1--------------");
	filename := "2/2.txt"
	width := 0
	height := 0
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("coulnt read file" + filename)
	}
	reg, err := regexp.Compile(`\b([a-zA-Z]+)`)
	if err != nil {
		fmt.Println(err)
	}
	items := reg.FindAllString(string(data), -1)

	for _, str := range items {
		count := make(map[string]int)
		for _, char := range strings.Split(string(str), "") {
			count[char]++
		}
		twidth, theight := findExact(count)
		width += twidth
		height += theight
	}

	fmt.Printf("width %d, height %d\n", width, height)
	fmt.Printf("part1 answer: %d\n", width*height)
	fmt.Println("-----------------PART 2--------------");

	for i := 0; i < len(items); i++ {
		wrongCount, wrongIndex := 0, 0
		for j := i + 1; j < len(items); j++ {
			wrongCount, wrongIndex = 0, 0
			charsi, charsj := strings.Split(items[i], ""), strings.Split(items[j], "")
			for k := 0; k < len(charsi); k++ {
				if charsi[k] != charsj[k] {
					wrongCount++
					wrongIndex = k
					if wrongCount > 1 {
						break
					}
				}
			}
			if wrongCount == 1 {
				fmt.Println("found the matching strings!")
				fmt.Printf("one of the strings: %s, differ char index: %d\n", items[i], wrongIndex)
				matchingSymbols := strings.Split(items[i], "")
				matchingSymbols = append(matchingSymbols[:wrongIndex], matchingSymbols[wrongIndex+1:]..., )
				fmt.Printf("part2 answer: %s", strings.Join(matchingSymbols, ""))
				break
			}
		}
	}
}

func findExact(chars map[string]int) (int, int) {
	twice := 0
	thrice := 0
	for _, value := range chars {
		if value == 2 && twice == 0{
			twice++
		}
		if value == 3 && thrice == 0 {
			thrice++
		}
	}

	return twice, thrice
}
