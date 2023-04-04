package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter you age : ")
	read := bufio.NewReader(os.Stdin)

	age, _ := read.ReadString('\n')

	ageInt, err := strconv.Atoi(strings.TrimSpace(age))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("You age after 5 years ", ageInt+5)
	}

}
