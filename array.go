package main

import "fmt"

func main(){
	// save dulu, isi kemudian (Manual)
	var names [3]string //[3] adalah limit data yang bisa dimasukan dalam 1 array
	names[0] = "imron"
	names[1] = "syabani"
	names[2] = "imronsyabani"
	fmt.Println(names[2])

	//langsung save dan isi arraynya 
	var values = [3]int {
		12,
		23,
		25,
	}
	fmt.Println(values)
	fmt.Println(values[2])

	//builtin function array
	fmt.Println(len(names))
	fmt.Println(len(values))

	fmt.Println(names[0])
	fmt.Println(values[1])

	//reassign value index 0
	values[0] = 22
	fmt.Println(values)

	// unik case jika buat array baru dan belum ada datanya ketika menggunakan
	// len() maka hasilnya sama dengan maksimal data
	var uniqArray [5]string;
	fmt.Println(len(uniqArray)) // hasilnya 5 walaupun belum ada data yang diassign

	//unlimite isi array
	var unlimitedArray = [...]string{}
	fmt.Println(unlimitedArray)
	fmt.Println(len(unlimitedArray))
}