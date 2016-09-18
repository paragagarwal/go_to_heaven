package main

import "fmt"

type Vertex struct {
	X, Y int
}

func add(v1, v2 Vertex) (Vertex) {
	x := v1.X+v2.X
	y := v1.Y+v2.Y
	return Vertex{ x,y}
}


var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// Simple usage of struct
	v_1 :=  Vertex{1, 1}
	v_2 :=  &Vertex{2, 2}
    fmt.Println(v_1)
    fmt.Println(v_1.X, v_1.Y)
    fmt.Println(*v_2)
    fmt.Println(v_2.X, v_2.Y)
    // Complex usage of struct
	fmt.Println(add(v_1, v_1))
	fmt.Println(add(v_1, *v_2))
	fmt.Println(add(*v_2, *v_2))
	// Struct Literals
	fmt.Println(v1, p,*p,  v2, v3)
}

 