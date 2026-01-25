package main

import "fmt"


func main (){
	fmt.Println(`"hell ya"`)
}

func doWork (success bool) error {
	fmt.Println(`"start: resource acquired"`)
	defer fmt.Println(`"clean up: resource released"`)
}
