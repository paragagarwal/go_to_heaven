// program to show usage of functions

package main

import "fmt"

func add(x, y int) int {
	return x+y
}

func multiply(x, y int) int {
	return x*y
}

func divide(x, y int) int {
	return x/y
}

func mod(x, y int) int {
	return x%y
}

func factorial(n int) int {
	if n > 0 {
		return n * factorial(n-1)
	}
	return 1
}

func main(){
	fmt.Println(add(1, 2))
	fmt.Println(multiply(1, 2))
	fmt.Println(divide(4, 2))
	fmt.Println(mod(4, 2))
	fmt.Println(add(add(2,2), add(2,2)))
	fmt.Println(factorial(10))
}

