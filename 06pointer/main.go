package main

import "fmt"

func main() {
	fmt.Println("--------pointers ---------------")
	var ptr *int
	fmt.Println("Default value of ptr is ", ptr) // Default value of ptr is  <nil>

	myAge := 30

	ptr = &myAge

	fmt.Println("value of ptr is ", *ptr)  //value of ptr is  30
	fmt.Println("address of ptr is ", ptr) // address of ptr is  0xc000016088

	*ptr += 5

	fmt.Println("Age after five years = ", myAge) //Age after five years =  35

}
