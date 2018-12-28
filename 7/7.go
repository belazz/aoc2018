package main

import (
	"awesomeProject/helpers"
	"fmt"
	"regexp"
	"sort"
)

type Worker struct {
	isAvailable bool
	currentStep string
	workTill int
}

func main() {
	// part 1
	correctOrder := ""
	nodes, noEdgeNodes := populate()
	for len(noEdgeNodes) != 0 {
		poppedNode := noEdgeNodes[0]
		correctOrder += poppedNode

		// remove from map: node -> [incoming edges]
		delete(nodes, poppedNode)
		// finally, remove edge as incoming to nodes
		nodes, noEdgeNodes = removeEdgeAndCollectNoEdgeNodes(poppedNode, &nodes)
	}

	fmt.Printf("Part1 answer: %v", correctOrder)

	// part 2
	allSteps := helpers.GenerateAsciiUpper()
	nodes, noEdgeNodes = populate()
	seconds := 0
	workersAmount := 5
	workers := make(map[int]Worker)
	for ;;seconds++ {
		// check if any workers completed their tasks
		for _, worker := range workers {
			if worker.workTill == seconds {
				worker.isAvailable = true
				// remove node as incoming edge to any other node if worker completed their task
				// collect new no-edge nodes
				nodes, noEdgeNodes = removeEdgeAndCollectNoEdgeNodes(worker.currentStep, &nodes)
			}
		}
		for index, node := range noEdgeNodes {
			workerId := getAvailableWorker(workers)
			if workerId != -1 {
				// put worker to work with a workTill second which shows when the currentStep completes
				workers[workerId] = Worker{isAvailable: false, currentStep: node, workTill: int([]rune(node)[0]) - 64 + 60}
				// remove node from slice
				noEdgeNodes = append(noEdgeNodes[:0], noEdgeNodes[index+1:]...)
			}
		}
	}
}

func getAvailableWorker(workers map[int]Worker) int {
	for key, item := range workers {
		if item.isAvailable {
			return key
		}
	}

	return -1
}

func populate() (map[string][]string, []string) {
	contents := helpers.ReadFileToString("7/input.txt")
	nodesRe := regexp.MustCompile(`Step ([A-Z]) must be finished before step ([A-Z]) can begin`)
	nodesStr := nodesRe.FindAllStringSubmatch(contents, -1)
	nodes := make(map[string][]string)
	noEdgeNodes := []string{}

	for _, fullmatch := range nodesStr {
		edge := fullmatch[1]
		node := fullmatch[2]
		if item, ok := nodes[node]; !ok {
			nodes[node] = []string{edge}
		} else {
			nodes[node] = append(item, edge)
		}
		if _, ok := nodes[edge]; !ok {
			nodes[edge] = []string{}
		}
	}
	noEdgeNodes = collectNoEdgeNodes(nodes)
	sort.Strings(noEdgeNodes)

	return nodes, noEdgeNodes
}

func collectNoEdgeNodes(nodes map[string][]string) []string {
	ret := []string{}
	for node, edges := range nodes {
		if len(edges) == 0 {
			ret = append(ret, node)
		}
	}

	sort.Strings(ret)
	return ret
}

func removeEdgeAndCollectNoEdgeNodes(edge string, nodes *map[string][]string) (map[string][]string, []string) {
	noEdgeNodes := []string{}
	for node, edges := range *nodes {
		(*nodes)[node] = removeEdge(edge, edges)

		if len((*nodes)[node]) == 0 {
			noEdgeNodes = append(noEdgeNodes, node)
		}
	}
	sort.Strings(noEdgeNodes)
	return *nodes, noEdgeNodes
}

func removeEdge(edge string, edges []string) []string {
	newEdges := []string{}

	for index, item := range edges {
		if item == edge {
			newEdges = append(newEdges, edges[:index]...)
			newEdges = append(newEdges, edges[index+1:]...)

			return newEdges
		}
	}

	return edges
}