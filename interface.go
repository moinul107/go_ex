package main

import "fmt"

type Email interface{
 
 Write(string, string, string) 
 Send()
 Read()
 
}

type Test struct{
 To string
 From string
 Subject string
 Message string
}


func (t Test) Write(to string, from string, subject string, message string){
 fmt.Println(to, from, subject)
}

func (t Test) send(){
 fmt.Println(t.To, "email sent")
}

func (t Test) Read(){
 fmt.Println(t.From, "received from")
}


func main(){

var tst Test
tst.To = "moinul107@gmail.com"
tst.From = "moinulh107@gmail.com"
tst.Subject = "Test email"
tst.Message = "This is a test"

tst.Write(tst.To, tst.From, tst.Subject, tst.Message)

}