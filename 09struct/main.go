package main

import "fmt"

func main() {
	fmt.Println("Struct")
	emp := Employee{"suraj", "surajshindecse@gmail.com", true, "04-08-2022"}
	fmt.Println(emp)                                                     // {suraj surajshindecse@gmail.com true 04-08-2022}
	fmt.Printf("Emp info is : %+v \n", emp)                              //Emp info is : {Name:suraj Email:surajshindecse@gmail.com IsActive:true DOJ:04-08-2022}
	fmt.Printf("Emp name is %v and email is %v \n", emp.Name, emp.Email) //Emp name is suraj and email is surajshindecse@gmail.com

}

type Employee struct {
	Name     string
	Email    string
	IsActive bool
	DOJ      string
}
