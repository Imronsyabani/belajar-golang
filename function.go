package main

import "fmt"

// function with parameter
func sayHello(firstName string, lastName string){
	fmt.Println("Hello, Selamat Pagi!", firstName, lastName)
}

// function with return value
// harus mereturn sesuai tipe data yang di declare
func getHello(nama string) string {
	// return 21 // meraise error untyped int constant as string value in return statement
	return "Hallo " + nama
}

func main(){
	sayHello("imron", "syabani")
	fmt.Println(getHello("Imronsyabani"))
}