package main

import "fmt"

func main() {
	fmt.Println(" Maps knowledge")

	language := make(map[string]string)

	language["JS"] = "JavaScript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("language Map", language)
	fmt.Println("JS from language Map", language["JS"])

	delete(language, "PY")
	fmt.Println("language Map", language)


	for key, value := range language {
		fmt.Printf("For key %v, value is %v \n",key, value)
	}
}
