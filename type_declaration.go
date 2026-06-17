package main

import "fmt"

func main() {
	// New Type
	type NoKTP string

	var ktpImron NoKTP = "12334353234"

	var deaNo string = "2222222"
	var ktpDea NoKTP = NoKTP(deaNo)

	fmt.Println(ktpImron)
	fmt.Println(ktpDea)
}