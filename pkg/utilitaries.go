package pkg

import (
	"fmt"
	"math/rand"
	"os"
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

	obj1TopLeft, obj1BottomRight := obj1.hitbox()[0], obj1.hitbox()[1]
	obj2TopLeft, obj2BottomRight := obj2.hitbox()[0], obj2.hitbox()[1]

	// Check for collision on the X-axis
	if obj1BottomRight.X < obj2TopLeft.X || obj1TopLeft.X > obj2BottomRight.X {
		return false // No collision on X-axis
	}

	// Check for collision on the Y-axis
	if obj1BottomRight.Y < obj2TopLeft.Y || obj1TopLeft.Y > obj2BottomRight.Y {
		return false // No collision on Y-axis
	}

	fmt.Print("Collision between ", obj1.Name(), " and ", obj2.Name(), "\n")

	return true // Collided on both axes
}
