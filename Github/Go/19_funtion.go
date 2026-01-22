package main

import "fmt"


func SumAndProduct (a int, b int) (int, int) (int, int){
	sum := a + b
	product := a * b
	div := a/b
	return sum, product, div
}

func main(){
	sum1, _, _ := SumAndProduct(10, 21)
	_, product, _ := SumAndProduct(5, 5)

	fmt.Println(sum1,product)
}