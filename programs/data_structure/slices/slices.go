package main

import "fmt"

func show(x []int, def string) {
	fmt.Println("slice definition :",def)
	fmt.Println("elements :",x)
	fmt.Println("capacity : ",cap(x))
	fmt.Println("length : ",len(x))
}

func main() {
	array := [6]int{0, 1, 2, 3, 4, 5}
	var s1 []int = array[1:4]
	var s2 []int = array[2:]
	var s3 []int = array[:4]
	var s4 []int = array[1:2]
	fmt.Println("original elements :",array)
	show(s1, "array[1:4]")
	show(s2, "array[2:]")
	show(s3, "array[:4]")
	show(s4, "array[1:2]")

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	// nil slice example
	var p []int
	fmt.Println(p, len(p), cap(p))
	if p == nil {
		fmt.Println("nil!")
	}
	// use of make
	a := make([]int, 5)
	show(a, "a")

	b := make([]int, 0, 5)
	show(b, "b")

	c := b[:2]
	show(c, "c")

	d := c[2:5]
	show(d, "d")

	// append to slice
	var app []int
	app = append(app, 0)
	show(app, "app")
	app = append(app, 1)
	show(app, "app")
	// slices and their range
	for i, v := range app {
		fmt.Printf("app[%d] ; %d\n", i, v)
	}
	for i  := range app {
		fmt.Printf("app[%d] ; %d\n", i, app[i])
	}
	for _, value  := range app {
		fmt.Printf(" %d\n", value)
	}
}