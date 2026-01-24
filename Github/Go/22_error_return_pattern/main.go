package main

import (
	"fmt"
	"strconv"
)


func main(){
	fmt.Println(`"hey"`)
}


func run () error{

}

func parseLevel(s string)(int, error){

	n, err := strconv.Atoi(s)
	if err != nil {
		return
	}
}
