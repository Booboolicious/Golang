package main

import "fmt"

func main (){
	ages := map[string]int{
		// map[keyType]valuetype
		"ezekiel": 28,
		"Aniema": 22,
	}

	fmt.Println(ages["ezekiel"], len(ages))

	var scores map[string]int

	fmt.Println(scores, scores["A"])
	
	scores = make(map[string]int)
	scores["A"] = 5
	fmt.Println(scores)
}