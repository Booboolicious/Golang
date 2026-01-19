package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName := "John"
	lastName := "Doe"
	age := 28.4
	fullName := fmt.Sprintf("%s-%s %.2f", firstName, lastName, age)
	fmt.Println(strings.ToUpper(fullName))

}