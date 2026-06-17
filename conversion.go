package main

import "fmt"

func main(){
	var nilai32 int32 = 54200
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println("Nilai int32 ==", nilai32);
	fmt.Println("Nilai int64 ==", nilai64);
	fmt.Println("Nilai int16 ==", nilai16);
	fmt.Println("Nilai int8 ==", nilai8);


	fmt.Println("=================");

	var name = "imronsyabani"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}