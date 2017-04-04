package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 800
	screenHeight = 600
	screenTitle  = "Burst"
)

var myGame *game

func main() {
	myGame = newGame()
	myGame.setActiveGamestate(GamestateIngame)

	if err := ebiten.Run(myGame.getCurrentUpdate(), screenWidth, screenHeight, 1, screenTitle); err != nil {
		log.Fatal(err)
	}
}
