package main

import "fmt"

func main() {
	var age = 23

	//shorthand
	// name := "Daud"

	name, email := "Immanuel Daud", "immanueldaud@outlook.com"

	fmt.Println(name, age, email)
	fmt.Printf("%T\n", age)
}
