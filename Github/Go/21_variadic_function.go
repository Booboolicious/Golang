package main

import "fmt"

func sumAll (nums ...int){

	total := 0

	for _, val := range nums {
		total+= val
	}
	return total
}