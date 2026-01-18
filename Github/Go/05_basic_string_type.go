package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fmt.Println(strings.ToUpper(fullName))
}