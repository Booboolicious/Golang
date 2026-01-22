package main

import "fmt"


func math (a int, b int) (int, int, int){
	sum := a + b
	product := a * b
	div := a/b
	return sum, product, div
}

func main(){
	sum1, _, _ := math(10, 21)
	_, product, _ := math(5, 5)

	fmt.Println(sum1,product)
}