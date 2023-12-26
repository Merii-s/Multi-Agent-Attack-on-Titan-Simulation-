package pkg

import (
	"container/heap"
	"math"
	"math/rand"
	"os"
	"reflect"
	"time"
)

func GetRandomCoords(topLeft Position, bottomRight Position) (int, int) {
	// Nouvelle seed pour le generateur
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	time.Sleep(1 * time.Millisecond)
	// Generation des coordonnees aleatoires
	randomX := random.Intn(bottomRight.X-topLeft.X) + topLeft.X
	randomY := random.Intn(bottomRight.Y-topLeft.Y) + topLeft.Y

	return randomX, randomY
}

func GetImagePath(imgName string) string {
	currentDir, _ := os.Getwd()
	path := currentDir + "/assets/" + imgName + ".png"
	return path
}

// DetectCollision checks if there is a collision between two objects using AABB collision detection
func DetectCollision(obj1, obj2 Object) bool {

	obj1TopLeft, obj1BottomRight := obj1.Hitbox()[0], obj1.Hitbox()[1]
	obj2TopLeft, obj2BottomRight := obj2.Hitbox()[0], obj2.Hitbox()[1]

	// Check for collision on the X-axis
	if obj1BottomRight.X < obj2TopLeft.X || obj1TopLeft.X > obj2BottomRight.X {
		return false // No collision on X-axis
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

// A modifier quand les constructeurs seront pret
func PlaceHuman(objs []Object, humans []Object, human *Object, tl_village Position, br_village Position) []Object {
	end := false
	for !end {
		counter := 0
		nb_grass := 0
		for _, obj := range objs {
			if obj.name != Grass {
				if DetectCollision(*human, obj) {
					x, y := GetRandomCoords(tl_village, br_village)
					human.SetPosition(Position{x, y})
					break
				} else {
					counter++
				}
			} else {
				nb_grass++
			}
		}

		for _, hu := range humans {
			if DetectCollision(*human, hu) {
				x, y := GetRandomCoords(tl_village, br_village)
				human.SetPosition(Position{x, y})
				break
			} else {
				counter++
			}
		}

		if counter == len(objs)+len(humans)-nb_grass {
			end = true
		}
	}

	humans = append(humans, *human)
	return humans
}

func PlaceTitan(titan *Object, titans []Object, W int, H int, dir int) []Object {
	tl_screen := Position{X: 0, Y: 0}
	br_screen := Position{X: W, Y: H}
	counter := 0
	end := false
	for !end {
		for _, ti := range titans {
			if DetectCollision(*titan, ti) {
				x, y := GetRandomCoords(tl_screen, br_screen)
				if dir == 0 {
					y = y - H
				} else if dir == 1 {
					x = x - W
				} else {
					x = x + W
				}
				titan.SetPosition(Position{x, y})
				break
			} else {
				counter++
			}
		}

		if counter == len(titans) {
			end = true
		}
	}

	titans = append(titans, *titan)
	return titans
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

// contains checks if the given list contains the specified object.
func contains[T any](list []T, target T) bool {
	for _, item := range list {
		if reflect.DeepEqual(item, target) {
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

func removeObjectsBehindPositions(perceivedObjects []Object, objectsBehindPositions []Position) []Object {
	// Filter out positions behind an obstacle if the center of the object is in the objectsBehindPositions list
	objectsToRemove := []Object{}
	for _, object := range perceivedObjects {
		if contains(objectsBehindPositions, object.Center()) || object.Name() == Grass {
			objectsToRemove = append(objectsToRemove, object)
		}
	}

	// Remove the objects in the objectsToRemove list from the perceivedObjects list
	perceivedObjects = removeObjects(perceivedObjects, objectsToRemove)

	return perceivedObjects
}

// generalize removeObjectsBehindPositions with any type

func removeAgentsBehindPositions(perceptedAgents []AgentI, objectsBehindPositions []Position) []AgentI {
	// Filter out positions behind an obstacle if the center of the object is in the objectsBehindPositions list
	objectsToRemove := []AgentI{}
	for _, object := range perceptedAgents {
		if contains(objectsBehindPositions, object.Pos()) {
			objectsToRemove = append(objectsToRemove, object)
		}
	}

	// Remove the objects in the objectsToRemove list from the perceivedObjects list
	perceptedAgents = removeAgents(perceptedAgents, objectsToRemove)

	return perceptedAgents
}

func removeObjects(perceivedObjects []Object, objectsToRemove []Object) []Object {
	// Remove the objects in the objectsToRemove list from the perceivedObjects list
	for _, object := range objectsToRemove {
		for i, perceptedObject := range perceivedObjects {
			if perceptedObject == object {
				perceivedObjects = append(perceivedObjects[:i], perceivedObjects[i+1:]...)
			}
		}
	}

	return perceivedObjects
}

func removeAgents(perceptedAgents []AgentI, objectsToRemove []AgentI) []AgentI {
	// Remove the objects in the objectsToRemove list from the perceptedObjects list
	for _, object := range objectsToRemove {
		for i, perceptedObject := range perceptedAgents {
			if perceptedObject == object {
				perceptedAgents = append(perceptedAgents[:i], perceptedAgents[i+1:]...)
			}
		}
	}

	return perceptedAgents
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

// define a function that takes two squares in parameters defined by top left and bottom right positions and returns true if they intersect
func IntersectSquare(topLeft1 Position, bottomRight1 Position, topLeft2 Position, bottomRight2 Position) bool {
	// Check if the two squares intersect
	if topLeft1.X > bottomRight2.X || topLeft2.X > bottomRight1.X {
		return false
	}
	if topLeft1.Y < bottomRight2.Y || topLeft2.Y < bottomRight1.Y {
		return false
	}

	return true
}

// TODO : Maybe use angle parameters instead of forced values
func getObjectsBehindPositions(position Position, angle float64, topLeftVision Position, bottomRightVision Position) []Position {
	objectsBehindPositions := []Position{}
	// the angle order follows the counter-clockwise order

	// if the position to avoid is in the straight right of the agent position
	// the agent can't see the positions behind it from 315 to 45 degrees following the perspective logic
	if angle > 345 && angle < 15 {
		for i := 0; i < bottomRightVision.X; i++ {
			objectsBehindPositions = append(objectsBehindPositions, Position{position.X + i, position.Y})
		}

		// if the position to avoid is in the bottom right quarter of the vision square
		// the agent can't see the positions in the bottom right quarter of the vision square
	} else if angle >= 15 && angle < 75 {
		for x := position.X; x <= bottomRightVision.X; x++ {
			for y := position.Y; y <= topLeftVision.Y; y++ {
				objectsBehindPositions = append(objectsBehindPositions, Position{x, y})
			}
		}
		// if the position to avoid is in the straight bottom of the agent position
		// the agent can't see the positions behind it from 45 to 135 degrees following the perspective logic
	} else if angle > 75 && angle < 105 {
		for i := 0; i < bottomRightVision.Y; i++ {
			objectsBehindPositions = append(objectsBehindPositions, Position{position.X, position.Y - i})
		}

		// if the position to avoid is in the bottom left of the agent position
		// the agent can't see the positions behind it from 90 to 180 degrees following the perspective logic
	} else if angle >= 105 && angle <= 165 {
		for i := 0; i < 8; i++ {
			objectsBehindPositions = append(objectsBehindPositions, Position{position.X - i, position.Y - i})
		}

		// if the position to avoid is in the straight left of the agent position
		// the agent can't see the positions behind it from 135 to 225 degrees following the perspective logic
	} else if angle > 165 && angle < 195 {
		for i := 0; i < topLeftVision.X; i++ {
			objectsBehindPositions = append(objectsBehindPositions, Position{position.X - i, position.Y})
		}

		// if the position to avoid is in the top left of the agent position
		// the agent can't see the positions behind it from 180 to 270 degrees following the perspective logic
	} else if angle >= 195 && angle <= 255 {
		for i := 0; i < 8; i++ {
			objectsBehindPositions = append(objectsBehindPositions, Position{position.X - i, position.Y + i})
		}

		// if the position to avoid is in the straight top of the agent position
		// the agent can't see the positions behind it from 225 to 315 degrees following the perspective logic
	} else if angle > 255 && angle < 285 {
		for i := 0; i < topLeftVision.Y; i++ {
			objectsBehindPositions = append(objectsBehindPositions, Position{position.X, position.Y + i})
		}

		// if the position to avoid is in the top right of the agent position
		// the agent can't see the positions behind it from 270 to 360 degrees and from 0 to 90 degrees following the perspective logic
	} else if angle >= 285 && angle <= 345 || angle >= 0 && angle <= 90 {
		for i := 0; i < 8; i++ {
			objectsBehindPositions = append(objectsBehindPositions, Position{position.X + i, position.Y + i})
		}

	}
	return objectsBehindPositions
}

// Calculate the opposite direction of a position
func OppositeDirection(currentPos, targetPos Position) Position {
	// TO DO: Check if the position is valid
	return Position{X: 2*currentPos.X - targetPos.X, Y: 2*currentPos.Y - targetPos.Y}
}
