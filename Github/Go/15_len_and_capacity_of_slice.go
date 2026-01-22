package main

import "fmt"

func main(){

	scores := make([]int, 0,5)
	fmt.Println(scores, len(scores), cap(scores))
	
	scores = append(scores, 100)
	scores = append(scores, 200, 5000)
	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 690, 15000)
	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 90, 150)
	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 0, 510)
	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 10, 50)
	fmt.Println(scores, len(scores), cap(scores))


	todo := []string{"sleep,", "eat,", "code &"}
	extra :=[]string{"fuck hard"}

	todo = append(todo, extra...)
	fmt.Println(todo)
}