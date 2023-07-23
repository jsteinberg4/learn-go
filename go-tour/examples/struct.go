package examples

import "fmt"

type Vertex struct {
	X, Y int
}

func StructExample() {
	var (
		v1 Vertex  = Vertex{X: 1, Y: 2}
		v2 *Vertex = &v1
	)

	fmt.Printf("Vertex(%d, %d)\t(Original Object)\n", v1.X, v1.Y)
	fmt.Printf("Vertex(%d, %d)\t(Pointer to original -- explicit dereference)\n", (*v2).X, (*v2).Y)
	fmt.Printf("Vertex(%d, %d)\t(Pointer to original -- implicit dereference)\n", v2.X, v2.Y)
}
