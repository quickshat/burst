package main

// Velocity id for impact map
const (
	VelocityGravity = 1
	VelocityShot    = 2
)

type velocity struct {
	LocalVector vector2f
	Strength    float64
}
