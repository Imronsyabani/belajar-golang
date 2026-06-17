package main

import "fmt"

func main(){
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Counter ke:", counter)
	// 	counter++
	// }

	// for-loop dengan statement
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Counter ke:", counter)
	// }

	slices := []string{"imron", "syabani", "imronsyabani"}
	// for i := 0; i < len(slices); i++{
	// 	fmt.Println(slices[i])
	// }

	for _, name := range slices{
		// fmt.Println("Index", index, "=", name)
		fmt.Println("=", name)
	}

	person := make(map[string]string)
	person["nama"] = "imronsyabani"
	person["umur"] = "23"
	person["job"] = "programmer"

	for key, value := range person {
		fmt.Println("Key", key, "Values", value)
	}
}