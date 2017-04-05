package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type inGameState struct {
	Parent      *game
	Active      bool
	Balls       []*ball
	Initialized bool
}

func newInGameState(parent *game) *inGameState {
	tmp := inGameState{}
	tmp.Parent = parent
	return &tmp
}

func (s *inGameState) update(i *ebiten.Image) error {
	s.logic()
	err := s.draw(i)
	return err
}

func (s *inGameState) isActive() bool {
	return s.Active
}

func (s *inGameState) setActive(b bool) {
	s.Active = b
}

func (s *inGameState) init() {
	ballSprite, _, _ := ebitenutil.NewImageFromFile("circle.png", ebiten.FilterNearest)
	s.Balls = append(s.Balls, newBall(15, &vector2f{250, 250}, false, ballSprite, s))
}

func (s *inGameState) logic() {
	if !s.Initialized {
		s.init()
		s.Initialized = true
	}
	// Main Update Loop
	for _, v := range s.Balls {
		v.update()
	}
}

func (s *inGameState) draw(g *ebiten.Image) error {
	for _, v := range s.Balls {
		g.DrawImage(v.Sprite, v.Settings)
	}
	return nil
}
