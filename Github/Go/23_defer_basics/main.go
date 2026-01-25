package main

import (
	."github.com/Booboolicious/Golang/my"
)


func main (){
	Log(`"Case 1: success"`)
	Log.Typeof(42)
	Log.Typeof("Hello")
	if err := doWork(true); err != nil {
		Log(`error`, err)
	}

	Log(`"Case 2: failed early"`)
	if err := doWork(false); err != nil {
		Log(`error`, err)
	}
}

func doWork (success bool) error {
	Log(`"start: resource acquired"`)
	defer Log(`"clean up: resource released"`)

	if !success {
		return  Log.Err(`something went wrong, "returning early"`)
		
	}
	Log(`work: doing somthing important`)
	Log(`work: It's over`)
	
	return nil
}


