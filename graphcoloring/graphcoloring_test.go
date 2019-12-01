package graphcoloring

import (
	"reflect"
	"testing"
)

var graphTest = Graph{
	AdjacencyList: []Node{
		Node{Value: 1, Links: []int{2, 3}, Color: -1},
		Node{Value: 2, Links: []int{1, 3}, Color: -1},
		Node{Value: 3, Links: []int{1, 2}, Color: -1},
	},
}

func TestGetFirstAvailableColor(t *testing.T) {
	unavailableColors := map[int]bool{0: false, 1: false, 2: true, 3: false}

	expectedColor := 2

	actualColor := GetFirstAvailableColor(unavailableColors)

	if !reflect.DeepEqual(expectedColor, actualColor) {
		t.Fatalf("\nExpected %v \nGot %v", expectedColor, actualColor)
	}
}

func TestGetFirstAvailableColor_NoColors(t *testing.T) {
	unavailableColors := map[int]bool{}

	expectedColor := 0

	actualColor := GetFirstAvailableColor(unavailableColors)

	if !reflect.DeepEqual(expectedColor, actualColor) {
		t.Fatalf("\nExpected %v \nGot %v", expectedColor, actualColor)
	}
}

func TestGetFirstAvailableColor_NoAvailableColors(t *testing.T) {
	unavailableColors := map[int]bool{0: false, 1: false, 2: false, 3: false}

	expectedColor := 0

	actualColor := GetFirstAvailableColor(unavailableColors)

	if !reflect.DeepEqual(expectedColor, actualColor) {
		t.Fatalf("\nExpected %v \nGot %v", expectedColor, actualColor)
	}
}

func TestInitColorMap(t *testing.T) {
	count := 4

	expectedColorMap := map[int]bool{0: true, 1: true, 2: true, 3: true}

	actualColorMap := InitColorMap(count)

	if !reflect.DeepEqual(expectedColorMap, actualColorMap) {
		t.Fatalf("\nExpected %v \nGot %v", expectedColorMap, actualColorMap)
	}
}

func TestInitColorMap_NoCount(t *testing.T) {
	count := 0

	expectedColorMap := map[int]bool{}

	actualColorMap := InitColorMap(count)

	if !reflect.DeepEqual(expectedColorMap, actualColorMap) {
		t.Fatalf("\nExpected %v \nGot %v", expectedColorMap, actualColorMap)
	}
}

func TestInitColorMap_InvalidCount(t *testing.T) {
	count := -1

	expectedColorMap := map[int]bool{}

	actualColorMap := InitColorMap(count)

	if !reflect.DeepEqual(expectedColorMap, actualColorMap) {
		t.Fatalf("\nExpected %v \nGot %v", expectedColorMap, actualColorMap)
	}
}

func TestFindSmallestNode(t *testing.T) {
	graph := graphTest

	expectedNode := Node{Value: 1, Links: []int{2, 3}, Color: -1}

	actualNode := graph.FindSmallestNode()

	if !reflect.DeepEqual(expectedNode, actualNode) {
		t.Fatalf("\nExpected %v \nGot %v", expectedNode, actualNode)
	}
}

func TestFindSmallestNode_NoGraph(t *testing.T) {
	var graph Graph

	expectedNode := Node{}

	actualNode := graph.FindSmallestNode()

	if !reflect.DeepEqual(expectedNode, actualNode) {
		t.Fatalf("\nExpected %v \nGot %v", expectedNode, actualNode)
	}
}

func TestFindNode(t *testing.T) {
	graph := graphTest
	value := 2

	expectedNode := Node{Value: 2, Links: []int{1, 3}, Color: -1}

	actualNode := graph.FindNode(value)

	if !reflect.DeepEqual(expectedNode, actualNode) {
		t.Fatalf("\nExpected %v \nGot %v", expectedNode, actualNode)
	}
}

func TestFindNode_NoGraph(t *testing.T) {
	var graph Graph
	value := 2

	expectedNode := Node{}

	actualNode := graph.FindNode(value)

	if !reflect.DeepEqual(expectedNode, actualNode) {
		t.Fatalf("\nExpected %v \nGot %v", expectedNode, actualNode)
	}
}

