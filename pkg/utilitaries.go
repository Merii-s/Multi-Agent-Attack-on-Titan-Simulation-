package pkg

import (
	"log"
	"math/rand"
	"os"
	"runtime"
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

// func GetPath_Win(imgName string) string {
// 	currentDir, _ := os.Getwd()
// 	path := currentDir + "/assets/" + imgName + ".png"
// 	return path
// }

func GetImagePath(imgName string) string {
	currentDir, _ := os.Getwd()
	var path string

	if runtime.GOOS == "windows" {
		path = currentDir + "/assets/" + imgName + ".png"
	} else if runtime.GOOS == "darwin" {
		path = currentDir + "/../../assets/" + imgName + ".png"
	} else {
		log.Fatal("OS not supported")
	}

	return path
}
