// Tipe Data Map
// * Map adalah tiipe data lain yang berisikan kumpulan data yang sama, namun kita bisa
//   menentukan jenis tipe data index yang akan kita gunakan
// * Sederhana nya, Map adalah tipe data kumpulan key-value, dimana kata kuncinya bersifat unik tidak boleh sama
// * Jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

package main

import "fmt"

func main(){
	// map[tipe_data_key]tipe_data_value
	person := map[string]string{
		"name": "imron",
		"age": "23",
	}
	person["address"] = "Tangsel"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["judul"] = "Impian"
	book["author"] = "syabani"
	book["ops"] = "salah"

	fmt.Println(book)
	delete(book, "ops")
	fmt.Println(book)
}