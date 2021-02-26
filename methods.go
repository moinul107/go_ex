package main

import "fmt"


type Doctor struct{
 Name string
 Education string
 Age int
 Salary float32
}

func (d Doctor)getName() string{
return d.Name
}

func (d Doctor)getSalaryInfo() float32{
return d.Salary

}


func main(){

var d Doctor
d.Name = "Tareq"
d.Education = "BDS"
d.Age = 30
d.Salary = 50000.50


fmt.Println(d.getName())
fmt.Println(d.getSalaryInfo())
 
}
