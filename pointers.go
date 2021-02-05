package main

import "fmt"

func main () {

	var x int = 18
	var y int
	var ip *int  //ip is pointer to int
	ip = &x //ip now points to x
	y = *ip // y is now 18
	fmt.Println(y)
}