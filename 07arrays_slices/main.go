package main

import (
	"fmt"
	"sort"
)

func main() {
	var friends [3]string

	friends[0] = "Pravin"
	friends[2] = "Akash"

	fmt.Println(friends)
	fmt.Println(len(friends)) // 3

	fri := []string{"Pravin", "Akash"}

	fri = append(fri, "Yuvraj", "Banty")

	fmt.Println(fri)
	// [Pravin Akash Yuvraj Banty]
	fri = append(fri[:3])
	fmt.Println(fri)
	// [Pravin Akash Yuvraj]
	fri = append(fri[1:3])
	fmt.Println(fri)
	// [Akash Yuvraj]
	fri = append(fri[1:])
	// [Yuvraj]
	fmt.Println(fri)

	values := make([]int, 3)

	values[0] = 4
	values[1] = 2
	values[2] = 6
	// values[3] = 9  // err

	values = append(values, 8, 3, 9)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	index := 2
	values = append(values[:index], values[index+1:]...)
	fmt.Println(values)

}
