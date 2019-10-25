package main

import (
	"aoc2018/helpers"
	"container/ring"
	"fmt"
	"regexp"
	"strconv"
	//"container/ring"
)

func main() {
	contents := helpers.ReadFileToString("9/test.txt")

	regex := regexp.MustCompile(`(\d+)\D+(\d+)`)
	playersAndNumbers := regex.FindAllStringSubmatch(contents, -1)

	playersCount, _ := strconv.Atoi(playersAndNumbers[0][1])
	marblesCount, _ := strconv.Atoi(playersAndNumbers[0][2])

	players := make([]int, playersCount)
	currentPlayer := 0

	circle := ring.New(1)
	circle = circle.Next()
	circle.Value = 0

	for i := 1; i <= marblesCount; i++ {
		newLink := ring.Ring{Value: i}

		if i%23 == 0 {
			players[currentPlayer] += i
			for j := 0; j < 7; j++ {
				circle = circle.Prev()
			}
			players[currentPlayer] += circle.Value.(int)
			circle = circle.Prev()
			circle.Unlink(1)
			circle = circle.Next()
		} else {
			circle = circle.Next()
			circle = circle.Link(&newLink).Prev()
		}

		currentPlayer++
		if currentPlayer >= playersCount {
			currentPlayer = 0
		}
	}

	max := 0
	for _, item := range players {
		//fmt.Println(item)
		if max < item {
			max = item
		}
	}
	fmt.Println(max)
}
