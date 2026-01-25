package main

import (
	"fmt"

	"github.com/Booboolicious/Golang/Github/Go/my"
	"github.com/Booboolicious/Golang/my"
)


func main (){
	my.Log(`"Case 1: success"`)
	my.Log.typeof(42)
	my.Log.typeof("Hello")
	if err := doWork(true); err != nil {
		my.Log(`error`, err)
	}

	my.Log(`"Case 1: fail early"`)
	if err := doWork(false); err != nil {
		my.Log(`error`, err)
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


