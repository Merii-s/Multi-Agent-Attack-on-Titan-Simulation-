package pkg

import (
	"container/heap"
	"math"
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

// define a function to get the vision of an agent
// it takes the agent position, the agent vision and positions he can't see behind in argument
// returns a list of positions that the agent can see (including the agent position)
// the vision is a square centered on the agent position
func GetVision(agentPos Position, agentVision int, toAvoid []Position) []Position {
	// Get the top left and bottom right positions of the vision square
	topLeft := Position{agentPos.X - agentVision, agentPos.Y + agentVision}
	bottomRight := Position{agentPos.X + agentVision, agentPos.Y - agentVision}

	// Get the positions inside the vision square
	positions := []Position{}
	for x := topLeft.X; x <= bottomRight.X; x++ {
		for y := bottomRight.Y; y <= topLeft.Y; y++ {
			positions = append(positions, Position{x, y})
		}
	}

	// Filter out positions to avoid regarding the agent position
	// If a toAvoid position is in the vision square, agent can't see behind it,
	// depending on the angle between the agent and the position to avoid, the agent can't see every positions behind it following the perspective logic
	for _, position := range positions {
		if !contains(toAvoid, position) {
			continue
		} else {
			// If the position to avoid is in the vision square, the agent can't see behind it
			// Get the angle between the agent and the position to avoid
			angle := getAngle(agentPos, position)

			// Get the positions behind the position to avoid
			positionsBehind := getPositionsBehind(position, angle, topLeft, bottomRight)

			// Filter out positions behind the position to avoid
			positions = filterOut(positionsBehind, positions)
		}
	}
	return positions
}

func remove(positions []Position, position Position) []Position {
	for i, p := range positions {
		if p == position {
			return append(positions[:i], positions[i+1:]...)
		}
	}
	return positions
}

func filterOut(positions []Position, toFilter []Position) []Position {
	filteredPositions := toFilter
	for _, position := range positions {
		filteredPositions = remove(filteredPositions, position)
	}
	return filteredPositions
}

// TODO : improve the angle for straight positions and maybe deepness of unseeable positions
func getPositionsBehind(position Position, angle float64, bottomLeft Position, topRight Position) []Position {
	positions := []Position{}
	// the angle order follows the counter-clockwise order
	// if the position to avoid is in the bottom right quarter of the vision square
	// the agent can't see the positions in the bottom right quarter of the vision square
	if angle >= 0 && angle < 90 {
		for x := position.X; x <= topRight.X; x++ {
			for y := position.Y; y <= bottomLeft.Y; y++ {
				positions = append(positions, Position{x, y})
			}
		}
		// if the position to avoid is in the bottom left quarter of the vision square
		// the agent can't see the positions in the bottom left quarter of the vision square
	} else if angle >= 90 && angle < 180 {
		for x := bottomLeft.X; x <= position.X; x++ {
			for y := position.Y; y <= bottomLeft.Y; y++ {
				positions = append(positions, Position{x, y})
			}
		}
		// if the position to avoid is in the top left quarter of the vision square
		// the agent can't see the positions in the top left quarter of the vision square
	} else if angle >= 180 && angle < 270 {
		for x := bottomLeft.X; x <= position.X; x++ {
			for y := topRight.Y; y <= position.Y; y++ {
				positions = append(positions, Position{x, y})
			}
		}
		// if the position to avoid is in the top right quarter of the vision square
		// the agent can't see the positions in the top right quarter of the vision square
	} else if angle >= 270 && angle < 360 {
		for x := position.X; x <= topRight.X; x++ {
			for y := topRight.Y; y <= position.Y; y++ {
				positions = append(positions, Position{x, y})
			}
		}
	}

	return positions
}

// define a function to get the angle between two positions
func getAngle(agentPos Position, position Position) float64 {
	// Calculate the vector from agentPos to position
	deltaX := position.X - agentPos.X
	deltaY := position.Y - agentPos.Y

	// Use Atan2 to get the angle in radians
	angleRad := math.Atan2(float64(deltaY), float64(deltaX))

	// Convert radians to degrees
	angleDeg := angleRad * 180 / math.Pi

	// Ensure the angle is in the range [0, 360)
	if angleDeg < 0 {
		angleDeg += 360
	}

	return angleDeg
}

// define a function to get the distance between two positions
func getDistance(agentPos Position, position Position) float64 {
	// Get the distance between the two positions using the Pythagorean theorem
	distance := math.Sqrt(float64((agentPos.X-position.X)*(agentPos.X-position.X) + (agentPos.Y-position.Y)*(agentPos.Y-position.Y)))

	return distance
}

func Contains(positions []Position, position Position) bool {
	for _, p := range positions {
		if p == position {
			return true
		}
	}
	return false
}
