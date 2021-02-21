package main

import "fmt"

/*
func add(x int, y int) int{
//body
r := x+y
return r
}
*/

/*
func add(x, y int) int{
//body
r := x+y
return r
}
*/

/*
func add(x, y int) (r int){
//body
r = x+y
return r
}
*/

/*
func add(x, y int) (r int){
//body
r = x+y
return
}
*/

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}


func main(){

// x := add (10,30)
a,p := rectangle(10,10)
fmt.Println(a,p)
}