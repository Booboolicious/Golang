package main

import "fmt"

func sumAll (nums ...int) int {

	total := 0

	for _, val := range nums {
		
		total += val
	}
	return total
}

func main (){
	fmt.Println(sumAll(1,2,3,5,9,3,4,7))

	values:= []int{5,7,6,4,8}

	fmt.Println(sumAll(values...))


	// Anonymous function

	res := func (n int) int{
		return n*2
	}
	fmt.Println(res(3))
}