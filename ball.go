package main

import "github.com/hajimehoshi/ebiten"

type ball struct {
	Size        float64
	R           float64
	Scale       float64
	Pos         *vector2f
	Impacts     map[int]*velocity
	Sprite      *ebiten.Image
	Settings    *ebiten.DrawImageOptions
	Stopped     bool
	IsColliding bool
	Parent      *inGameState
}

func newBall(size float64, pos *vector2f, stop bool, i *ebiten.Image, p *inGameState) *ball {
	tmp := ball{}
	tmp.Impacts = make(map[int]*velocity)
	tmp.Pos = pos
	tmp.Stopped = stop
	tmp.Size = size
	tmp.R = size / 2
	tmp.Sprite = i
	tmp.Scale = size / 100
	tmp.Settings = &ebiten.DrawImageOptions{}
	tmp.Settings.GeoM.Scale(tmp.Scale, tmp.Scale)
	tmp.Parent = p
	return &tmp
}

func (b *ball) update() {
	if b.Parent.Parent.Tick%tickrate != 0 {
		return
	}
	// Perform calulateable changes
	b.applyGravity()
	b.calculateResultingVector()
	// Update Render Settings
	b.Settings = &ebiten.DrawImageOptions{}
	b.Settings.GeoM.Scale(b.Scale, b.Scale)
	b.Settings.GeoM.Translate(b.Pos.X, b.Pos.Y)
}

func (b *ball) applyGravity() {
	if b.Stopped || b.IsColliding {
		return
	}
	if b.Impacts[VelocityGravity] == nil || b.Pos.Y > screenHeight {
		b.Impacts[VelocityGravity] = &velocity{gravityVecNorm, gravityScaleNorm}
	}
	if b.Pos.Y > screenHeight {
		b.Pos.X = 250
		b.Pos.Y = 250
	}
	b.Impacts[VelocityGravity].Strength = b.Impacts[VelocityGravity].Strength * 1.05
}

func (b *ball) calculateResultingVector() {
	resolution := &vector2f{0, 0}
	for _, v := range b.Impacts {
		resolution = v.LocalVector.scale(v.Strength).summarize(*resolution)
	}
	b.Pos = b.Pos.summarize(*resolution)
}
