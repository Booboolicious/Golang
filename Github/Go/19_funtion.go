package main

import "fmt"


func add (a int, b int) int {
	return a + b
}

func main(){
	sum1 := add(10, 11)

	fmt.Println(sum1)
}