package helpers

import (
	"fmt"
	"io/ioutil"
)

func ReadFileToString(filepath string) string {
	contentsBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Specified file was not found")
	}

	return string(contentsBytes)
}

func GenerateAsciiLower() []rune {
	asciiLower := []rune{}

	for i := 'a'; i < 'z'; i++ {
		asciiLower = append(asciiLower, i)
	}

	return asciiLower
}

func GenerateAsciiUpper() []rune {
	asciiLower := []rune{}

	for i := 'A'; i < 'Z'; i++ {
		asciiLower = append(asciiLower, i)
	}

	return asciiLower
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func Taxicab(x1, y1, x2, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

func ContainsInt(needle int, haystack []int) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}

	return false
}

func ContainsIntInMap(needle int, haystack map[string]int) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}

	return false
}

func IntMin(x int, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func IntMax(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
