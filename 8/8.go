package main

import (
	"awesomeProject/helpers"
	"fmt"
	"regexp"
	"strconv"
)

type Tree struct {
	children []*Tree
	metadata []int
	value    int
}

func main() {
	allNumbersRe := regexp.MustCompile(`(\d+)`)
	contents := helpers.ReadFileToString("8/input.txt")
	numbersStr := allNumbersRe.FindAllString(contents, -1)
	numbers := []int{}
	for _, item := range numbersStr {
		number, _ := strconv.Atoi(item)
		numbers = append(numbers, number)
	}
	index := 0
	root := &Tree{}
	// parse tree
	index, root = traverse(index, numbers, root)
	// sum metadata
	metadataSum := sumMetadata(root, 0)
	fmt.Printf("Part1 answer: %v\n", metadataSum)

	sumValues(root, 0)
	fmt.Printf("Part2 answer: %v\n", root.value)
}

func traverse(index int, numbers []int, root *Tree) (int, *Tree) {
	// if children count >= 1
	childrenCount := numbers[index]
	metadataCount := numbers[index+1]
	// shift index by 2, first 2 numbers contain children count and metadata entries count
	index += 2
	// iterate over children
	for i := 0; i < childrenCount; i++ {
		// create children along the way
		child := &Tree{}
		root.children = append(root.children, child)
		index, _ = traverse(index, numbers, child)
	}

	// collect metadata after traversing the children
	i := index
	for ; i < metadataCount+index; i++ {
		root.metadata = append(root.metadata, numbers[i])
	}
	return i, root
}

func sumMetadata(root *Tree, metadataSum int) int {
	for _, child := range root.children {
		metadataSum = sumMetadata(child, metadataSum)
	}

	for _, entry := range root.metadata {
		metadataSum += entry
	}

	return metadataSum
}

func sumValues(root *Tree, valuesSum int) int {
	for _, child := range root.children {
		root.value = sumValues(child, valuesSum)
	}

	if root.children != nil {
		for _, entry := range root.metadata {
			if entry-1 < len(root.children) {
				root.value += root.children[entry-1].value
			}
		}
	} else {
		for _, entry := range root.metadata {
			root.value += entry
		}
	}

	return valuesSum
}
