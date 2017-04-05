package main

type vector2f struct {
	X float64
	Y float64
}

func (v *vector2f) substract(v2 vector2f) *vector2f {
	res := vector2f{}
	res.X = v.X - v2.X
	res.Y = v.Y - v2.Y
	return &res
}

func (v *vector2f) summarize(v2 vector2f) *vector2f {
	res := vector2f{0, 0}
	res.X = v.X + v2.X
	res.Y = v.Y + v2.Y
	return &res
}

func (v *vector2f) scale(f float64) *vector2f {
	res := vector2f{}
	res.X = v.X * f
	res.Y = v.Y * f
	return &res
}

func (v *vector2f) scalarProduct(v2 vector2f) float64 {
	return v.X*v2.X + v.Y*v2.Y
}
