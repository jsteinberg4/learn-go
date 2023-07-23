package examples

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y float64
}

func (v Vector) Norm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector<%.2f, %.2f>", v.X, v.Y)
}

func main() {
	var vec Vector = Vector{5, 0}
	// fmt.Printf("Vector: %v\n", vec)
	// fmt.Printf("Norm: %f", vec.Norm())

	var vStringer fmt.Stringer = vec
	fmt.Println(vStringer)
}
