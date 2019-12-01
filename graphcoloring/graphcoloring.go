package graphcoloring

import "fmt"

type Graph struct {
	AdjacencyList []Node
}

type Node struct {
	Value int
	Links []int
	Color int
}

func (graph *Graph) Display() {
	for _, node := range graph.AdjacencyList {
		fmt.Println(node)
	}
}

func (graph *Graph) AddNode(value int, links []int) {
	node := Node{value, links, -1}
	graph.AdjacencyList = append(graph.AdjacencyList, node)
}

func (graph *Graph) FindNode(value int) Node {
	var nodeFound Node

	for _, node := range graph.AdjacencyList {
		if node.Value == value {
			nodeFound = node
		}
	}

	return nodeFound
}

func (graph *Graph) FindMinNode() Node {
	var minNode Node

	if len(graph.AdjacencyList) > 0 {
		minNode = graph.AdjacencyList[0]

		for index := 1; index < len(graph.AdjacencyList); index++ {
			node := graph.AdjacencyList[index]
			if node.Value < minNode.Value {
				minNode = node
			}
		}
	}

	return minNode
}

func (graph *Graph) GenerateColors() {
	verticesCount := len(graph.AdjacencyList)
	if verticesCount > 0 {
		graph.AdjacencyList[0].Color = 0
		availableColorList := InitColorMap(verticesCount)

		for index := 1; index < verticesCount; index++ {
			availableColors := availableColorList
			for _, link := range graph.AdjacencyList[index].Links {
				linkNode := graph.FindNode(link)
				if linkNode.Color != -1 {
					availableColors[linkNode.Color] = false
				}
			}
			graph.AdjacencyList[index].Color = GetFirstAvailableColor(availableColors)
		}
	}
}

func InitColorMap(count int) map[int]bool {
	colorMap := make(map[int]bool)

	for index := 0; index < count; index++ {
		colorMap[index] = true
	}

	return colorMap
}

func GetFirstAvailableColor(unavailableColors map[int]bool) int {
	var availableColor int

	for index := 0; index < len(unavailableColors); index++ {
		if unavailableColors[index] == true {
			availableColor = index
			break
		}
	}

	return availableColor
}
