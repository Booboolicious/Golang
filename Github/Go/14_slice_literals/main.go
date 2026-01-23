package main

import "fmt"

func main(){
	result := []string{"Ezekiel", "Augustine"}
	result = append(result, "Abaessien")
	fmt.Println(result, result[1], result[len(result)-2])
	result[0] = "Aniema"
	fmt.Println(result, result[1], result[len(result)-2])
}