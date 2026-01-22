package main

import "fmt"

func main (){
	ages := map[string]int{
		// map[keyType]valuetype
		"ezekiel": 28,
		"Aniema": 22,
	}

	fmt.Println(ages["ezekiel"], len(ages))
}