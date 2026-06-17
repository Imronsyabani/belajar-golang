package main

import "fmt"

func main(){

	nama1 := "imron"
	nama2 := "Imron"

	var nameEqual = nama1 == nama2
	var sameLen = len(nama1) == len(nama2)

	fmt.Println(nameEqual)
	fmt.Println(sameLen)

	fmt.Println("=================");
	// && and operator
	fmt.Println("And Opearator 1 :", true && true);
	fmt.Println("And Opearator 2 :", true && false)
	
	// || or operator
	fmt.Println("Or Opearator 1 :", true || true);
	fmt.Println("Or Opearator 2 :", true || false)

	// ! not operator
	fmt.Println("Not Opearator 1 :", !true);
	fmt.Println("Not Opearator 2 :", !false)

}