package main

import (
	"github.com/hajimehoshi/ebiten"
)

// Game state identifiers
const (
	GamestateIngame = 0
)

type game struct {
	Tick int
	IGS  *inGameState
}

func newGame() *game {
	tmp := game{}
	tmp.IGS = newInGameState(&tmp)
	return &tmp
}

func (g *game) getCurrentUpdate() func(*ebiten.Image) error {
	if g.IGS.isActive() {
		return g.IGS.update
	}
	g.Tick++
	return nil
}

func (g *game) setActiveGamestate(state int) {
	//disable all
	g.IGS.setActive(false)
	//enable new state
	if state == GamestateIngame {
		g.IGS.setActive(true)
	}
}
