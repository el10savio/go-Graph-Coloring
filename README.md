# Graph Coloring
A greedy graph coloring algorithm implemented in Go

Here for a given graph it determines the best possible solution to color each node so that it does not share the same color with its neighbors.

**Methodology**

It first assigns the first node with the first available color and then for each corresponding node creates a hash map of a list of available colors and marks the colors as false if its neighbors contains that color. The maximum number of colors the graph could have is the number of nodes it contains.