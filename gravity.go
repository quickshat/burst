package main

var gravityVecNorm = vector2f{0, 0.20}
var gravityScaleNorm = 1.01

type gravity interface {
	applyGravity()
}
