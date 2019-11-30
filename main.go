package main

import "fmt"

type Graph struct {
	AdjacencyList []Node
}

type Node struct {
	Value int
	Links []int
	Color int
}

func (graph *Graph) addNode(value int, links []int) {
	node := Node{value, links, -1}
	graph.AdjacencyList = append(graph.AdjacencyList, node)
}

func (graph *Graph) display() {
	for _, node := range graph.AdjacencyList {
		fmt.Println(node)
	}
}

func (graph *Graph) findNode(value int) Node {
	var nodeFound Node

	for _, node := range graph.AdjacencyList {
		if node.Value == value {
			nodeFound = node
		}
	}

	return nodeFound
}

func (graph *Graph) findMinNode() Node {
	var minNode Node

	for _, node := range graph.AdjacencyList {
		if node.Value < minNode.Value {
			minNode = node
		}
	}

	return minNode
}

func (graph *Graph) generateColors() {
	verticesCount := len(graph.AdjacencyList)
	graph.AdjacencyList[0].Color = 0

	for index := 1; index < verticesCount; index++ {
		availableColors := initColorMap(verticesCount)
		for _, link := range graph.AdjacencyList[index].Links {
			linkNode := graph.findNode(link)
			if linkNode.Color != -1 {
				availableColors[linkNode.Color] = false
			}
		}
		graph.AdjacencyList[index].Color = getFirstAvailableColor(availableColors)
	}
}

func initColorMap(count int) map[int]bool {
	colorMap := make(map[int]bool)

	for index := 0; index < count; index++ {
		colorMap[index] = true
	}

	return colorMap
}

func getFirstAvailableColor(unavailableColors map[int]bool) int {
	var availableColor int

	for index := 0; index < len(unavailableColors); index++ {
		if unavailableColors[index] == true {
			availableColor = index
			break
		}
	}

	return availableColor
}

func main() {
}
