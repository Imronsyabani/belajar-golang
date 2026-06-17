package main

import "fmt"

func main(){
	name := "imron"

	switch name {
		case "imron":
			fmt.Println("Hello Imron")
		case "syabani":
			fmt.Println("Hello Syabani")
		default:
			fmt.Println("Aku tidak tau siapa kamu")
	}

	// Switch short Statement
	// switch length := len(name); length > 8 {
	// 	case true:
	// 		fmt.Println("Nama tidak valid")
	// 	case false:
	// 		fmt.Println("Nama valid")
	// }

	// Switch tanpa kondisi
	length := len(name)
	switch {
		case length > 8:
			fmt.Println("Nama Tidak valid")
		case length < 8: 
			fmt.Println("Nama Valid") 
	}
}