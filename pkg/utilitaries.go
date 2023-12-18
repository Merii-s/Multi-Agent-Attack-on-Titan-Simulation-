package pkg

import (
	"container/heap"
	"math/rand"
	"os"
	"time"
)

func GetRandomCoords(topLeft Position, bottomRight Position) (int, int) {
	// Verification des parametres d'entree
	if topLeft.X >= bottomRight.X || topLeft.Y <= bottomRight.Y {
		return 0, 0
	}

	// Nouvelle seed pour le generateur
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generation des coordonnees aleatoires
	randomX := random.Intn(bottomRight.X-topLeft.X) + topLeft.X
	randomY := random.Intn(topLeft.Y-bottomRight.Y) + bottomRight.Y

	return randomX, randomY
}

func GetPath(imgName string) string {
	currentDir, _ := os.Getwd()
	path := currentDir + "/assets/" + imgName + ".png"
	return path
}

// define a function that takes a position in argument and an agent position
// and returns the shortest path to reach the position from the agent position

type Node struct {
	position Position
	parent   *Node
	g        int
	h        int
	f        int
}

type NodeHeap []*Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].f < h[j].f }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	node := old[n-1]
	*h = old[0 : n-1]
	return node
}

func calculateHeuristic(start Position, target Position) int {
	dx := abs(target.X - start.X)
	dy := abs(target.Y - start.Y)
	return dx + dy + min(dx, dy)
}

func getNeighbors(position Position, toAvoid []Position) []Position {
	neighbors := []Position{
		{position.X, position.Y - 1},     // above
		{position.X, position.Y + 1},     // below
		{position.X - 1, position.Y},     // left
		{position.X + 1, position.Y},     // right
		{position.X - 1, position.Y - 1}, // top left
		{position.X + 1, position.Y - 1}, // top right
		{position.X - 1, position.Y + 1}, // bottom left
		{position.X + 1, position.Y + 1}, // bottom right
	}

	// Filter out positions to avoid
	filteredNeighbors := []Position{}
	for _, neighbor := range neighbors {
		if !contains(toAvoid, neighbor) {
			filteredNeighbors = append(filteredNeighbors, neighbor)
		}
	}

	return filteredNeighbors
}

func contains(positions []Position, position Position) bool {
	for _, p := range positions {
		if p == position {
			return true
		}
	}
	return false
}

func GetShortestPath(pos Position, agentPos Position, agentSpeed int, toAvoid []Position) Position {
	openSet := make(NodeHeap, 0)
	closedSet := make(NodeHeap, 0)

	startNode := &Node{
		position: agentPos,
		parent:   nil,
		g:        0,
		h:        calculateHeuristic(agentPos, pos),
		f:        0,
	}

	heap.Push(&openSet, startNode)

	for len(openSet) > 0 {
		currentNode := heap.Pop(&openSet).(*Node)

		// If the current node is the target node, reconstruct the path
		if currentNode.position == pos {
			path := []Position{}
			current := currentNode
			for current != nil {
				path = append(path, current.position)
				current = current.parent
			}
			// Return the position of the first node in the path
			return path[len(path)-1]
		}

		// Add the current node to the closed set
		heap.Push(&closedSet, currentNode)

		// Get the neighboring positions of the current node
		neighbors := getNeighbors(currentNode.position, toAvoid)

		for _, neighbor := range neighbors {
			// Skip if the neighbor is in the closed set
			if containsNode(&closedSet, neighbor) {
				continue
			}

			// Calculate the tentative g value for the neighbor
			tentativeG := currentNode.g + agentSpeed

			// Check if the neighbor is not in the open set or the tentative g value is lower
			neighborNode := getNode(&openSet, neighbor)
			if neighborNode == nil || tentativeG < neighborNode.g {
				// Create a new node for the neighbor
				neighborNode = &Node{
					position: neighbor,
					parent:   currentNode,
					g:        tentativeG,
					h:        calculateHeuristic(neighbor, pos),
					f:        tentativeG + calculateHeuristic(neighbor, pos),
				}

				// Add the neighbor node to the open set
				heap.Push(&openSet, neighborNode)
			}
		}
	}

	// If no path is found, return the agent position
	return agentPos
}

func containsNode(nodes *NodeHeap, position Position) bool {
	for _, node := range *nodes {
		if node.position.Equals(position) {
			return true
		}
	}
	return false
}

func getNode(nodes *NodeHeap, position Position) *Node {
	for _, node := range *nodes {
		if node.position.Equals(position) {
			return node
		}
	}
	return nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (p Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}
