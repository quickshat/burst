package main

import (
	"log"

	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 800
	screenHeight = 600
	screenTitle  = "Burst"
	tickrate     = 2
)

var myGame *game

func main() {
	rand.Seed(int64((time.Now().Nanosecond())))
	myGame = newGame()
	myGame.setActiveGamestate(GamestateIngame)

	if err := ebiten.Run(myGame.getCurrentUpdate(), screenWidth, screenHeight, 1, screenTitle); err != nil {
		log.Fatal(err)
	}
}
