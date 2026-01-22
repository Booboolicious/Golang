package main 

import "fmt"

func main (){
	m := map[string]int{
		"a":10,
		"d":0,
		"e":210,
		"k":120,
	}
	fmt.Println("a", m["a"])
	fmt.Println("d", m["d"])
	fmt.Println("e", m["e"])
	fmt.Println("b", m["b"])
}