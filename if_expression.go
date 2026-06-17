package main

import "fmt"

func main(){
	name := "imron"

	if name == "syabani"{
		fmt.Println("hello syabani")
	} else if name == "imron"{
		fmt.Println("hello imron")
	} else {
		fmt.Println("Aku tidak tau kamu siapa.")
	}

	// if Short Statement
	if length := len(name); length > 8 {
		fmt.Println("Nama lebih dari 5 karakter")
	} else {
		fmt.Println("Nama kamu valid")
	}
}