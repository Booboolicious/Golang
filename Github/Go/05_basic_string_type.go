package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName := "John"
	lastName := "Doe"
	age := 28
	height := 128.4
	fullName := fmt.Sprintf("%s-%s %d %.2f", firstName, lastName, age, height)
	fmt.Println(strings.ToUpper(fullName))

}