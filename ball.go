package main

import "github.com/hajimehoshi/ebiten"

type ball struct {
	Size     float64
	R        float64
	Scale    float64
	Pos      vector2f
	Impacts  map[int]velocity
	Sprite   *ebiten.Image
	Settings *ebiten.DrawImageOptions
	Stopped  bool
}

func newBall(size float64, pos vector2f, stop bool, i *ebiten.Image) *ball {
	tmp := ball{}
	tmp.Impacts = make(map[int]velocity)
	tmp.Pos = pos
	tmp.Stopped = stop
	tmp.Size = size
	tmp.R = size / 2
	tmp.Sprite = i
	tmp.Scale = size / 100
	tmp.Settings = &ebiten.DrawImageOptions{}
	tmp.Settings.GeoM.Scale(tmp.Scale, tmp.Scale)
	return &tmp
}

func (b *ball) update() {
	// Perform calulateable changes

	// Update Render Settings
	b.Settings = &ebiten.DrawImageOptions{}
	b.Settings.GeoM.Scale(b.Scale, b.Scale)
	b.Settings.GeoM.Translate(b.Pos.X, b.Pos.Y)
}
