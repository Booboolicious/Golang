package main

import (
	"fmt"
	"strconv"
)


func main(){
	fmt.Println(`"hey"`)
}


func run () error{
	input := `3`

	level, err := parseLevel(input)

	if err != nil{
		return err
	}
	fmt.Println(`Selected Level`, level)
	
}

func parseLevel(s string)(int, error){

	n, err := strconv.Atoi(s)
	if err != nil {
		return 404, fmt.Errorf(`level must be a number`)
	}
	return  n, nil

}
