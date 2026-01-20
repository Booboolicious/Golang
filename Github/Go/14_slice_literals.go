package main

import "fmt"

func main(){
	result := []string{"Ezekiel", "Augustine"}
	result[3] = "aba"
	fmt.Println(result, result[1], result[len(result)-2])
}