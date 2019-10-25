package main

import (
	"aoc2018/helpers"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	contents := helpers.ReadFileToString("10/input.txt")
	regex := regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>`)
	matches := regex.FindAllStringSubmatch(contents, -1)
	points := parseIntoMap(matches)
	a := area(bounds(points))
	for t := 1; ; t++ {
		currentBounds := bounds(state(points, t))
		currentArea := area(currentBounds)
		if a > currentArea {
			a = currentArea
		} else {
			display(state(points, t-1))
			fmt.Printf("min area at t: %d\n", t-1)
			break
		}
	}
}

func display(points []map[string][]int) {
	boundingBox := bounds(points)
	x0, y0, x1, y1 := boundingBox[0], boundingBox[1], boundingBox[2], boundingBox[3]
	for ; y0 <= y1; y0++ {
		for j := x0; j <= x1; j++ {
			if isPoint(points, j, y0) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func isPoint(points []map[string][]int, x int, y int) bool {
	for _, point := range points {
		if (point["position"][0] == x) && (point["position"][1] == y) {
			return true
		}
	}

	return false
}

func bounds(points []map[string][]int) []int {
	x0, y0, x1, y1 :=
		points[0]["position"][0], points[0]["position"][1], points[0]["position"][0], points[0]["position"][1]
	for _, point := range points {
		x0 = helpers.IntMin(x0, point["position"][0])
		x1 = helpers.IntMax(x1, point["position"][0])

		y0 = helpers.IntMin(y0, point["position"][1])
		y1 = helpers.IntMax(y1, point["position"][1])
	}

	return []int{x0, y0, x1, y1}
}

func area(bounds []int) int {
	return helpers.Abs((bounds[0] - bounds[2]) * (bounds[1] - bounds[3]))
}

func state(points []map[string][]int, t int) []map[string][]int {
	var pointsAtT []map[string][]int
	for _, point := range points {
		newPoint := make(map[string][]int)
		newPoint["position"] = []int{point["position"][0] + (point["velocity"][0] * t), point["position"][1] + (point["velocity"][1] * t)}
		pointsAtT = append(pointsAtT, newPoint)
	}

	return pointsAtT
}

func parseIntoMap(matches [][]string) []map[string][]int {
	var points []map[string][]int
	for _, match := range matches {
		point := make(map[string][]int)
		var ints []int
		for _, submatch := range match {
			i, _ := strconv.Atoi(submatch)
			ints = append(ints, i)
		}
		point["position"] = []int{ints[1], ints[2]}
		point["velocity"] = []int{ints[3], ints[4]}
		points = append(points, point)
	}

	return points
}
