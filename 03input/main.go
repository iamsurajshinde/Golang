package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	nameMessage := "Enter your name"
	fmt.Println(nameMessage)

	read := bufio.NewReader(os.Stdin)

	name, _ := read.ReadString('\n')

	fmt.Println("Your name is ", name)

}
