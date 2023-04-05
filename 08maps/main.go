package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps")

	lang := make(map[string]string)
	lang["PY"] = "Python"
	lang["JV"] = "Java"
	lang["JS"] = "JavaScript"
	lang["CS"] = "Csharp"

	fmt.Println("list of languages : ", lang)

	for key, value := range lang {
		fmt.Printf("key %s has value %s \n", key, value)
	}

	delete(lang, "JV")
	fmt.Println("list of languages : ", lang)

}
