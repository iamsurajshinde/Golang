package main

import "fmt"

func main() {
	var name string = "suraj"
	fmt.Println(name)
	fmt.Printf("My name is type of %T \n", name)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn is type of %T \n", isLoggedIn)

	var smallInt uint8 = 255 //256 err
	fmt.Println(smallInt)
	fmt.Printf("smallInt is type of %T \n", smallInt)

	var smallInt64 uint64 = 656667 // uint16, uint32, uint, int
	fmt.Println(smallInt64)
	fmt.Printf("smallInt64 is type of %T \n", smallInt64)

	var float float64 = 656667.7886662738 // float32
	fmt.Println(float)
	fmt.Printf("float is type of %T \n", float)

	// implicit type
	var github = "http://github.com/iamsurajshinde"
	fmt.Println(github)

	// no var v
	linkedIn := "linkedin.com/in/iamsurajshinde"
	fmt.Println(linkedIn)

	const jwtToken string = "FesKDkeslOQlsdfqSe"

	fmt.Println(jwtToken)

}