func TestFindNode_NoValue(t *testing.T) {
	graph := graphTest
	value := 0

	expectedNode := Node{}

	actualNode := graph.FindNode(value)

	if !reflect.DeepEqual(expectedNode, actualNode) {
		t.Fatalf("\nExpected %v \nGot %v", expectedNode, actualNode)
	}
}

func TestAddNode(t *testing.T) {
	graph := graphTest
	value := 4
	links := []int{}

	expectedGraph := Graph{
		AdjacencyList: []Node{
			Node{Value: 1, Links: []int{2, 3}, Color: -1},
			Node{Value: 2, Links: []int{1, 3}, Color: -1},
			Node{Value: 3, Links: []int{1, 2}, Color: -1},
			Node{Value: 4, Links: []int{}, Color: -1},
		},
	}

	graph.AddNode(value, links)
	actualGraph := graph

	if !reflect.DeepEqual(expectedGraph, actualGraph) {
		t.Fatalf("\nExpected %v \nGot %v", expectedGraph, actualGraph)
	}
}

func TestAddNode_EmptyGraph(t *testing.T) {
	var graph Graph
	value := 1
	links := []int{}

	expectedGraph := Graph{
		AdjacencyList: []Node{
			Node{Value: 1, Links: []int{}, Color: -1},
		},
	}

	graph.AddNode(value, links)
	actualGraph := graph

	if !reflect.DeepEqual(expectedGraph, actualGraph) {
		t.Fatalf("\nExpected %v \nGot %v", expectedGraph, actualGraph)
	}
}

func TestGenerateColors(t *testing.T) {
	graph := graphTest

	expectedGraph := Graph{
		AdjacencyList: []Node{
			Node{Value: 1, Links: []int{2, 3}, Color: 0},
			Node{Value: 2, Links: []int{1, 3}, Color: 1},
			Node{Value: 3, Links: []int{1, 2}, Color: 2},
		},
	}

	graph.GenerateColors()
	actualGraph := graph

	if !reflect.DeepEqual(expectedGraph, actualGraph) {
		t.Fatalf("\nExpected %v \nGot %v", expectedGraph, actualGraph)
	}
}

func TestGenerateColors_NoGraph(t *testing.T) {
	var graph Graph

	var expectedGraph Graph

	graph.GenerateColors()
	actualGraph := graph

	if !reflect.DeepEqual(expectedGraph, actualGraph) {
		t.Fatalf("\nExpected %v \nGot %v", expectedGraph, actualGraph)
	}
}

func TestGenerateColors_SingleNode(t *testing.T) {
	var graph Graph
	graph.AddNode(1, []int{})

	expectedGraph := Graph{
		AdjacencyList: []Node{
			Node{Value: 1, Links: []int{}, Color: 0},
		},
	}

	graph.GenerateColors()
	actualGraph := graph

	if !reflect.DeepEqual(expectedGraph, actualGraph) {
		t.Fatalf("\nExpected %v \nGot %v", expectedGraph, actualGraph)
	}
}

func TestGenerateColors_TwoSeperatedNodes(t *testing.T) {
	var graph Graph
	graph.AddNode(1, []int{})
	graph.AddNode(2, []int{})

	expectedGraph := Graph{
		AdjacencyList: []Node{
			Node{Value: 1, Links: []int{}, Color: 0},
			Node{Value: 2, Links: []int{}, Color: 0},
		},
	}

	graph.GenerateColors()
	actualGraph := graph

	if !reflect.DeepEqual(expectedGraph, actualGraph) {
		t.Fatalf("\nExpected %v \nGot %v", expectedGraph, actualGraph)
	}
}

func TestGenerateColors_TwoConnectedNodes(t *testing.T) {
	var graph Graph
	graph.AddNode(1, []int{2})
	graph.AddNode(2, []int{1})

	expectedGraph := Graph{
		AdjacencyList: []Node{
			Node{Value: 1, Links: []int{2}, Color: 0},
			Node{Value: 2, Links: []int{1}, Color: 1},
		},
	}

	graph.GenerateColors()
	actualGraph := graph

	if !reflect.DeepEqual(expectedGraph, actualGraph) {
		t.Fatalf("\nExpected %v \nGot %v", expectedGraph, actualGraph)
	}
}
