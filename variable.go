package main

import "fmt"

func main(){
	// version 1 untuk assign variable
	// var name string;
	// var umur uint8;
	
	// name = "Imron Syabani"
	// umur = 23
	
	// fmt.Println("Nama ==", name);
	// fmt.Println("Umur ==", umur);
	
	// name = "Dea Nurul Fatimah"
	// umur = 24
	
	// fmt.Println("Nama ==", name);
	// fmt.Println("Umur ==", umur);
	
	
	// version 2 untuk assign variable tanpa tipe data
	// var name = "Imron Syabani"
	// var umur = 23
	
	// fmt.Println("Nama ==", name);
	// fmt.Println("Umur ==", umur);
	
	// name = "Dea Nurul Fatimah"
	// umur = 24
	
	// fmt.Println("Nama ==", name);
	// fmt.Println("Umur ==", umur);
	
	// version 3 untuk assign variable dengan keyword ':='
	// variabel yang menggunakan keyword ':=' hanya boleh sekali diassign dengan nama yang sama
	// seperti constanta tapi bisa diubah
	name := "Imron Syabani"
	umur := 23
	
	fmt.Println("Nama ==", name);
	fmt.Println("Umur ==", umur);
	
	name = "Dea Nurul Fatimah"
	umur = 24
	
	fmt.Println("Nama ==", name);
	fmt.Println("Umur ==", umur);
	
}