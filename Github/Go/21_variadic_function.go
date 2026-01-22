package main

import "fmt"

func sumAll (nums ...int) int {

	total := 0

	for _, val := range nums {
		fmt.Println(val)
		total += val
	}
	return total
}

func main (){
	fmt.Println(sumAll(1,2,3,5,9,3,4,7))
}