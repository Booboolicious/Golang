package main

import "fmt"

func main(){
	age := 28
	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
	}
}