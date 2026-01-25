package main

import (
	"errors"
	"fmt"
)


func main (){
	log(`"Case 1: success"`)
	if err := doWork(true); err != nil {
		log(`error`, err)
	}

	log(`"Case 1: fail early"`)
	if err := doWork(false); err != nil {
		log(`error`, err)
	}
}

func doWork (success bool) error {
	log(`"start: resource acquired"`)
	defer log(`"clean up: resource released"`)

	if !success {
		return  errors.New(`something went wrong, "returning early"`)
		
	}
	log(`work: doing somthing important`)
	log(`work: It's over`)
	
	return nil
}


func log(input ...any) {
    fmt.Println(input...)
}