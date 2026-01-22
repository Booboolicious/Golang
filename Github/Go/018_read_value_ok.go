package main 

import "fmt"

func main (){
	m := map[string]int{
		"a":10,
		"d":0,
		"e":210,
		"k":120,
	}
	// Loop over the map
	for key, value := range m {
		fmt.Println(key, value)
	}

	valD, okD := m["d"]
	fmt.Println(valD, okD)

	valB, okB := m["b"]
	fmt.Println(valB, okB)


	if val, ok:= m["k"]; ok{
		fmt.Println(val)
	} else {
		fmt.Println("Not in the map")
	}


}