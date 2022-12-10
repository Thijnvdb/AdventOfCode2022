package types

import "math"

type Vector struct {
	X, Y int
}

func ZeroVector() Vector {
	return Vector{X: 0, Y: 0}
}

func GetDistanceBetween(a Vector, b Vector) int {
	dx := math.Abs(float64(b.X - a.X))
	dy := math.Abs(float64(b.Y - a.Y))

	min := math.Min(dx, dy)
	max := math.Max(dx, dy)

	diagonalSteps := min
	straightSteps := max - min

	return int(math.Sqrt(2)*diagonalSteps + straightSteps)
}

func (vec *Vector) Add(vector Vector) {
	vec.X += vector.X
	vec.Y += vector.Y
}

func (vec *Vector) Subtract(vector Vector) {
	vec.X -= vector.X
	vec.Y -= vector.Y
}

func (vec *Vector) Equals(vector Vector) bool {
	return vec.X == vector.X && vec.Y == vector.Y
}

func VectorAdd(a Vector, b Vector) Vector {
	res := a
	res.X += b.X
	res.Y += b.Y
	return res
}

func VectorSubtract(a Vector, b Vector) Vector {
	res := a
	res.X -= b.X
	res.Y -= b.Y
	return res
}
