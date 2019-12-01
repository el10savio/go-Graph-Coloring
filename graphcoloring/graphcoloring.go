package graphcoloring

import "fmt"

// Graph depicts the graph structure. The graph
// relationship is represented using an
// adjacency list that describes each
// node with links to its neighbors
type Graph struct {
	AdjacencyList []Node
}

// Node depicts each individual node in the graph
// The value here is considered as its index and
// is unique. Each node also describes the
// nodes it is linked to and its color
type Node struct {
	Value int
	Links []int
	Color int
}

// Display prints out each node
// in the graph's adjacency list
func (graph *Graph) Display() {
	for _, node := range graph.AdjacencyList {
		fmt.Println(node)
	}
}

// AddNode takes the value and links to
// create a Node and update the
// graph's adjacency list
func (graph *Graph) AddNode(value int, links []int) {
	node := Node{value, links, -1}
	graph.AdjacencyList = append(graph.AdjacencyList, node)
}

// FindNode takes a value and returns
// a node in the graph which
// contains that value
func (graph *Graph) FindNode(value int) Node {
	var nodeFound Node

	for _, node := range graph.AdjacencyList {
		if node.Value == value {
			nodeFound = node
		}
	}

	return nodeFound
}

// FindSmallestNode returns a node
// in the graph which has
// the smallest value
func (graph *Graph) FindSmallestNode() Node {
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

// GenerateColors is the greedy graph coloring implementation.
// It first assigns the first node with the first available
// color and then for each corresponding node creates a
// hash map of a list of available colors and marks
// the colors as false if its neighbors contains
// that color. The maximum number of colors
// the graph could have is the number of
// nodes it contains
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

// InitColorMap returns a hash map of the
// list of possible colors and whether
// the color is available or not
func InitColorMap(count int) map[int]bool {
	colorMap := make(map[int]bool)

	for index := 0; index < count; index++ {
		colorMap[index] = true
	}

	return colorMap
}

// GetFirstAvailableColor iterates through the hash map
// of the list of possible colors and whether
// returns the first available color it finds
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
