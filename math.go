package main

import "fmt"

func main(){
	var a = 15;
	var b = 10;

	fmt.Println(a + b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a - b)
	fmt.Println(a % b)

	// Augmented Assignment
	a += b
	a -= b
	a /= b
	a *= b
	fmt.Println(a + b)

	// unary operator
	var c = 10
	var d = 10
	c++
	fmt.Println(c)
	d--
	fmt.Println(d)
	fmt.Println(-d)
	fmt.Println(-c)
	fmt.Println(!true)
	fmt.Println(!false)
}