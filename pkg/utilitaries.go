package pkg

import (
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
