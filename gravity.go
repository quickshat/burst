package main

var gravityVecNorm = vector2f{0, 0.5}
var gravityScaleNorm = 1.5

type gravity interface {
	applyGravity()
}
