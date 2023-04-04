package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()

	fmt.Println("Current Time is ", presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Mon"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.September, 2, 0, 0, 0, 0, time.Local)

	fmt.Println(createdDate.Format("02-01-2006 Monday"))

}
