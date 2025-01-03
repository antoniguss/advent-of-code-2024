package util

// Vector represents a point in 2D space
type Vector struct {
	X int
	Y int
}

func (v Vector) Get() (X, Y int) {
	return v.X, v.Y
}

// withinBounds checks if the Vector is within the grid boundaries
func (v Vector) WithinBounds(limX, limY int) bool {
	return v.X >= 0 && v.X < limX && v.Y >= 0 && v.Y < limY
}

// add adds two Vectors
func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y}
}

// Subtracts v2 from v1
func (v1 Vector) Sub(v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y}
}

// Subtracts v2 from v1
func (v1 Vector) Scale(scalar int) Vector {
	return Vector{v1.X * scalar, v1.Y * scalar}
}

func (v Vector) Rotate90Right() Vector {
	return Vector{X: -v.Y, Y: v.X}
}

func (v Vector) Rotate90Left() Vector {
	return Vector{X: v.Y, Y: -v.X}
}

func (v Vector) Above(d int) Vector {
	return v.Add(Vector{X: 0, Y: d})
}

func (v Vector) Below(d int) Vector {
	return v.Add(Vector{X: 0, Y: -d})
}

func (v Vector) Left(d int) Vector {
	return v.Add(Vector{X: -d, Y: 0})
}

func (v Vector) Right(d int) Vector {
	return v.Add(Vector{X: d, Y: 0})
}
