package main

import (
	"fmt"
)

// Vertex is just a name it can be anything
type Vertex struct {
	X int 
	Y int
}

func pointer_to_struct() {
	v := Vertex{6, 7}
	p := &v 
	p.X = 1e9
	fmt.Println(v)
}

// Why putting them all in the var block? and not just assign them a value like v1 :=
// That is because := can only be used inside a function, and at package level we must use
// var, const, type, struct
var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1}	// Y:0 is implicit
	v3 = Vertex{} // X:0 and Y:0
	p = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println((v.X))	// 4
	fmt.Println(v)	// {4 2}
	// *********************
	fmt.Println("After calling pointer_to_struct")
	pointer_to_struct()
	fmt.Println("Examples of Struct Literals")
	fmt.Println(v1, p, v2, v3)
	// To check the address of pointer need to use %p with printf,
	// without it go would dereference the pointer and give the value at the pointer
	fmt.Printf("%p %v\n", p, *p)
}
