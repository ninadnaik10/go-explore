package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
method on type Vertex
(v Vertex) is the special receiver argument
This is similar to methods in class in Java
Because go does not have class
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
This is pointer receiver compared to value receiver in the above example.
It is used to modify the original value since it is referenced by pointer and not value.
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Abser interface {
	Abs() float64
	Scale(float64)
}

/*
Error interface
*/
type error interface {
	Error() string
}

/*
Type parameters
*/
// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	defer fmt.Println("Defer statement")
	v := Vertex{3, 4}

	/*
		No need to write (&v).Scale(10). That's because for convenience Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
	*/
	v.Scale(10)

	p := &v
	// In this case, the method call p.Abs() is interpreted as (*p).Abs().
	fmt.Println(p.Abs())

	fmt.Println(v.Abs())

	var a Abser = &Vertex{5, 6}
	a.Scale(10)
	describe(a)
	fmt.Println(a.Abs())

	var i interface{}
	desc(i)

	i = 42
	desc(i)

	i = "hello"
	desc(i)

	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(ss)

	fmt.Println("Exiting now")
}

func describe(abser Abser) {
	fmt.Printf("(%v, %T)\n", abser, abser)
}

func desc(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
