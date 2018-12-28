package helpers

import (
	"fmt"
	"io/ioutil"
)

func ReadFileToString(filepath string) string {
	contentsBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Specified file not found")
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

