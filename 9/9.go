package main

import (
	"aoc2018/datastruct"
	"aoc2018/helpers"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	contents := helpers.ReadFileToString("9/test.txt")
	fmt.Println(contents)

	list := datastruct.LinkedList{Head: nil}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		list = list.Append(&datastruct.Node{Value: rand.Int() % 10000})
	}
	fmt.Println('d')
}
