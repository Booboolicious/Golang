package main

import (
	"errors"
	"fmt"
)


func main (){
	fmt.Println(`"Case 1: success"`)

	if err := doWork(true); err != nil {
		fmt.
	}
}

func doWork (success bool) error {
	fmt.Println(`"start: resource acquired"`)
	defer fmt.Println(`"clean up: resource released"`)

	if !success {
		return  errors.New(`something went wrong, "returning early"`)
		
		fmt.Println(`work: doing somthing important`)
		fmt.Println(`work: It's over`)
		
		return nil
	}
}


func log(input any){
	fmt.Println(log)
}