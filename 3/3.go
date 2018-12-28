package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	contents, err := ioutil.ReadFile("3/test.txt")
	if err != nil {
		fmt.Println("file test.txt not found")
	}

	regex := regexp.MustCompile(`(#(\d+) @ (\d+),(\d+): (\d+)x(\d+))`)
	coords := regex.FindAllSubmatch(contents, -1)

	fabric := [1000][1000]int{}

	for _, fullmatch := range coords {
		id, _ := strconv.Atoi(string(fullmatch[2]))
		i, _ := strconv.Atoi(string(fullmatch[4]))
		j, _ := strconv.Atoi(string(fullmatch[3]))
		width, _ := strconv.Atoi(string(fullmatch[5]))
		height, _ := strconv.Atoi(string(fullmatch[6]))
		//fmt.Printf("ID: %v, j: %v, i: %v, width: %v, height: %v\n", id, j, i, width, height)
		fabric[i][j]++
	}

	var count int
	for _, row := range fabric {
		for _, item := range row {
			if item > 0 {
				count++
			}
		}
	}

	fmt.Printf("How many square inches of fabric are within two or more claims? Answer: %v", count)
}
