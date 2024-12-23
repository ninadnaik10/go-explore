package main

import (
	"fmt"
)

// import "variables"

type todo struct {
	ID   string
	name string
}

var todos = []todo{
	{ID: "1", name: "hello"},
}

func main() {
	// variables()

	// fmt.Println("hello world")
	// var a uint8 = 12
	// b := "this"
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(add(2, 3))
	// fmt.Println(todos)

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello world"))
	// 	fmt.Println(r.Body)
	// })
	// http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(r.Body)
	// 	var nums Numbers
	//     if err := json.NewDecoder(r.Body).Decode(&nums); err != nil {
	//         http.Error(w, err.Error(), http.StatusBadRequest)
	//         return
	//     }

	//     sum := nums.A + nums.B

	// })
	// http.ListenAndServe(":8080", nil)

	// server()
	// immutable datatype
	var x int = 5
	y := x
	y = 7
	fmt.Println(x, y)

	// mutable datatype slices
	var a []int = []int{3, 4, 5}
	b := a
	b[0] = 100
	fmt.Println(a, b)

	//immutable datatype array
	var a1 [3]int = [3]int{3, 4, 5}
	b1 := a1
	b1[0] = 100
	fmt.Println(a1, b1)

	var a2 []int = []int{1, 2, 3}
	fmt.Println(&a2[0], &a2[1])

	/*
			 Mutable Data Types:
		   - Slice
		   - Map
		   - Channels

		Immutable Data Types:

		   - Boolean, Int, Float
		   - Pointers
		   - String
		   - Interfaces
	*/

	// pointers and references

	x1 := 7
	y1 := &x1
	*y1 = 8
	// var y2 *string
	// fmt.Println(*y2)
	// above code demonstrates nil pointer dereference and causes
	fmt.Println(x1)

}

func add(x, y int) (z1, z2 int) {
	defer fmt.Println("hello", z1, z2)
	z1 = x + y
	z2 = x - y
	return
}
