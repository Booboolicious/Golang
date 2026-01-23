package main

import "fmt"


func math (a int, b int) (sum int, product int, div int){
	sum = a + b
	product = a * b
	div = a/b
	return
}

func main(){
	sum1, _, _ := math(10, 21)
	_, product, _ := math(5, 5)
	_, _, div := math(25, 5)

	fmt.Println(sum1,product, div)
}