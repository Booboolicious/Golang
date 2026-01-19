package main

import "fmt"
var sum int = 0

func main(){
	for i := 0; i <= 5; i++ {
		sum = i
		fmt.Println(i, sum)
	}


}