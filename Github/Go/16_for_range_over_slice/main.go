package main

import "fmt"

func main (){
	view := []int {10, 20, 30, 40}

	total := 0

	for index, value := range view {
		fmt.Println("Day", index, "views", value)
		total += value
	}

	fmt.Println(total)
}