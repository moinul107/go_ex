package main

import "fmt"
import "reflect"

func main(){

//array
//Long hand way
//var students [2] string

//fmt.Println(students)

//fmt.Println(len(students))

//data set/assign
//students[0] = "Moinul"
//students[1] = "Hasan"

//data pull, data retrieve, data get
//fmt.Println(students[0])

//short hand way, string literals

//students := [...] string{"Moinul", "Hasan", "Shelby", "PB"}
//fmt.Println(students[3])

var students [2]string
//students[0] = "Moinul"
//students[1] = "Hasan"

//x := students [0:2]
//x := make([]string, 0)
var fruits []string
fruits= append(fruits, "Mango", "Apple")

fmt.Println(fruits)
fmt.Printf("%T\n", fruits)
fmt.Printf(%T\n, students)

a := reflect.TypeOf(students).Kind().String()
fmt.Println(a)



}