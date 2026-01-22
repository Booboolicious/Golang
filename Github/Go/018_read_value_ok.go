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

	val, okB := m["b"]
	fmt.Println(val, okB)

	key := "p"

	if val, ok:= m[key]; ok{
		fmt.Println(key, val)
	} else {
		fmt.Println(key, "is not in the map")
	}

	for key := range m {
		fmt.Println(key)
	}

}