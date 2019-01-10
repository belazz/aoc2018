package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

type Claim struct {
	row    int
	col    int
	width  int
	height int
}

func main() {
	contents, err := ioutil.ReadFile("3/input.txt")
	if err != nil {
		fmt.Println("file test.txt not found")
	}

	regex := regexp.MustCompile(`(#(\d+) @ (\d+),(\d+): (\d+)x(\d+))`)
	coords := regex.FindAllSubmatch(contents, -1)

	fabric := [1000][1000]int{}
	claims := make(map[int]Claim)
	for _, fullmatch := range coords {
		id, _ := strconv.Atoi(string(fullmatch[2]))
		i, _ := strconv.Atoi(string(fullmatch[4]))
		j, _ := strconv.Atoi(string(fullmatch[3]))
		width, _ := strconv.Atoi(string(fullmatch[5]))
		height, _ := strconv.Atoi(string(fullmatch[6]))
		claims[id] = Claim{i, j, width, height}

		for row := i; row < height+i; row++ {
			for col := j; col < width+j; col++ {
				fabric[row][col]++
			}
		}
		//fmt.Printf("ID: %v, j: %v, i: %v, width: %v, height: %v\n", id, j, i, width, height)
	}

	var count int
	for _, row := range fabric {
		for _, item := range row {
			if item > 1 {
				count++
			}
		}
	}

	fmt.Printf("Part1: How many square inches of fabric are within two or more claims? Answer: %v\n", count)

	for id, claim := range claims {
		found := true
		for row := claim.row; row < claim.row+claim.height; row++ {
			for col := claim.col; col < claim.col+claim.width; col++ {
				if fabric[row][col] != 1 {
					found = false
					break
				}
			}
		}

		if found {
			fmt.Printf("Part2: What is the ID of the only claim that doesn't overlap? Answer: %v", id)
			break
		}
	}
}
