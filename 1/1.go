package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	input, err := ioutil.ReadFile("1/test.txt")
	if err != nil {
		log.Fatal("file test.txt not found")
	}
	count := 0
	var isFound = false
	var processed = []int{0}
	for !isFound {
		for _, element := range bytes.Split(input, []byte("\r\n")) {
			if intElem, err := strconv.Atoi(string(element)); err == nil {
				count += intElem
				if contains(processed, count) {
					fmt.Println(count)
					isFound = true
					break
				}
				processed = append(processed, count)
			} else {
				fmt.Println(err)
			}
		}
	}
}

func contains(haystack []int, needle int) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}

	return false
}
