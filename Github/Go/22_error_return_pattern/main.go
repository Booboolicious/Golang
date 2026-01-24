package main

import (
	"fmt"
	"log"
	"strconv"
)


func main(){
	fmt.Println(typ `"hey"`)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}


func run () error{
	input := `30`

	level, err := parseLevel(input)

	if err != nil{
		return err
	}
	fmt.Println(
		`Selected Level`,
		level,
	)
	return nil
}

func parseLevel(s string)(int, error){

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf(`level must be a number`)
	}
	if n < 1 || n > 7 {
		return 0, fmt.Errorf(`"Level must be btwn one and seven"`)
	}
	return  n, nil

}
