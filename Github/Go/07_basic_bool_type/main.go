package main

import "fmt"

func main(){
	isSunny := true
	isRainy := true
	canGoOutside := isSunny && !isRainy
	fmt.Println(canGoOutside)
}