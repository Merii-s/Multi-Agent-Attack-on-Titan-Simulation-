package pkg

import (
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

	return true // Collided on both axes
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
