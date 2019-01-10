package main

import (
	"2018/helpers"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type Distance struct {
	id              string
	distanceToPoint int
	point           Point
}

type Point struct {
	x int
	y int
}

func main() {
	contents := helpers.ReadFileToString("6/input.txt")
	bounds := 1000
	shift := 400
	offset := 5
	points := make(map[string]int)
	regionCount := 0

	digitsReg := regexp.MustCompile(`(\d+), (\d+)`)
	digitsStr := digitsReg.FindAllStringSubmatch(contents, -1)
	for x := 0; x < bounds; x++ {
		for y := 0; y < bounds; y++ {
			distanceToPoints := []Distance{}
			totalDistance := 0

			for _, fullmatch := range digitsStr {
				xCoord, _ := strconv.Atoi(fullmatch[1])
				yCoord, _ := strconv.Atoi(fullmatch[2])
				pointId := fullmatch[1] + fullmatch[2]
				distanceToCurrentPoint := helpers.Taxicab(x, y, xCoord+shift, yCoord+shift)
				totalDistance += distanceToCurrentPoint
				distanceToPoints = append(distanceToPoints, Distance{pointId, distanceToCurrentPoint, Point{x, y}})
			}

			//part 1
			sort.Slice(distanceToPoints, func(i, j int) bool {
				return distanceToPoints[i].distanceToPoint < distanceToPoints[j].distanceToPoint
			})
			if distanceToPoints[0].distanceToPoint != distanceToPoints[1].distanceToPoint {
				points[distanceToPoints[0].id]++
			}
			//part 2
			if totalDistance < 10000 {
				regionCount++
			}
		}
	}
	currentMax := 0
	for id, item := range points {
		if item > bounds*offset {
			continue
		} else {
			if currentMax < item {
				currentMax = item
			}
		}

		fmt.Printf("Point, id:%v, count: %v\n", id, item)
	}

	fmt.Printf("Part1 answer: %v\n", currentMax)
	fmt.Printf("Part2 answer: %v", regionCount)
}
