package main

import (
	."github.com/Booboolicious/Golang/my"
)


func main (){
	my.Log(`"Case 1: success"`)
	my.Log.Typeof(42)
	my.Log.Typeof("Hello")
	if err := doWork(true); err != nil {
		my.Log(`error`, err)
	}

	my.Log(`"Case 1: fail early"`)
	if err := doWork(false); err != nil {
		my.Log(`error`, err)
	}
}

func doWork (success bool) error {
	my.Log(`"start: resource acquired"`)
	defer my.Log(`"clean up: resource released"`)

	if !success {
		return  my.Log.Err(`something went wrong, "returning early"`)
		
	}
	my.Log(`work: doing somthing important`)
	my.Log(`work: It's over`)
	
	return nil
}


