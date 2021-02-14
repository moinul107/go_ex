package main

import "fmt"

func main(){
/*
	fmt.Print("Enter you age here: ")
	var age int
	fmt.Scanf("%d", &age)
*/

//if boolean_expression

/*
if age < 3 { //0 to 2
	fmt.Println("infant")
}else if age>2 && age<13 { //2 to 12
	fmt.Println("Children")
}else if age>12 && age <= 19 {
	fmt.Println("Teenager")
}else{
	fmt.Println("Adult")
}
	//fmt.Println(age)
*/

/* switch age {
case 1,2:
	fmt.Println("infant")
	fallthrough
case 3,4,5,6,7,8,10,11,12:
	fmt.Println("children")
case 13,14,15,16,17,18,19:
	fmt.Println("teenager")
default:
	fmt.Println("Adult")
}
*/
/*
for i:=1; i<=9; i++ {
fmt.Println(i)
}
*/
/* students := []string{ "Asgor", "Moinul", "Anonnya" }

for i, std := range students {
	fmt.Println(i, std)
}
*/

i:=0
for i<10 {
fmt.Println(i, "Hello")
i++
}

}




