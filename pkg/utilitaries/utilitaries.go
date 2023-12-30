package pkg

import (
	obj "AOT/pkg/obj"
	types "AOT/pkg/types"
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"time"
)

func GetRandomCoords(topLeft types.Position, bottomRight types.Position) (int, int) {
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

func CreateAgentID(agentNb int, agentType types.ObjectName) types.Id {
	return types.Id(fmt.Sprint("AGTID", fmt.Sprint(agentNb), "_", string(agentType)))
}

// DetectCollision checks if there is a collision between two objects using AABB collision detection
func DetectCollision(obj1, obj2 obj.Object) bool {

	obj1TopLeft, obj1BottomRight := obj1.Hitbox()[0], obj1.Hitbox()[1]
	obj2TopLeft, obj2BottomRight := obj2.Hitbox()[0], obj2.Hitbox()[1]

	// Check for collision on the X-axis
	if obj1BottomRight.X < obj2TopLeft.X || obj1TopLeft.X > obj2BottomRight.X {
		return false // No collision on X-axis
	}

	// Check for collision on the Y-axis
	if obj1BottomRight.Y < obj2TopLeft.Y || obj1TopLeft.Y > obj2BottomRight.Y {
		return false // No collision on Y-axis
	}

	return true // Collided on both axes
}

func GetNeighbors(position types.Position, speed int, toAvoid []types.Position) []types.Position {
	neighbors := []types.Position{
		{X: position.X, Y: position.Y - speed},         // above
		{X: position.X, Y: position.Y + speed},         // below
		{X: position.X - speed, Y: position.Y},         // left
		{X: position.X + speed, Y: position.Y},         // right
		{X: position.X - speed, Y: position.Y - speed}, // top left
		{X: position.X + speed, Y: position.Y - speed}, // top right
		{X: position.X - speed, Y: position.Y + speed}, // bottom left
		{X: position.X + speed, Y: position.Y + speed}, // bottom right
	}

	// Filter out positions to avoid
	filteredNeighbors := []types.Position{}
	for _, neighbor := range neighbors {
		if !Contains(toAvoid, neighbor) {
			filteredNeighbors = append(filteredNeighbors, neighbor)
		}
	}

	return filteredNeighbors
}

// contains checks if the given list contains the specified object.
func Contains[T any](list []T, target T) bool {
	for _, item := range list {
		if reflect.DeepEqual(item, target) {
			return true
		}
	}
	return false
}

func PositionsBehindObjects(perceivedObjects []obj.Object, positionsBehindObjects []types.Position) []obj.Object {
	// Filter out positions behind an obstacle if the center of the object is in the positionsBehindObjects list
	objectsToRemove := []obj.Object{}
	for _, object := range perceivedObjects {
		if Contains(positionsBehindObjects, object.Center()) || object.Name() == types.Grass {
			objectsToRemove = append(objectsToRemove, object)
		}
	}

	// Remove the objects in the objectsToRemove list from the perceivedObjects list
	perceivedObjects = RemoveObjects(perceivedObjects, objectsToRemove)

	return perceivedObjects
}

func RemoveObjects(perceivedObjects []obj.Object, objectsToRemove []obj.Object) []obj.Object {
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

// define a function to get the angle between two positions
func GetAngle(agentPos types.Position, position types.Position) float64 {
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
func IntersectSquare(topLeft1 types.Position, bottomRight1 types.Position, topLeft2 types.Position, bottomRight2 types.Position) bool {
	// Check if the two squares intersect
	if topLeft1.X > bottomRight2.X || topLeft2.X > bottomRight1.X || topLeft1.Y > bottomRight2.Y || topLeft2.Y > bottomRight1.Y {
		return false
	}

	return true
}

func GetPositionsBehindObject(position types.Position, angle float64, topLeftVision types.Position, bottomRightVision types.Position) []types.Position {
	positionsBehindObjects := []types.Position{}
	// the angle order follows the counter-clockwise order

	// if the position to avoid is in the straight right of the agent position
	// the agent can't see the positions behind it from 315 to 45 degrees following the perspective logic
	if angle > 345 && angle < 15 {
		for i := 0; i < bottomRightVision.X; i++ {
			positionsBehindObjects = append(positionsBehindObjects, types.Position{X: position.X + i, Y: position.Y})
		}

		// if the position to avoid is in the bottom right quarter of the vision square
		// the agent can't see the positions in the bottom right quarter of the vision square
	} else if angle >= 15 && angle < 75 {
		for x := position.X; x <= bottomRightVision.X; x++ {
			for y := position.Y; y <= topLeftVision.Y; y++ {
				positionsBehindObjects = append(positionsBehindObjects, types.Position{X: x, Y: y})
			}
		}
		// if the position to avoid is in the straight bottom of the agent position
		// the agent can't see the positions behind it from 45 to 135 degrees following the perspective logic
	} else if angle > 75 && angle < 105 {
		for i := 0; i < bottomRightVision.Y; i++ {
			positionsBehindObjects = append(positionsBehindObjects, types.Position{X: position.X, Y: position.Y - i})
		}

		// if the position to avoid is in the bottom left of the agent position
		// the agent can't see the positions behind it from 90 to 180 degrees following the perspective logic
	} else if angle >= 105 && angle <= 165 {
		for i := 0; i < 8; i++ {
			positionsBehindObjects = append(positionsBehindObjects, types.Position{X: position.X - i, Y: position.Y - i})
		}

		// if the position to avoid is in the straight left of the agent position
		// the agent can't see the positions behind it from 135 to 225 degrees following the perspective logic
	} else if angle > 165 && angle < 195 {
		for i := 0; i < topLeftVision.X; i++ {
			positionsBehindObjects = append(positionsBehindObjects, types.Position{X: position.X - i, Y: position.Y})
		}

		// if the position to avoid is in the top left of the agent position
		// the agent can't see the positions behind it from 180 to 270 degrees following the perspective logic
	} else if angle >= 195 && angle <= 255 {
		for i := 0; i < 8; i++ {
			positionsBehindObjects = append(positionsBehindObjects, types.Position{X: position.X - i, Y: position.Y + i})
		}

		// if the position to avoid is in the straight top of the agent position
		// the agent can't see the positions behind it from 225 to 315 degrees following the perspective logic
	} else if angle > 255 && angle < 285 {
		for i := 0; i < topLeftVision.Y; i++ {
			positionsBehindObjects = append(positionsBehindObjects, types.Position{X: position.X, Y: position.Y + i})
		}

		// if the position to avoid is in the top right of the agent position
		// the agent can't see the positions behind it from 270 to 360 degrees and from 0 to 90 degrees following the perspective logic
	} else if angle >= 285 && angle <= 345 || angle >= 0 && angle <= 90 {
		for i := 0; i < 8; i++ {
			positionsBehindObjects = append(positionsBehindObjects, types.Position{X: position.X + i, Y: position.Y + i})
		}

	}
	return positionsBehindObjects
}

// Calculate the opposite direction of a position
func OppositeDirection(currentPos, targetPos types.Position) types.Position {
	// Check if the position is valid
	xToGo := 2*currentPos.X - targetPos.X
	yToGo := 2*currentPos.Y - targetPos.Y
	if xToGo < 0 || yToGo < 0 {
		return currentPos
	}
	// TODO: Add more checks if needed

	return types.Position{X: xToGo, Y: yToGo}
}

func IsOutOfScreen(pos types.Position, W int, H int) bool {
	return pos.X < 0 || pos.X > W || pos.Y < 0 || pos.Y > H
}

func GetPositionsInHitbox(tl types.Position, br types.Position) (inHitboxPositions []types.Position) {
	for posX := tl.X; posX <= br.X; posX++ {
		for posY := tl.Y; posY <= br.Y; posY++ {
			inHitboxPositions = append(inHitboxPositions, types.Position{X: posX, Y: posY})
		}
	}
	return
}

func ClosestObject(objects []obj.Object, position types.Position) (obj.Object, types.Position) {
	// Get the closest position from the list
	closestObject := objects[0]
	closestObjectPosition := objects[0].Hitbox()[0]
	for _, object := range objects {
		for _, pos := range GetPositionsInHitbox(object.Hitbox()[0], object.Hitbox()[1]) {
			if position.Distance(pos) < position.Distance(closestObjectPosition) {
				closestObject = object
				closestObjectPosition = pos
			}
		}
	}
	return closestObject, closestObjectPosition
}
