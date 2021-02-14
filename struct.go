package main

import "fmt"


//type keyword indicates extra data type.

/*
type book struct {
Title string
Author string
ISBN string
Price float32
Pages int
}
/*

/*
type Student struct{
Color []string
Name string
Age string
Roll int
}
*/


/*
var b1 book
b1.Title = "An introduction to programming in GO"
b1.Author = "CALEB DOXY"
b1.ISBN = "978-1478355823"
b1.Price = 10.50
b1.Pages = 165
fmt.Println(b1) 
*/

//string literals: A literal is a notation for representing a fixed value in source code.
//anonymous struct

b1 := struct{
Title string
Author string
ISBN string
Price float32
Pages int
}{
 Title: "An Introduction to programming in GO",
 Author: "Caleb Doxy",
 ISBN : "655555",
 Price: 01.50,
 Pages: 1111,
}

fmt.Println(b1)

