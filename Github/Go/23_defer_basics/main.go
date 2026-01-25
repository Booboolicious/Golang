package main

import (
	"fmt"
	// "reflect"
)


func main (){
	log(`"Case 1: success"`)
	log.typeof(42)
	log.typeof("Hello")
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
		return  log.err(`something went wrong, "returning early"`)
		
	}
	log(`work: doing somthing important`)
	log(`work: It's over`)
	
	return nil
}


