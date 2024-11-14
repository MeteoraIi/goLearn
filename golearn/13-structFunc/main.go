package main

import "fmt"

type user struct{
	name string
	age int
}

func(u user)checkName(name string) bool{
	return u.name == name
}

func(u *user)setName(name string){
	u.name = name
}

func main(){
	u := user{"mete", 25}

	fmt.Println(u.checkName("wu"))

	u.setName("wu")

	fmt.Println(u.checkName("wu"))
}