package main

import (
	"aoc2018/helpers"
	"fmt"
	"strconv"
)

func main() {
	serialNumber, _ := strconv.Atoi(helpers.ReadFileToString("11/input.txt"))
	grid := makeGrid(serialNumber)
	x, y := part1(grid)
	fmt.Printf("part 1 coordinates at %d, %d", x, y)

	fmt.Println()

	x, y, size := part2(grid)
	fmt.Printf("part 2 coordinates at %d, %d with size %d", x, y, size)
}

func part1(grid [][]int) (x, y int) {
	coords, maxArea := []int{0, 0}, 0
	s := 3
	for x := 0; x < 300-s+1; x++ {
		for y := 0; y < 300-s+1; y++ {
			newArea := area(x, y, s, grid)
			if newArea > maxArea {
				maxArea = newArea
				coords = []int{x + 1, y + 1}
			}
		}
	}

	return coords[0], coords[1]
}

func part2(grid [][]int) (x, y, size int) {
	coords, maxArea, maxSize := []int{0, 0}, 0, 0
	for s := 2; s < 300; s++ {
		for x := 0; x < 300-s; x++ {
			for y := 0; y < 300-s; y++ {
				newArea := area(x, y, s, grid)
				if newArea > maxArea {
					maxArea = newArea
					coords = []int{x + 1, y + 1}
					maxSize = s
				}
			}
		}

	}

	return coords[0], coords[1], maxSize
}

func makeGrid(serialNumber int) [][]int {
	grid := make([][]int, 300)
	for x := 0; x < 300; x++ {
		grid[x] = make([]int, 300)
		for y := 0; y < 300; y++ {
			grid[x][y] = powerLevel(serialNumber, x, y)
			grid[x][y] = summedArea(grid, x, y)
		}
	}

	return grid
}

func summedArea(grid [][]int, x int, y int) int {
	sum := grid[x][y]
	if x >= 1 {
		sum += grid[x-1][y]
	}
	if y >= 1 {
		sum += grid[x][y-1]
	}
	if x >= 1 && y >= 1 {
		sum -= grid[x-1][y-1]
	}

	return sum
}

//Find the fuel cell's rack ID, which is its X coordinate plus 10.
//Begin with a power level of the rack ID times the Y coordinate.
//Increase the power level by the value of the grid serial number (your puzzle input).
//Set the power level to itself multiplied by the rack ID.
//Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
//Subtract 5 from the power level.

func powerLevel(serialNumber int, x int, y int) int {
	rackId := x + 1 + 10
	powerLevel := rackId * (y + 1)
	powerLevel += serialNumber
	powerLevel = powerLevel * rackId
	powerLevel = (powerLevel / 100) % 10

	return powerLevel - 5
}

func area(x, y int, size int, grid [][]int) int {
	area, x0, y0, x1, y1 := 0, x-1, y-1, x+size-1, y+size-1
	if x0 >= 0 && y0 >= 0 {
		area += grid[x0][y0]
	}
	if x0 >= 0 {
		area -= grid[x0][y1]
	}
	if y0 >= 0 {
		area -= grid[x1][y0]
	}
	area += grid[x1][y1]

	return area
}
