package types

type Vector struct {
	X, Y int
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
	res.X += b.X
	res.Y += b.Y
	return res
}
