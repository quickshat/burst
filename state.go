package main

import (
	"github.com/hajimehoshi/ebiten"
)

type state interface {
	update(*ebiten.Image) error
	isActive() bool
	setActive(bool)
}
