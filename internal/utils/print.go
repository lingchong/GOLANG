package utils

type Vetor struct {
	X, Y, Z int
}

// 查找100以内构成三角的所有可能
func (v *Vetor) Abs(z int) {
	v.X = v.X * z
	v.Y = v.Y * z
}
