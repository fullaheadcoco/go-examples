package main

import "fmt"

// https://go.dev/blog/declaration-syntax
func main() {
	var x1 int     // variable x is int
	var x2 *int    // variable x is pointer to int
	var x3 [5]int  // variable x is 5 array of int
	var x4 [5]*int // variable x is 5 array of pointer to int
	var x5 *[5]int // variable x is pointer to 5 araray of int

	x1 = 1
	x2 = &x1
	x3 = [5]int{1, 2, 3, 4, 5}
	x4 = [5]*int{&x3[0], &x3[1], &x3[2], &x3[3], &x3[4]}
	x5 = &x3

	fmt.Println("x1 : ", x1)
	fmt.Println("x2 : ", x2, " / *x2 : ", *x2)
	fmt.Println("x3 : ", x3)
	fmt.Println("x4 : ", x4)
	fmt.Println(
		"*x4[0] : ", *x4[0], "\n",
		"*x4[1] : ", *x4[1], "\n",
		"*x4[2] : ", *x4[2], "\n",
		"*x4[3] : ", *x4[3], "\n",
		"*x4[4] : ", *x4[4])
	fmt.Println("x5 : ", x5, " / *x5 : ", *x5)
}
